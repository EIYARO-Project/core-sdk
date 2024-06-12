package api

import (
	"encoding/json"
	"errors"
	"io"

	"github.com/EIYARO-Project/core-sdk/api/resources"
	"github.com/EIYARO-Project/core-sdk/client"
)

type AccessTokenResource struct {
	client client.ClientInterface
}

func NewAccessTokenResource(client client.ClientInterface) AccessTokenResource {
	return AccessTokenResource{
		client: client,
	}
}

func (at AccessTokenResource) List() ([]Resources, error) {
	response, err := at.client.Get("list-access-tokens")
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
		var apiMessage APiMessageSuccessArray[resources.AccessToken]
		err := json.Unmarshal(body, &apiMessage)
		if err != nil {
			return nil, err
		}
		var result []Resources
		for value := range apiMessage.Data {
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
