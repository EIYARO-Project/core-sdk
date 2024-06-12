package resources

import (
	"encoding/json"
	"fmt"
)

type Account struct {
	ID         string   `json:"id"`
	Alias      string   `json:"alias"`
	XPubs      []string `json:"xpubs"`
	Quorum     uint64   `json:"quorum"`
	KeyIndex   uint64   `json:"key_index"`
	DeriveRule uint64   `json:"derive_rule"`
}

func (a *Account) Marshal() ([]byte, error) {
	result, err := json.Marshal(a)
	return result, err
}

func (a *Account) MarshalIndent() ([]byte, error) {
	result, err := json.MarshalIndent(a, "", "    ")
	return result, err
}

func (a *Account) String() string {
	j, err := a.Marshal()
	if err != nil {
		fmt.Printf("Error encoding Account to JSON: %s", err)
	}

	return string(j)
}

func (a *Account) StringIndent() string {
	j, err := a.MarshalIndent()
	if err != nil {
		fmt.Printf("Error encoding Account to JSON: %s", err)
	}

	return string(j)
}
