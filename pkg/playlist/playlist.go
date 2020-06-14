package playlist

import (
	"errors"
	"fmt"

	"github.com/google/uuid"
)

type Playlist struct {
	Id    uuid.UUID   `json:"Id"`
	Name  string      `json:"Name"` // kinda like the name of a playlist
	Songs []uuid.UUID `json:"Songs"`
}

func (q *Playlist) AddSong(song_id uuid.UUID) error {
	q.Songs = append(q.Songs, song_id)
	return nil
}

func (q *Playlist) Pop() (uuid.UUID, error) {
	if len(q.Songs) > 0 { // theres something in the list
		return q.Songs[0], nil
	}
	return uuid.Nil, errors.New("playlist is empty")
}

func (q *Playlist) Remove(id uuid.UUID) error {
	for k, v := range q.Songs {
		if v == id {
			q.Songs = append(q.Songs[:k], q.Songs[k+1:]...)
			return nil
		}
	}

	return errors.New(fmt.Sprintf("file %s not found", id))
}

func (q *Playlist) Length() int {
	return len(q.Songs)
}

// an error is a missing file for example
func (q *Playlist) HasErrors() bool {
	// for _, f := range q.Songs {
	// #TODO(@krizzle): use the librarymanager to check if the file behind the id is existing
	// }
	return false
}
