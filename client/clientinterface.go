package client

import (
	"net/http"
)

type ClientInterface interface {
	// The GET HTTP method
	Get(path string) (*http.Response, error)

	// The POST HTTP method
	Post(path string, body string) (*http.Response, error)
}
