package api

import (
	"errors"
	"strings"

	"github.com/EIYARO-Project/core-sdk/client"
)

type Api struct {
	client client.ClientInterface
}

func NewApi(client client.ClientInterface) *Api {
	return &Api{
		client: client,
	}
}

func (a *Api) Resource(resourceName string) (ResourceInterface[Resources], error) {
	switch strings.ToLower(resourceName) {
	case "accesstoken":
		result := NewAccessTokenResource(a.client)
		return result, nil
	default:
		return nil, errors.New("unknown resource type")
	}
}
