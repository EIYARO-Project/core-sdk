package client

import (
	"net/http"
)

type ClientInterface interface {
	// The GET HTTP method
	Get(path string) (*http.Response, error)
}
