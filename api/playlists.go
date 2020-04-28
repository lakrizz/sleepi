package api

import (
	"errors"

	"github.com/lakrizz/sleepi/pkg/playlist"
)

func (a *Api) GetPlaylists() ([]*playlist.Playlist, error) {
	if a.Playlists == nil {
		return nil, errors.New("there's no alarm manager, we cannot do anything :(")
	}
	return a.Playlists.Playlists, nil
}
