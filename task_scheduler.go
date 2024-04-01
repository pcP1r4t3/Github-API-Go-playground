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
	repoName     string
	ID           int
}

type LanguagesWorkerTaskItem struct {
	languages any
	ID        int
}

type LicenseWorkerTaskItem struct {
	License any
	ID      int
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

func (t *TaskScheduler) workerLicense(in chan GeneratorTaskItem) chan LicenseWorkerTaskItem {
	out := make(chan LicenseWorkerTaskItem)

	for repo := range in {
		go func(repoIdx int) {
			license := getLicence(repo.repoName, t.githubAccessToken)
			out <- LicenseWorkerTaskItem{
				License: license,
				ID:      repoIdx,
			}
		}(repo.ID)
	}

	return out

}

func (t *TaskScheduler) merge() {
	var wg sync.WaitGroup
	wg.Add(2 * len(t.repos))
	chRepos1 := t.generateSourceChannel()
	chRepos2 := t.generateSourceChannel()
	chLangs := t.workerLanguages(chRepos1)
	chLicense := t.workerLicense(chRepos2)

	go func() {
		for l := range chLangs {
			t.repos[l.ID]["languages"] = l.languages
			wg.Done()
		}
	}()

	go func() {
		for l := range chLicense {
			t.repos[l.ID]["license"] = l.License
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
				ID:           repoIdx,
				languagesURL: repo["languages_url"].(string),
				repoName:     repo["full_name"].(string),
			}
			ch <- item
		}
	}()

	return ch
}
