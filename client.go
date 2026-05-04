package bixgo

import (
	"net/http"
	"time"
)

var DefaultTimeout = time.Second * 15

type Client struct {
	baseURL    string
	auth       string
	httpClient *http.Client
}

func NewClient(baseURL, auth string) *Client {
	return &Client{
		baseURL: baseURL,
		auth:    auth,
		httpClient: &http.Client{
			Timeout: DefaultTimeout,
		},
	}
}

func NewClientWithTimeout(baseURL, auth string, timeout time.Duration) *Client {
	return &Client{
		baseURL: baseURL,
		auth:    auth,
		httpClient: &http.Client{
			Timeout: timeout,
		},
	}
}
