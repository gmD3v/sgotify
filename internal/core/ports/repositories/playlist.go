package repositories

import "gmd3v.com/sgotify/internal/core/domain"

type PlaylistRepository interface {
	GetPlaylists() ([]domain.Playlist, error)
	GetPlaylist(id string) ([]domain.Playlist, error)
	CreatePlaylist(name string, isPublic bool, isCollaborative bool) (domain.Playlist, error)
	AddSongToPlaylist(playlistID string, songID string) error
	RemoveSongFromPlaylist(playlistID string, songID string) error
	DeletePlaylist(id string) error
	UpdatePlaylist(id string, name string, isPublic bool, isCollaborative bool) error
}
