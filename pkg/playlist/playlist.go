package playlist

import (
	"errors"
	"os"

	"github.com/google/uuid"
)

type Playlist struct {
	Id    uuid.UUID `json:"Id"`
	Name  string    `json:"Name"` // kinda like the name of a playlist
	Files []string  `json:"Files"`
}

func (q *Playlist) AddSong(filename string) error {
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
	return "", errors.New("playlist is empty")
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
