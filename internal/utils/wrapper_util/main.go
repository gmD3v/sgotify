package wrapperutil

import spotifyauth "github.com/zmb3/spotify/v2/auth"

var auth *spotifyauth.Authenticator

type WrapperUtil struct {
}

func NewWrapperUtil() *WrapperUtil {
	auth = spotifyauth.New(spotifyauth.WithRedirectURL("http://localhost:8080/callback"), authScopes)
	return &WrapperUtil{}
}

func (w *WrapperUtil) GetAuthenticator() *spotifyauth.Authenticator {
	auth = spotifyauth.New(spotifyauth.WithRedirectURL("http://localhost:8080/callback"), authScopes)
	return auth
}

func (w *WrapperUtil) GetAuthScopes() spotifyauth.AuthenticatorOption {
	return authScopes
}
