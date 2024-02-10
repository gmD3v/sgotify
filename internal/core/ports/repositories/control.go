package repositories

import "gmd3v.com/sgotify/internal/core/domain"

type ControlRepository interface {
	PlaySong(token string, songID string) error
	PauseSong(token string) error
	NextSong(token string) error
	PreviousSong(token string) error
	GetSong(token string) (*domain.Song, error)
}
