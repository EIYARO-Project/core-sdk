package api

type MessageData interface {
	NetInfo | Difficulty | any
}

type APiMessageSuccess[T MessageData] struct {
	Status string `json:"status"`
	Data   T      `json:"data"`
}

type APiMessageFail struct {
	Code      string `json:"code"`
	Message   string `json:"msg"`
	Status    string `json:"status"`
	Detail    string `json:"detail"`
	Temporary bool   `json:"temporary"`
}
