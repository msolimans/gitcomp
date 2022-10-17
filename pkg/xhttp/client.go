package xhttp

import (
	"bytes"
	"context"
	"encoding/json"
	"io"
	"net/http"

	"golang.org/x/oauth2"
)

//I can wrap this in a struct and add more functionality to it
//fo example like add specific timeout to requests and generalize that or send specific headers and unify requests
//accepts token and returns http client
func NewHttpClient(ctx context.Context, token string) *http.Client {
	ts := oauth2.StaticTokenSource(
		&oauth2.Token{AccessToken: token},
	)
	return oauth2.NewClient(ctx, ts)
}

// wrapper around http.NewRequest
func NewRequest(method, urlStr string, headers map[string]string, body interface{}) (*http.Request, error) {
	var buf io.ReadWriter
	if body != nil {
		buf = &bytes.Buffer{}
		enc := json.NewEncoder(buf)
		enc.SetEscapeHTML(false)
		err := enc.Encode(body)
		if err != nil {
			return nil, err
		}
	}
	
	req, err := http.NewRequest(method, urlStr, buf)
	if err != nil {
		return nil, err
	}

	if body != nil {
		req.Header.Set("Content-Type", "application/json")
	}
	
	for k, h := range headers {
		req.Header.Set(k, h)
	}
	
	return req, nil
}
