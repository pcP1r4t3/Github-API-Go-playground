package main

import (
	"encoding/json"
	"fmt"
	"net/http"
)

const (
	GithhubPublicReposEndpoint  = "https://api.github.com/repositories"
	GithhubRepoLicensesEndpoint = "https://api.github.com/repos/%s/license"
	EmptyResponse               = "{}"
)

func getRepos(token string) ([]map[string]any, error) {
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

func getLanguages(repoLangsEndpoint, token string) any {
	var languages any

	client := http.Client{}
	req := NewRequestBuilder().
		WithMethod(http.MethodGet).
		WithURL(repoLangsEndpoint).
		WithTokenAuth(token).
		Build()
	res, err := client.Do(req)
	if err != nil {
		return EmptyResponse
	}
	if res.StatusCode != http.StatusOK {
		return EmptyResponse
	}

	if err = json.NewDecoder(res.Body).Decode(&languages); err != nil {
		return EmptyResponse
	}

	return languages
}

func getLicence(repoName, token string) any {
	var (
		license            map[string]any
		licenseInnerObject any
	)

	client := http.Client{}
	req := NewRequestBuilder().
		WithMethod(http.MethodGet).
		WithURL(fmt.Sprintf(GithhubRepoLicensesEndpoint, repoName)).
		WithTokenAuth(token).
		Build()
	res, err := client.Do(req)
	if err != nil {
		return EmptyResponse
	}
	if res.StatusCode != http.StatusOK {
		return EmptyResponse
	}

	if err = json.NewDecoder(res.Body).Decode(&license); err != nil {
		return EmptyResponse
	}
	licenseInnerObject, ok := license["license"]
	if !ok {
		return EmptyResponse
	}

	return licenseInnerObject
}
