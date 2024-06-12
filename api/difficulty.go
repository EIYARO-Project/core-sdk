package api

import (
	"encoding/json"
	"fmt"
)

type Difficulty struct {
	Hash        string `json:"hash"`
	BlockHeight uint64 `json:"height"`
	Bits        uint64 `json:"bits"`
	Difficulty  string `json:"difficulty"`
}

func NewApiMessageDifficulty(data []byte) (*APiMessageSuccess[Difficulty], error) {
	var apiMessage APiMessageSuccess[Difficulty]
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
