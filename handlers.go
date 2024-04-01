package main

import (
	"encoding/json"
	"github.com/sirupsen/logrus"
	"net/http"
	"strings"
)

type MyCustomHandler struct {
	GithubAccessToken string
	logger            logrus.FieldLogger
}

func NewMyCustomHandler(logger logrus.FieldLogger, token string) *MyCustomHandler {
	return &MyCustomHandler{
		GithubAccessToken: token,
		logger:            logger,
	}
}

func (h *MyCustomHandler) reposHandler(w http.ResponseWriter, r *http.Request, _ map[string]string) error {

	//query params
	language := r.URL.Query().Get("language")
	license := r.URL.Query().Get("license")

	repos, err := getRepos(h.GithubAccessToken)
	if err != nil {
		h.logger.WithError(err).Error("Failed to get public repositories")
	}

	t := NewTaskScheduler(h.GithubAccessToken, repos)
	t.merge()

	repos = h.filterByQueryParams(repos, language, license)

	w.Header().Add("Content-Type", "application/json")
	w.WriteHeader(http.StatusOK)

	if err = json.NewEncoder(w).Encode(repos); err != nil {
		h.logger.WithError(err).Error("Fail to encode JSON")
	}

	return nil
}

func (h *MyCustomHandler) filterByQueryParams(repos []map[string]any, language string, license string) []map[string]any {
	var filteredRepos []map[string]any
	for _, r := range repos {
		found := false

		repoLicense, ok := r["license"].(map[string]any)
		if !ok {
			continue
		}
		if strings.ToUpper(repoLicense["key"].(string)) == strings.ToUpper(license) {
			found = true
		}

		repoLanguages, ok := r["languages"].(map[string]any)
		if !ok {
			continue
		}

		for l, _ := range repoLanguages {
			if strings.ToUpper(l) == strings.ToUpper(language) {
				found = true
				break
			}
		}

		if found {
			filteredRepos = append(filteredRepos, r)
		}
	}

	return filteredRepos
}
