package api

import (
	"encoding/json"
	"fmt"
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
