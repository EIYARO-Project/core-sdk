package api

import (
	"encoding/json"
	"errors"
	"fmt"
	"io"
)

type NetInfo struct {
	Listening    bool        `json:"listening"`
	Syncing      bool        `json:"syncing"`
	Mining       bool        `json:"mining"`
	PeerCount    int         `json:"peer_count"`
	CurrentBlock uint64      `json:"current_block"`
	HighestBlock uint64      `json:"highest_block"`
	NetworkID    string      `json:"network_id"`
	VersionInfo  VersionInfo `json:"version_info"`
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

	success, err := isMessageSuccess(body)
	if err != nil {
		return nil, err
	}

	if success {
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

func NewApiMessageNetInfo(data []byte) (*APiMessageSuccessObject[NetInfo], error) {
	var apiMessage APiMessageSuccessObject[NetInfo]
	err := json.Unmarshal(data, &apiMessage)
	return &apiMessage, err
}

func NewNetInfo(data []byte) (*NetInfo, error) {
	var netInfo NetInfo
	err := json.Unmarshal(data, &netInfo)
	return &netInfo, err
}

func (ni *NetInfo) Marshal() ([]byte, error) {
	result, err := json.Marshal(ni)
	return result, err
}

func (ni *NetInfo) MarshalIndent() ([]byte, error) {
	result, err := json.MarshalIndent(ni, "", "    ")
	return result, err
}

func (ni *NetInfo) String() string {
	j, err := ni.Marshal()
	if err != nil {
		fmt.Printf("Error encoding NetInfo to JSON: %s", err)
	}

	return string(j)
}

func (ni *NetInfo) StringIndent() string {
	j, err := ni.MarshalIndent()
	if err != nil {
		fmt.Printf("Error encoding NetInfo to JSON: %s", err)
	}

	return string(j)
}

type VersionInfo struct {
	Version    string `json:"version"`
	Update     int    `json:"update"`
	NewVersion string `json:"new_version"`
}
