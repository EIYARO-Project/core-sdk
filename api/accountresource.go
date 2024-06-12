package api

import (
	"encoding/json"
	"errors"
	"io"

	"github.com/EIYARO-Project/core-sdk/api/resources"
	"github.com/EIYARO-Project/core-sdk/client"
)

type AccountResource struct {
	client client.ClientInterface
}

func NewAccountResource(client client.ClientInterface) AccountResource {
	return AccountResource{
		client: client,
	}
}

func (at AccountResource) List() ([]Resources, error) {
	response, err := at.client.Post("list-accounts", "{}")
	if err != nil {
		return []Resources{}, err
	}

	body, err := io.ReadAll(response.Body)
	if err != nil {
		return nil, err
	}

	var message map[string]interface{}
	if err := json.Unmarshal(body, &message); err != nil {
		return nil, err
	}

	status, ok := message["status"]
	if !ok {
		return nil, errors.New("did not find field status")
	}

	if status == "success" {
		var apiMessage APiMessageSuccessArray[resources.Account]
		err := json.Unmarshal(body, &apiMessage)
		if err != nil {
			return nil, err
		}
		var result []Resources
		for _, value := range apiMessage.Data {
			result = append(result, value)
		}
		return result, nil

	} else {
		return nil, errors.New("the call to the API returned a status of fail")
	}
}

// func (at AccessTokenResource) View() (Resources, error) {
// 	result := AccessToken{}
// 	return result, nil
// }
