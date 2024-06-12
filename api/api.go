package api

import (
	"encoding/json"
	"errors"
	"io"

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

func (a *Api) NetInfo() (*NetInfo, error) {
	response, err := a.client.Get("net-info")
	if err != nil {
		return nil, err
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
		return nil, errors.New("dif not find field status")
	}

	if status == "success" {
		apiMessage, err := NewApiMessageNetInfo(body)
		if err != nil {
			return nil, err
		}
		result := apiMessage.Data

		return &result, err

	} else {
		return nil, errors.New("the call to the API returned a status of fail")
	}
}

func (a *Api) Difficulty(block_height uint64, block_hash string) (*Difficulty, error) {
	blockRequest := BlockRequest{
		BlockHeight: block_height,
		BlockHash:   block_hash,
	}
	response, err := a.client.Post("get-difficulty", blockRequest.String())
	if err != nil {
		return nil, err
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
		return nil, errors.New("dif not find field status")
	}

	if status == "success" {
		apiMessage, err := NewApiMessageDifficulty(body)
		if err != nil {
			return nil, err
		}
		result := apiMessage.Data

		return &result, err

	} else {
		return nil, errors.New("the call to the API returned a status of fail")
	}
}
