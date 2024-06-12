package api

import (
	"encoding/json"
	"fmt"
)

type BlockRequest struct {
	BlockHeight uint64 `json:"block_height"`
	BlockHash   string `json:"block_hash"`
}

func (br *BlockRequest) String() string {
	j, err := br.Marshal()
	if err != nil {
		fmt.Printf("Error encoding NetInfo to JSON: %s", err)
	}

	return string(j)
}

func (br *BlockRequest) Marshal() ([]byte, error) {
	result, err := json.Marshal(br)
	return result, err
}
