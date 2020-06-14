package playlist

import (
	"encoding/json"
	"errors"
	"fmt"
	"io/ioutil"

	"github.com/adrg/xdg"
	"github.com/google/uuid"
)

type PlaylistManager struct {
	filename  string      `json:"-"`
	Playlists []*Playlist `json:"Playlists"`
}

func (pm *PlaylistManager) GetPlaylist(id uuid.UUID) (*Playlist, error) {
	for _, v := range pm.Playlists {
		if id == v.Id {
			return v, nil
		}
	}
	return nil, errors.New(fmt.Sprintf("playlist with id %s not found", id.String()))
}

func (pm *PlaylistManager) CreatePlaylist(name string, files []uuid.UUID) error {
	for _, v := range pm.Playlists {
		if v.Name == name {
			return errors.New(fmt.Sprintf("a playlist with the name %s already exists", name))
		}
	}

	id, err := uuid.NewRandom()
	if err != nil {
		return err
	}
	pm.Playlists = append(pm.Playlists, &Playlist{
		Name:  name,
		Songs: files,
		Id:    id,
	})

	return pm.Save()
}

func LoadConfig() (*PlaylistManager, error) {
	filename, err := xdg.ConfigFile("sleepi/playlists.json")
	if err != nil {
		return nil, err
	}

	dat, err := ioutil.ReadFile(filename)
	if err != nil {
		return &PlaylistManager{
			Playlists: []*Playlist{},
			filename:  filename,
		}, nil
	}

	var c *PlaylistManager
	err = json.Unmarshal(dat, &c)
	if err != nil {
		return nil, err
	}
	c.filename = filename
	return c, nil
}

func (c *PlaylistManager) Save() error {
	dat, err := json.Marshal(c)
	if err != nil {
		return err
	}
	fmt.Println(string(dat))
	err = ioutil.WriteFile(c.filename, dat, 0755)
	return err
}
