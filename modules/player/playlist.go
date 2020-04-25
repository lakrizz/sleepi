package player

import (
	"encoding/json"
	"errors"
	"io/ioutil"
	"os"
)

type Playlist struct {
	Name  string   `json:"Name"` // kinda like the name of a playlist
	Files []string `json:"Files"`
}

func LoadPlaylist(filename string) (*Playlist, error) {
	if _, err := os.Stat(filename); err == os.ErrNotExist {
		return nil, errors.New("playlist file does not exist, could not load, eh?")
	}

	dat, err := ioutil.ReadFile(filename)
	if err != nil {
		return nil, err
	}

	var playlist *Playlist
	err = json.Unmarshal(dat, &playlist)
	if err != nil {
		return nil, err
	}
	return playlist, nil
}

func (q *Playlist) Add(filename string) error {
	_, err := os.Stat(filename)
	if os.IsNotExist(err) {
		return errors.New(err.Error())
	}
	q.Files = append(q.Files, filename)
	return nil
}

func (q *Playlist) Pop() (string, error) {
	if len(q.Files) > 0 { // theres something in the list
		return q.Files[0], nil
	}
	return "", errors.New("Playlist is empty")
}

func (q *Playlist) RemoveAt(idx int) error {
	if len(q.Files) < idx {
		return errors.New("index out of range")
	}
	return nil
}

func (q *Playlist) Length() int {
	return len(q.Files)
}
