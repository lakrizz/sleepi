package manager

import (
	"encoding/json"
	"errors"
	"fmt"
	"io/ioutil"
	"os"
	"path"
	"path/filepath"

	"github.com/adrg/xdg"
	"github.com/google/uuid"
	"krizz.org/sleepi/pkg/playlist"
)

type PlaylistManager struct {
	Playlists []*playlist.Playlist
}

var playlists_file_name string = "playlists.json"

func getPlaylistManager() (*PlaylistManager, error) {
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
		return pm.Save()
	}

	if playlist == nil {
		return errors.New("you are trying to add an empty playlist")
	}

	return fmt.Errorf("a playlist with the id %v already exists - is this a coincidence?", playlist.Id.String())
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
	for _, v := range pm.Playlists {
		if !v.Valid() {
			return fmt.Errorf("can't save playlists. playlist %s is not valid", v.Name)
		}
	}

	filename := path.Join(xdg.UserDirs.Documents, "sleepi", playlists_file_name)

	dat, err := json.Marshal(pm)
	if err != nil {
		return err
	}

	err = ioutil.WriteFile(filename, dat, 07777)
	return err
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

func (pm *PlaylistManager) GetPlaylistById(id uuid.UUID) (*playlist.Playlist, error) {
	if id == uuid.Nil {
		return nil, errors.New("playlist id can't be null")
	}

	for _, v := range pm.Playlists {
		if v.Id == id {
			return v, nil
		}
	}

	return nil, fmt.Errorf("no playlist found for id %s", id.String())
}

func (pm *PlaylistManager) UpdatePlaylist(playlist *playlist.Playlist) error {
	if pm.isInList(playlist) {
		for _, v := range pm.Playlists {
			if v.Id == playlist.Id {
				v.Files = playlist.Files
				break
			}
		}
		return pm.Save()
	}

	if playlist == nil {
		return errors.New("you are trying to update an empty playlist")
	}

	return fmt.Errorf("a playlist with the id %v does not exist", playlist.Id.String())
}
