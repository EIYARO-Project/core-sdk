package api

import (
	"encoding/json"
	"errors"
	"fmt"
	"io"

	"github.com/EIYARO-Project/core-sdk/api/types"
)

type GetWorkJson struct {
	BlockHeader types.BlockHeader `json:"block_header"`
	Seed        string            `json:"seed"`
}

func (a *Api) GetWorkJson() (*GetWorkJson, error) {
	response, err := a.client.Get("get-work-json")
	if err != nil {
		return nil, err
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
		apiMessage, err := NewApiMessageGetWorkJSON(body)
		if err != nil {
			return nil, err
		}
		result := apiMessage.Data

		return &result, err

	} else {
		return nil, errors.New("the call to the API returned a status of fail")
	}
}

func NewApiMessageGetWorkJSON(data []byte) (*APiMessageSuccessObject[GetWorkJson], error) {
	var apiMessage APiMessageSuccessObject[GetWorkJson]
	err := json.Unmarshal(data, &apiMessage)
	return &apiMessage, err
}

func NewGetWorkJson(data []byte) (*GetWorkJson, error) {
	var getWorkJson GetWorkJson
	err := json.Unmarshal(data, &getWorkJson)
	return &getWorkJson, err
}

func (gwj *GetWorkJson) Marshal() ([]byte, error) {
	result, err := json.Marshal(gwj)
	return result, err
}

func (gwj *GetWorkJson) MarshalIndent() ([]byte, error) {
	result, err := json.MarshalIndent(gwj, "", "    ")
	return result, err
}

func (gwj *GetWorkJson) String() string {
	j, err := gwj.Marshal()
	if err != nil {
		fmt.Printf("Error encoding NetInfo to JSON: %s", err)
	}

	return string(j)
}

func (gwj *GetWorkJson) StringIndent() string {
	j, err := gwj.MarshalIndent()
	if err != nil {
		fmt.Printf("Error encoding NetInfo to JSON: %s", err)
	}

	return string(j)
}
