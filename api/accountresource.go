package api

import (
	"encoding/json"
	"errors"
	"io"
	"net/http"
	"strings"

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

func (ar AccountResource) List() ([]Resources, error) {
	response, err := ar.client.Post("list-accounts", "{}")
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

func (ar AccountResource) View(a ...string) (Resources, error) {
	if len(a) != 2 {
		return nil, errors.New("incorrect number of parameters: Need two")
	} else {
		first := a[0]
		second := a[1]

		var response *http.Response
		var err error
		switch strings.ToLower(first) {
		case "id":
			response, err = ar.client.Post("list-accounts", "{\"id\":\""+second+"\"}")
			if err != nil {
				return nil, err
			}
		case "alias":
			response, err = ar.client.Post("list-accounts", "{\"alias\":\""+second+"\"}")
			if err != nil {
				return nil, err
			}
		default:
			return nil, errors.New("Unknown filter named " + string(first))
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
			var apiMessage APiMessageSuccessArray[resources.Account]
			err := json.Unmarshal(body, &apiMessage)
			if err != nil {
				return nil, err
			}

			if len(apiMessage.Data) == 0 {
				return nil, errors.New("no accounts returned")
			}

			if len(apiMessage.Data) > 1 {
				return nil, errors.New("more than one account returned")
			}

			return apiMessage.Data[0], nil

		} else {
			return nil, errors.New("the call to the API returned a status of fail")
		}
	}
}
