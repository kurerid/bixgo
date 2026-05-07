package bixgo

import (
	"bytes"
	"context"
	"encoding/json"
	"errors"
	"fmt"
	"io"
	"net/http"
)

/*
Call

	Декодирует ответ в result.
	result должен быть указателем на одну из структур Response или ListResponse.
	Позволяет вызывать методы, возвращающие как одиночный ответ так и ответ со списком
*/
func (c *Client) Call(ctx context.Context, method string, params Params, result any) error {
	if result == nil {
		return errors.New("result must be not nil")
	}

	if !isBixgoResponseStruct(result) {
		return errors.New("result must be a pointer to bixgo.Response or bixgo.ListResponse")
	}

	rawResponse, err := c.callRaw(ctx, method, params)
	if err != nil {
		return err
	}

	return json.Unmarshal(rawResponse, result)
}

func (c *Client) callRaw(
	ctx context.Context,
	method string,
	params Params,
) (rawResponse, error) {
	if c.auth.IsExpired() {
		err := c.auth.Refresh()
		if err != nil {
			return nil, err
		}
	}

	endpoint := fmt.Sprintf("%s/rest/%s.json?auth=%s", c.baseURL, method, c.auth.authToken)
	var body io.Reader
	if params != nil {
		b, err := json.Marshal(params)
		if err != nil {
			return nil, fmt.Errorf("marshal params: %w", err)
		}
		body = bytes.NewBuffer(b)
	}

	req, err := http.NewRequestWithContext(ctx, http.MethodPost, endpoint, body)
	if err != nil {
		return nil, fmt.Errorf("create request: %w", err)
	}

	req.Header.Set("Content-Type", "application/json")

	resp, err := c.httpClient.Do(req)
	if err != nil {
		return nil, fmt.Errorf("do request: %w", err)
	}
	defer resp.Body.Close()

	responseBytes, err := io.ReadAll(resp.Body)
	if err != nil {
		return nil, fmt.Errorf("failed to read response: %w", err)
	}
	// можно добавить проверку HTTP статуса
	if resp.StatusCode != http.StatusOK {
		return nil, fmt.Errorf("http error %d: %s", resp.StatusCode, string(responseBytes))
	}

	var raw rawResponse = responseBytes

	// обработка ошибки Bitrix
	if bitrixErr := raw.BitrixError(); bitrixErr != nil {
		return nil, fmt.Errorf("bitrix error: %s", bitrixErr.Error())
	}

	return raw, nil
}

func isBixgoResponseStruct(result any) bool {
	_, ok := result.(BixgoResponseMarker)
	return ok
}
