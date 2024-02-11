package auth

import "gmd3v.com/sgotify/internal/core/ports/repositories"

type service struct {
	repo      repositories.AuthRepository
	tokenrepo repositories.TokenRepository
}

func NewService(repo repositories.AuthRepository, tokenrepo repositories.TokenRepository) *service {
	return &service{repo: repo, tokenrepo: tokenrepo}
}

func (s *service) Login(code string) error {
	accesTk, refreshTk, err := s.repo.GetToken(code)
	if err != nil {
		return err
	}
	err = s.tokenrepo.SaveToken(accesTk, refreshTk)
	if err != nil {
		return err
	}
	return nil
}

func (s *service) Refresh(refresh string) error {
	accessTk, refreTk, err := s.repo.RefreshToken(refresh)
	if err != nil {
		return err
	}
	err = s.tokenrepo.SaveToken(accessTk, refreTk)
	if err != nil {
		return err
	}
	return nil
}

func (s *service) Logout() error {
	return s.tokenrepo.DeleteToken()
}

func (s *service) RedirectUrl() string {
	return s.repo.GetAuthURL()
}
