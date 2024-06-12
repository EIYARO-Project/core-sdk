package api

import (
	"encoding/json"
	"errors"
	"fmt"
	"io"
)

type Difficulty struct {
	Hash        string `json:"hash"`
	BlockHeight uint64 `json:"height"`
	Bits        uint64 `json:"bits"`
	Difficulty  string `json:"difficulty"`
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
		return nil, errors.New("did not find field status")
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

func NewApiMessageDifficulty(data []byte) (*APiMessageSuccessObject[Difficulty], error) {
	var apiMessage APiMessageSuccessObject[Difficulty]
	err := json.Unmarshal(data, &apiMessage)
	return &apiMessage, err
}

func NewDifficulty(data []byte) (*Difficulty, error) {
	var difficulty Difficulty
	err := json.Unmarshal(data, &difficulty)
	return &difficulty, err
}

func (d *Difficulty) Marshal() ([]byte, error) {
	result, err := json.Marshal(d)
	return result, err
}

func (d *Difficulty) MarshalIndent() ([]byte, error) {
	result, err := json.MarshalIndent(d, "", "    ")
	return result, err
}

func (d *Difficulty) String() string {
	j, err := d.Marshal()
	if err != nil {
		fmt.Printf("Error encoding Difficulty to JSON: %s", err)
	}

	return string(j)
}

func (d *Difficulty) StringIndent() string {
	j, err := d.MarshalIndent()
	if err != nil {
		fmt.Printf("Error encoding Difficulty to JSON: %s", err)
	}

	return string(j)
}
