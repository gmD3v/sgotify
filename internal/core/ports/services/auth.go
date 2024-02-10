package services

type AuthService interface {
	Login(code string) error
	Refresh(accessToken, refresh string) error
	Logout() error
	RedirectUrl() string
}
