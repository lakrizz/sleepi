package manager

import (
	"encoding/json"
	"errors"
	"io/ioutil"
	"os"
	"path"
	"path/filepath"

	"github.com/adrg/xdg"
	"krizz.org/sleepi/pkg/playlist"
)

type PlaylistManager struct {
	Playlists []*playlist.Playlist
}

var playlists_file_name string = "playlists.json"

func GetPlaylistManager() (*PlaylistManager, error) {
	folder := path.Join(xdg.UserDirs.Documents, "sleepi")
	if _, err := os.Stat(folder); os.IsNotExist(err) {
		err := os.Mkdir(folder, 0777)
		if err != nil {
			return nil, err
		}
	}

	pl, err := searchPlaylistFile(folder)
	if err != nil {
		return &PlaylistManager{Playlists: make([]*playlist.Playlist, 0)}, nil
	}
	return pl, nil
}

func (pm *PlaylistManager) AddPlaylist(playlist *playlist.Playlist) error {
	if !pm.isInList(playlist) {
		pm.Playlists = append(pm.Playlists, playlist)
	}
	return nil
}

func (pm *PlaylistManager) isInList(playlist *playlist.Playlist) bool {
	for _, v := range pm.Playlists {
		if v.Id == playlist.Id {
			return true
		}
	}
	return false
}

func (pm *PlaylistManager) Save() error {
	return nil
}

func searchPlaylistFile(folder string) (*PlaylistManager, error) {
	if _, err := os.Stat(folder); os.IsNotExist(err) {
		return nil, err
	}

	files, err := ioutil.ReadDir(folder)
	if err != nil {
		return nil, err
	}

	for _, v := range files {
		if v.Name() == playlists_file_name {
			// found it, read it!
			dat, err := ioutil.ReadFile(filepath.Join(folder, v.Name()))
			if err != nil {
				return nil, err
			}

			var pl_manager *PlaylistManager
			err = json.Unmarshal(dat, &pl_manager)
			if err != nil {
				return nil, err
			}

			return pl_manager, nil
		}
	}

	return nil, errors.New("no config found")
}
