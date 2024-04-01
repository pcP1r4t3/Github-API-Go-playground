package main

import (
	"encoding/json"
	"fmt"
	"net/http"
)

const GithhubPublicReposEndpoint = "https://api.github.com/repositories"

func getRepos(token string) ([]map[string]any, error) {
	fmt.Println("repos token", token)
	var repos []map[string]any

	client := http.Client{}
	req := NewRequestBuilder().
		WithMethod(http.MethodGet).
		WithURL(GithhubPublicReposEndpoint).
		WithTokenAuth(token).
		Build()
	res, err := client.Do(req)
	if err != nil {
		return nil, err
	}
	if res.StatusCode != http.StatusOK {
		return nil, err
	}

	if err = json.NewDecoder(res.Body).Decode(&repos); err != nil {
		return nil, err
	}

	return repos, nil
}

func getLanguages(repoLangsEndpoint, token string) map[string]any {
	var languages map[string]any

	client := http.Client{}
	req := NewRequestBuilder().
		WithMethod(http.MethodGet).
		WithURL(repoLangsEndpoint).
		WithTokenAuth(token).
		Build()
	res, err := client.Do(req)
	if err != nil {
		return nil
	}
	if res.StatusCode != http.StatusOK {
		return nil
	}

	if err = json.NewDecoder(res.Body).Decode(&languages); err != nil {
		return nil
	}

	return languages
}
