package bixgo

import (
	"encoding/json"
	"fmt"
)

// Response представляет собой структуру ответа от API Bitrix24 для одиночного элемента.
type Response[T any] struct {
	Result T `json:"result"`
	Error
}

// ListResponse представляет собой структуру ответа от API Bitrix24 для списков.
type ListResponse[T any] struct {
	Result []T `json:"result"`
	Total  int `json:"total"`
	Error
}

type Error struct {
	ErrorTitle       string `json:"error"`
	ErrorDescription string `json:"error_description"`
}

func (e Error) Error() string {
	return fmt.Sprintf("%s\n%s", e.ErrorTitle, e.ErrorDescription)
}

type rawResponse json.RawMessage

func (r rawResponse) BitrixError() error {
	var bitrixErr Error
	if err := json.Unmarshal(r, &bitrixErr); err != nil {
		return err
	}
	return &bitrixErr
}
