package authrepo

import (
	"context"

	wrapperutil "gmd3v.com/sgotify/internal/utils/wrapper_util"
	"golang.org/x/oauth2"
)

var utils = wrapperutil.NewWrapperUtil()

func NewWRPAuth() *WRPAuth {
	return &WRPAuth{}
}

type WRPAuth struct {
	AccessToken        string `json:"access_token"`
	TokenType          string `json:"token_type"`
	ExpiresIn          int    `json:"expires_in"`
	RefreshAccessToken string `json:"refresh_token"`
	Scope              string `json:"scope"`
}

// GetToken is a function that returns a token and a refresh token
func (w *WRPAuth) GetToken(code string) (string, string, error) {
	context := context.Background()
	println("Code", code)
	tok, err := utils.GetAuthenticator().Exchange(context, code)
	if err != nil {
		return "", "", err
	}
	return tok.AccessToken, tok.RefreshToken, nil
}

// RefreshToken is a function that returns a token and a refresh token
func (w *WRPAuth) RefreshToken(accessToken, refresh string) (string, string, error) {
	context := context.Background()
	tok, err := utils.GetAuthenticator().RefreshToken(context, &oauth2.Token{RefreshToken: refresh, AccessToken: accessToken})
	if err != nil {
		return "", "", err
	}
	return tok.AccessToken, tok.RefreshToken, nil
}

// Get URL for authentication
func (w *WRPAuth) GetAuthURL() string {
	return utils.GetAuthenticator().AuthURL("abc123")
}
