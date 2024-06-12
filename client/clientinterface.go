package client

import (
	"net/http"
)

type ClientInterface interface {
	// The GET HTTP method
	Get(endpoint string) (*http.Response, error)
}
