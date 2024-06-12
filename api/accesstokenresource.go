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

	success, err := isMessageSuccess(body)
	if err != nil {
		return nil, err
	}

	if success {
		var apiMessage APiMessageSuccessArray[resources.AccessToken]
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

func (at AccessTokenResource) View(a ...string) (Resources, error) {
	result := resources.AccessToken{}
	return result, nil
}
