package api

import (
	"errors"

	"github.com/google/uuid"
	"github.com/lakrizz/sleepi/pkg/playlist"
)

func (a *Api) GetPlaylists() ([]*playlist.Playlist, error) {
	if a.Playlists == nil {
		return nil, errors.New("there's no alarm manager, we cannot do anything :(")
	}
	return a.Playlists.Playlists, nil
}

func (a *Api) GetPlaylist(id_s string) (*playlist.Playlist, error) {
	if a.Playlists == nil {
		return nil, errors.New("there's no alarm manager, we cannot do anything :(")
	}

	id, err := uuid.Parse(id_s)
	if err != nil {
		return nil, err
	}

	return a.Playlists.GetPlaylist(id)
}

func (a *Api) DeleteSongsFromPlaylist(playlist_id string, song_ids []string) error {
	if a.Playlists == nil {
		return errors.New("there's no alarm manager, we cannot do anything :(")
	}

	playlist, err := a.GetPlaylist(playlist_id)
	if err != nil {
		return err
	}

	for _, v := range song_ids {
		id, err := uuid.Parse(v)
		if err != nil {
			return err
		}
		if err = playlist.Remove(id); err != nil {
			return err
		}
	}

	err = a.Playlists.Save()
	return err
}
