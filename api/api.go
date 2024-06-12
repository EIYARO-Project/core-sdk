package api

import (
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
