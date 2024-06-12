package client

import (
	"io"
	"net/http"
	"net/url"
)

type StdClient struct {
	baseURL     string
	accessToken string
	client      http.Client
}

func (sc StdClient) Get(path string) (*http.Response, error) {
	url, err := url.Parse(sc.baseURL)
	if err != nil {
		return nil, err
	}
	url.Path = path

	var body io.Reader
	request, err := http.NewRequest("GET", url.String(), body)
	if err != nil {
		return nil, err
	}
	request.Header.Add("Content-Type", "application/json")
	request.Header.Add("User-Agent", "EIYARO/v0.1.0")

	response, err := sc.client.Do(request)
	if err != nil {
		return nil, err
	}

	return response, nil
}

func NewStdClient(baseURL string, accessToken string) *StdClient {
	return &StdClient{
		baseURL:     baseURL,
		accessToken: accessToken,
		client:      http.Client{},
	}
}
