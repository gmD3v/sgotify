package controlrepo

import (
	"context"

	"github.com/zmb3/spotify/v2"
	"gmd3v.com/sgotify/internal/core/domain"
	wrapperutil "gmd3v.com/sgotify/internal/utils/wrapper_util"
	"golang.org/x/oauth2"
)

var utils = wrapperutil.NewWrapperUtil()

type WrpControl struct {
}

func NewWrpControl() *WrpControl {
	return &WrpControl{}
}

func (s *WrpControl) PlaySong(token string, songID string) error {
	ctx := context.Background()
	client := spotify.New(utils.GetAuthenticator().Client(ctx, &oauth2.Token{AccessToken: token, TokenType: "Bearer"}))
	err := client.Play(ctx)
	if err != nil {
		return err
	}
	return nil
}

func (s *WrpControl) PauseSong(token string) error {
	ctx := context.Background()
	client := spotify.New(utils.GetAuthenticator().Client(ctx, &oauth2.Token{AccessToken: token}))
	err := client.Pause(ctx)
	if err != nil {
		return err
	}
	return nil
}

func (s *WrpControl) NextSong(token string) error {
	ctx := context.Background()
	client := spotify.New(utils.GetAuthenticator().Client(ctx, &oauth2.Token{AccessToken: token}))
	err := client.Next(ctx)
	if err != nil {
		return err
	}
	return nil
}

func (s *WrpControl) PreviousSong(token string) error {
	ctx := context.Background()
	client := spotify.New(utils.GetAuthenticator().Client(ctx, &oauth2.Token{AccessToken: token}))
	err := client.Previous(ctx)
	if err != nil {
		return err
	}
	return nil
}

func (s *WrpControl) GetSong(token string) (*domain.Song, error) {
	ctx := context.Background()
	client := spotify.New(utils.GetAuthenticator().Client(ctx, &oauth2.Token{AccessToken: token}))
	cp, err := client.PlayerCurrentlyPlaying(ctx)
	if err != nil {
		return nil, err
	}
	return &domain.Song{
		ID:          cp.Item.ID.String(),
		Name:        cp.Item.Name,
		Artist:      cp.Item.Artists[0].Name,
		Album:       cp.Item.Album.Name,
		ReleaseDate: cp.Item.Album.ReleaseDate,
		Length:      cp.Item.Duration / 1000,
		IsPlaying:   cp.Playing,
		URI:         string(cp.Item.URI),
	}, nil
}

func (s *WrpControl) GetCurrentlyPlaying(token string) (*spotify.CurrentlyPlaying, error) {
	ctx := context.Background()
	client := spotify.New(utils.GetAuthenticator().Client(ctx, &oauth2.Token{AccessToken: token}))
	cp, err := client.PlayerCurrentlyPlaying(ctx)
	if err != nil {
		return nil, err
	}
	return cp, nil
}

func (s *WrpControl) GetDevices(token string) ([]spotify.PlayerDevice, error) {
	ctx := context.Background()
	client := spotify.New(utils.GetAuthenticator().Client(ctx, &oauth2.Token{AccessToken: token}))
	devices, err := client.PlayerDevices(ctx)
	if err != nil {
		return nil, err
	}
	return devices, nil
}
