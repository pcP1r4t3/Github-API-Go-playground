package main

import (
	"fmt"
	"github.com/gorilla/mux"
	"net/http"
)

type RequestBuilder struct {
	method    string
	url       string
	tokenAuth string
	params    []string
}

func NewRequestBuilder() *RequestBuilder {
	return &RequestBuilder{}
}

func (r *RequestBuilder) WithMethod(method string) *RequestBuilder {
	r.method = method
	return r
}

func (r *RequestBuilder) WithURL(url string) *RequestBuilder {
	r.url = url
	return r
}

func (r *RequestBuilder) WithURLParams(url string, params ...string) *RequestBuilder {
	r.url = url
	r.params = params
	return r
}

func (r *RequestBuilder) WithTokenAuth(token string) *RequestBuilder {
	r.tokenAuth = token
	return r
}

func (r *RequestBuilder) Build() *http.Request {
	return r.createRequest()
}

func (r *RequestBuilder) createRequest() *http.Request {
	req, err := http.NewRequest(r.method, r.url, nil)
	if err != nil {
		return nil
	}

	if len(r.params) > 0 {
		i := 0
		varsKeys := mux.Vars(req)
		for key, _ := range varsKeys {
			varsKeys[key] = r.params[i]
			i++
		}
		mux.SetURLVars(req, varsKeys)
	}

	if r.tokenAuth != "" {
		req.Header.Add("Authorization", fmt.Sprintf("Bearer %s", r.tokenAuth))
	}

	return req
}
