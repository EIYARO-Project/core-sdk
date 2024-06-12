package api

import (
	"encoding/json"
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
	case "account":
		result := NewAccountResource(a.client)
		return result, nil
	default:
		return nil, errors.New("unknown resource type")
	}
}

func isMessageSuccess(body []byte) (bool, error) {
	var message map[string]interface{}
	if err := json.Unmarshal(body, &message); err != nil {
		return false, err
	}

	status, ok := message["status"]
	if !ok {
		return false, errors.New("did not find field status")
	}

	return (status == "success"), nil
}
