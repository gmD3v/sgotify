package repositories

type AuthRepository interface {
	GetToken(code string) (token string, refresh string, err error)
	RefreshToken(accessToken, refreshToken string) (string, string, error)
	GetAuthURL() string
}
