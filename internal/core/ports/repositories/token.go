package repositories

type TokenRepository interface {
	SaveToken(accessToken, refreshToken string) error
	GetSavedToken() (string, string, error)
	DeleteToken() error
}
