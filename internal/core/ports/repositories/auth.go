package repositories

type AuthRepository interface {
	GetToken(code string) (token string, refresh string, err error)
	RefreshToken(refreshToken string) (string, string, error)
	GetAuthURL() string
	IsTokenValid(token string) bool
}
