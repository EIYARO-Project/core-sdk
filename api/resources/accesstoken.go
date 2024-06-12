package resources

import (
	"encoding/json"
	"fmt"
)

type AccessToken struct {
	ID        string `json:"id"`
	Token     string `json:"token"`
	CreatedAt string `json:"created_at"`
}

func NewAccessToken(data []byte) (AccessToken, error) {
	var accessToken AccessToken
	err := json.Unmarshal(data, &accessToken)
	return accessToken, err
}

func (at *AccessToken) Marshal() ([]byte, error) {
	result, err := json.Marshal(at)
	return result, err
}

func (at *AccessToken) MarshalIndent() ([]byte, error) {
	result, err := json.MarshalIndent(at, "", "    ")
	return result, err
}

func (at *AccessToken) String() string {
	j, err := at.Marshal()
	if err != nil {
		fmt.Printf("Error encoding AccessToken to JSON: %s", err)
	}

	return string(j)
}

func (at *AccessToken) StringIndent() string {
	j, err := at.MarshalIndent()
	if err != nil {
		fmt.Printf("Error encoding AccessToken to JSON: %s", err)
	}

	return string(j)
}
