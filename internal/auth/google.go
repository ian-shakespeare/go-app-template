package auth

import (
	"context"
	"encoding/json"

	"github.com/ian-shakespeare/go-app-template/internal/env"
	"golang.org/x/oauth2"
	"golang.org/x/oauth2/google"
)

const googleUserInfoURL = "https://www.googleapis.com/oauth2/v2/userinfo"

type GoogleOAuth2 struct {
	oauth2.Config
}

func NewGoogleOAuth2(clientID, clientSecret string) *GoogleOAuth2 {
	return &GoogleOAuth2{
		oauth2.Config{
			ClientID:     clientID,
			ClientSecret: clientSecret,
			Endpoint:     google.Endpoint,
			RedirectURL:  env.Fallback("BASE_URL", "http://localhost:8000") + "/auth/google/callback",
		},
	}
}

func (g *GoogleOAuth2) ExchangeAuthCode(ctx context.Context, code string) (Profile, error) {
	var p Profile

	token, err := g.Exchange(ctx, code)
	if err != nil {
		return p, err
	}

	client := g.Client(ctx, token)
	res, err := client.Get(googleUserInfoURL)
	if err != nil {
		return p, err
	}
	defer res.Body.Close()

	decoder := json.NewDecoder(res.Body)
	err = decoder.Decode(&p)
	return p, err
}
