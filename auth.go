package bixgo

import "time"

type ClientAuth struct {
	authToken    string
	refreshToken string
	authAt       time.Time
	expiresIn    time.Duration
}

func NewClientAuth(
	authToken,
	refreshToken string,
	authAt time.Time,
	expiresIn time.Duration,
) *ClientAuth {
	return &ClientAuth{
		authToken:    authToken,
		refreshToken: refreshToken,
		authAt:       authAt,
		expiresIn:    expiresIn,
	}
}

func (c *ClientAuth) Refresh() {
	if !c.IsExpired() {
		return
	}
}

func (c *ClientAuth) IsExpired() bool {
	if c.authAt.Add(c.expiresIn).Before(time.Now()) {
		return true
	}
	return false
}
