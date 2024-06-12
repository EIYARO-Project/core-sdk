package api

import "github.com/EIYARO-Project/core-sdk/api/resources"

type MessageData interface {
	NetInfo |
		Difficulty |
		GetWorkJson |
		resources.AccessToken |
		resources.Account |
		any
}

type APiMessageSuccessObject[T MessageData] struct {
	Status string `json:"status"`
	Data   T      `json:"data"`
}

type APiMessageSuccessArray[T MessageData] struct {
	Status string `json:"status"`
	Data   []T    `json:"data"`
}

type APiMessageFail struct {
	Code      string `json:"code"`
	Message   string `json:"msg"`
	Status    string `json:"status"`
	Detail    string `json:"detail"`
	Temporary bool   `json:"temporary"`
}
