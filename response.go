package bixgo

import (
	"encoding/json"
	"fmt"
)

// ResponseMarker — marker interface for all Bitrix24 response types
type BixgoResponseMarker interface {
	isBitrixResponse()
}

// Response представляет собой структуру ответа от API Bitrix24 для одиночного элемента.
type Response[T any] struct {
	Result T `json:"result"`
	Error
}

func (r Response[T]) isBitrixResponse() {}

// ListResponse представляет собой структуру ответа от API Bitrix24 для списков.
type ListResponse[T any] struct {
	Result T   `json:"result"`
	Total  int `json:"total"`
	Error
}

func (l ListResponse[T]) isBitrixResponse() {}

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
	if bitrixErr == (Error{}) {
		return nil
	}
	return bitrixErr
}
