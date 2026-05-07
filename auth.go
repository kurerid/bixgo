package bixgo

import (
	"encoding/json"
	"fmt"
	"io"
	"net/http"
	"time"
)

type ClientAuth struct {
	authToken    string
	refreshToken string
	expiresAt    time.Time
	clientId     string
	clientSecret string
}

type refreshResponse struct {
	AccessToken  string `json:"access_token"`
	ExpiresIn    int    `json:"expires_in"`
	RefreshToken string `json:"refresh_token"`
}

const refreshEndpoint = "https://oauth.bitrix24.tech/oauth/token/?grant_type=refresh_token&client_id=%s&client_secret=%s&refresh_token=%s"

func NewClientAuth(
	authToken,
	refreshToken string,
	expiresAt time.Time,
	clientId, clientSecret string,
) *ClientAuth {
	return &ClientAuth{
		authToken:    authToken,
		refreshToken: refreshToken,
		expiresAt:    expiresAt,
		clientId:     clientId,
		clientSecret: clientSecret,
	}
}

func (c *ClientAuth) Refresh() error {
	if !c.IsExpired() {
		return nil
	}
	response, err := http.DefaultClient.Post(
		c.refreshEndpoint(),
		"application/json",
		nil,
	)
	if err != nil {
		return err
	}

	if response.StatusCode != http.StatusOK {
		bodyBytes, _ := io.ReadAll(response.Body)
		return fmt.Errorf(
			"failed to refresh token. status code: %d, body: %s",
			response.StatusCode,
			string(bodyBytes),
		)
	}

	var decodedResponse refreshResponse
	err = json.NewDecoder(response.Body).Decode(&decodedResponse)
	if err != nil {
		return err
	}
	c.authToken = decodedResponse.AccessToken
	c.refreshToken = decodedResponse.RefreshToken
	c.expiresAt = time.Unix(int64(decodedResponse.ExpiresIn), 0)
	return nil
}

func (c *ClientAuth) IsExpired() bool {
	if c.expiresAt.Before(time.Now().Add(-time.Minute)) {
		return true
	}
	return false
}

func (c *ClientAuth) refreshEndpoint() string {
	return fmt.Sprintf(refreshEndpoint, c.clientId, c.clientSecret, c.refreshToken)
}
