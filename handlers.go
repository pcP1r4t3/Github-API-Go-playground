package main

import (
	"encoding/json"
	"github.com/sirupsen/logrus"
	"net/http"
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

	repos, err := getRepos(h.GithubAccessToken)
	if err != nil {
		h.logger.WithError(err).Error("Failed to get public repositories")
	}

	t := NewTaskScheduler(h.GithubAccessToken, repos)
	t.merge()

	w.Header().Add("Content-Type", "application/json")
	w.WriteHeader(http.StatusOK)

	if err = json.NewEncoder(w).Encode(repos); err != nil {
		h.logger.WithError(err).Error("Fail to encode JSON")
	}

	return nil
}
