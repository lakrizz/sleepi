package playlist

import (
	"errors"
	"fmt"
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

func (q *Playlist) Remove(filename string) error {
	for k, v := range q.Files {
		if v == filename {
			q.Files = append(q.Files[:k], q.Files[k+1:]...)
			return nil
		}
	}

	return errors.New(fmt.Sprintf("file %s not found", filename))
}

func (q *Playlist) Length() int {
	return len(q.Files)
}

// an error is a missing file for example
func (q *Playlist) HasErrors() bool {
	for _, f := range q.Files {
		if _, err := os.Stat(f); err == os.ErrNotExist {
			return true
		}
	}
	return false
}
