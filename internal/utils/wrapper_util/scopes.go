package wrapperutil

import (
	spotifyauth "github.com/zmb3/spotify/v2/auth"
)

var authScopes = spotifyauth.WithScopes(spotifyauth.ScopeUserReadPrivate, spotifyauth.ScopeUserModifyPlaybackState, spotifyauth.ScopeUserReadCurrentlyPlaying, spotifyauth.ScopeUserReadPlaybackState, spotifyauth.ScopeUserLibraryRead, spotifyauth.ScopeUserLibraryModify, spotifyauth.ScopeUserFollowRead, spotifyauth.ScopeUserFollowModify, spotifyauth.ScopePlaylistReadPrivate, spotifyauth.ScopePlaylistModifyPublic, spotifyauth.ScopePlaylistModifyPrivate, spotifyauth.ScopePlaylistReadCollaborative, spotifyauth.ScopeUserTopRead, spotifyauth.ScopeUserReadRecentlyPlayed, spotifyauth.ScopeUserReadEmail)
