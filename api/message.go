package api

import "encoding/json"

type MessageData interface {
	NetInfo | any
}

type APiMessageSuccess[T MessageData] struct {
	Status string `json:"status"`
	Data   T      `json:"data"`
}

func NewApiMessageNetInfo(data []byte) (*APiMessageSuccess[NetInfo], error) {
	var apiMessage APiMessageSuccess[NetInfo]
	err := json.Unmarshal(data, &apiMessage)
	return &apiMessage, err
}

type APiMessageFail struct {
	Code      string `json:"code"`
	Message   string `json:"msg"`
	Status    string `json:"status"`
	Detail    string `json:"detail"`
	Temporary bool   `json:"temporary"`
}
