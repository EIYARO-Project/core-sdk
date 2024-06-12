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

type AccessTokenResource struct {
	client client.ClientInterface
}

func NewAccessTokenResource(client client.ClientInterface) AccessTokenResource {
	return AccessTokenResource{
		client: client,
	}
}

func (at AccessTokenResource) List() ([]Resources, error) {
	response, err := at.client.Post("list-access-tokens", "{}")
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
	if len(a) != 2 {
		return nil, errors.New("incorrect number of parameters: Need two")
	} else {
		first := a[0]
		second := a[1]

		var response *http.Response
		var err error
		switch strings.ToLower(first) {
		case "id":
			response, err = at.client.Post("list-access-tokens", "{\"id\":\""+second+"\"}")
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
			var apiMessage APiMessageSuccessArray[resources.AccessToken]
			err := json.Unmarshal(body, &apiMessage)
			if err != nil {
				return nil, err
			}

			if len(apiMessage.Data) == 0 {
				return nil, errors.New("no access tokens returned")
			}

			var result []Resources
			if len(apiMessage.Data) > 1 {
				for _, value := range apiMessage.Data {
					if value.ID == second {
						result = append(result, value)
					}
				}

				if len(result) == 0 {
					return nil, errors.New("no access tokens returned")
				}
			}

			return result[0], nil

		} else {
			return nil, errors.New("the call to the API returned a status of fail")
		}
	}
}
