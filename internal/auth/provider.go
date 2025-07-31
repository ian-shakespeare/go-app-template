package auth

import (
	"context"

	"golang.org/x/oauth2"
)

type Profile struct {
	Email   string `json:"email"`
	Picture string `json:"picture"`
}

type OAuth2Provider interface {
	AuthCodeURL(state string, opts ...oauth2.AuthCodeOption) string
	ExchangeAuthCode(ctx context.Context, code string) (Profile, error)
}
