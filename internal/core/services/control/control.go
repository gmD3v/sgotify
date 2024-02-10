package control

import (
	"gmd3v.com/sgotify/internal/core/domain"
	"gmd3v.com/sgotify/internal/core/ports/repositories"
)

type ControlService struct {
	repo      repositories.ControlRepository
	tokenRepo repositories.TokenRepository
}

func NewService(repo repositories.ControlRepository, token repositories.TokenRepository) *ControlService {
	return &ControlService{repo: repo, tokenRepo: token}
}

func (s *ControlService) PlaySong(songID string) error {
	token, _, err := s.tokenRepo.GetSavedToken()
	if err != nil {
		return err
	}
	return s.repo.PlaySong(token, songID)
}

func (s *ControlService) PauseSong() error {
	token, _, err := s.tokenRepo.GetSavedToken()
	if err != nil {
		return err
	}
	return s.repo.PauseSong(token)

}

func (s *ControlService) NextSong() error {
	token, _, err := s.tokenRepo.GetSavedToken()
	if err != nil {
		return err
	}
	return s.repo.NextSong(token)
}

func (s *ControlService) PreviousSong() error {
	token, _, err := s.tokenRepo.GetSavedToken()
	if err != nil {
		return err
	}
	return s.repo.PreviousSong(token)
}

func (s *ControlService) GetSong() (domain.Song, error) {
	token, _, err := s.tokenRepo.GetSavedToken()
	if err != nil {
		return domain.Song{}, err

	}
	_, err = s.repo.GetSong(token)
	if err != nil {
		return domain.Song{}, err
	}
	return domain.Song{}, nil
}
