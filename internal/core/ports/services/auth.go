package services

type AuthService interface {
	Login(code string) error
	Refresh(refresh string) error
	Logout() error
	RedirectUrl() string
}
