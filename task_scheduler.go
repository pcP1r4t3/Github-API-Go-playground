package main

import (
	"sync"
)

type TaskScheduler struct {
	repos             []map[string]any
	githubAccessToken string
}

type GeneratorTaskItem struct {
	languagesURL string
	ID           int
}

type LanguagesWorkerTaskItem struct {
	languages map[string]any
	ID        int
}

func NewTaskScheduler(token string, repos []map[string]any) *TaskScheduler {
	return &TaskScheduler{
		githubAccessToken: token,
		repos:             repos,
	}
}

func (t *TaskScheduler) workerLanguages(in chan GeneratorTaskItem) chan LanguagesWorkerTaskItem {
	out := make(chan LanguagesWorkerTaskItem)

	for repo := range in {
		go func(repoIdx int) {
			lang := getLanguages(repo.languagesURL, t.githubAccessToken)
			out <- LanguagesWorkerTaskItem{
				languages: lang,
				ID:        repoIdx,
			}
		}(repo.ID)
	}

	return out

}

func (t *TaskScheduler) merge() {
	var wg sync.WaitGroup
	wg.Add(len(t.repos))
	chRepos := t.generateSourceChannel()
	chLangs := t.workerLanguages(chRepos)

	go func() {
		for l := range chLangs {
			t.repos[l.ID]["languages"] = l.languages
			wg.Done()
		}
	}()

	wg.Wait()

}
func (t *TaskScheduler) generateSourceChannel() chan GeneratorTaskItem {
	ch := make(chan GeneratorTaskItem)
	go func() {
		defer close(ch)
		for repoIdx, repo := range t.repos {
			item := GeneratorTaskItem{
				languagesURL: repo["languages_url"].(string),
				ID:           repoIdx,
			}
			ch <- item
		}
	}()

	return ch
}
