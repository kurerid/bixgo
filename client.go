package bixgo

import (
	"net/http"
	"time"
)

var DefaultTimeout = time.Second * 15

type Client struct {
	baseURL    string
	auth       *ClientAuth
	httpClient *http.Client
}

func NewClient(baseURL string, auth *ClientAuth) *Client {
	return &Client{
		baseURL: baseURL,
		auth:    auth,
		httpClient: &http.Client{
			Timeout: DefaultTimeout,
		},
	}
}

func NewClientWithTimeout(baseURL string, auth *ClientAuth, timeout time.Duration) *Client {
	return &Client{
		baseURL: baseURL,
		auth:    auth,
		httpClient: &http.Client{
			Timeout: timeout,
		},
	}
}
