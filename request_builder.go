package main

import (
	"fmt"
	"net/http"
)

type RequestBuilder struct {
	method    string
	url       string
	tokenAuth string
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

	if r.tokenAuth != "" {
		req.Header.Add("Authorization", fmt.Sprintf("Bearer %s", r.tokenAuth))
	}

	return req
}
