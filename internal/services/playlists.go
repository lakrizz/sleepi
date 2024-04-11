package services

import (
	"encoding/json"
	"errors"
	"fmt"
	"os"
	"slices"

	"github.com/google/uuid"

	"github.com/lakrizz/sleepi/config"
	"github.com/lakrizz/sleepi/pkg/playlist"
)

type PlaylistService struct {
	cfg       *config.Config
	Playlists []*playlist.Playlist
}

func NewPlaylistService(cfg *config.Config) (*PlaylistService, error) {
	dat, err := os.ReadFile(cfg.PlaylistsFileName)
	if err != nil {
		return &PlaylistService{Playlists: make([]*playlist.Playlist, 0)}, nil
	}

	ps := &PlaylistService{cfg: cfg}
	err = json.Unmarshal(dat, &ps)
	if err != nil {
		return nil, err
	}

	return ps, nil
}

func (pm *PlaylistService) AddPlaylist(plist *playlist.Playlist) error {
	if !slices.ContainsFunc(pm.Playlists, func(e *playlist.Playlist) bool { return e.ID == plist.ID }) {
		pm.Playlists = append(pm.Playlists, plist)
		return pm.Save()
	}

	if plist == nil {
		return errors.New("you are trying to add an empty playlist")
	}

	return fmt.Errorf("a playlist with the id %s already exists - is this a coincidence?", plist.ID)
}

func (pm *PlaylistService) Save() error {
	for _, v := range pm.Playlists {
		if !v.Valid() {
			return fmt.Errorf("can't save playlists. playlist %s is not valid", v.Name)
		}
	}

	dat, err := json.Marshal(pm)
	if err != nil {
		return err
	}

	err = os.WriteFile(pm.cfg.PlaylistsFileName, dat, 0777)
	return err
}

func (pm *PlaylistService) GetPlaylistById(id uuid.UUID) (*playlist.Playlist, error) {
	if id == uuid.Nil {
		return nil, errors.New("playlist id can't be null")
	}

	for _, v := range pm.Playlists {
		if v.ID == id {
			return v, nil
		}
	}

	return nil, fmt.Errorf("no playlist found for id %s", id.String())
}

func (pl *PlaylistService) GetAllPlaylists() ([]*playlist.Playlist, error) {
	return pl.Playlists, nil
}
