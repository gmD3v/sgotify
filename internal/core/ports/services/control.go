package services

import "gmd3v.com/sgotify/internal/core/domain"

type ControlService interface {
	PlaySong(songID string) error
	PauseSong() error
	NextSong() error
	PreviousSong() error
	GetSong() (domain.Song, error)
}
