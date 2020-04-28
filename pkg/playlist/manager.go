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
	filename  string                  `json:"-"`
	Playlists map[uuid.UUID]*playlist `json:"Playlists"`
}

func (pm *PlaylistManager) GetPlaylist(id uuid.UUID) (*playlist, error) {
	for k, v := range pm.Playlists {
		if id == k {
			return v, nil
		}
	}
	return nil, errors.New(fmt.Sprintf("playlist with id %s not found", id.String()))
}

func (pm *PlaylistManager) CreatePlaylist(name string, files []string) error {
	for _, v := range pm.Playlists {
		if v.Name == name {
			return errors.New(fmt.Sprintf("a playlist with the name %s already exists", name))
		}
	}

	id, err := uuid.NewRandom()
	if err != nil {
		return err
	}
	pm.Playlists[id] = &playlist{
		Name:  name,
		Files: files,
		Id:    id,
	}

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
			Playlists: map[uuid.UUID]*playlist{},
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
