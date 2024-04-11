package playlist

import (
	"errors"
	"math/rand"

	"github.com/google/uuid"

	"github.com/lakrizz/sleepi/pkg/models"
)

// a playlist is nothing more than a loose collection of files that reside in "the" library
// we are NOT expecting a library-reference in the struct but more so a list of files
// thus the playlist is NOT managed from within and is not aware of itself

type Playlist struct {
	ID    uuid.UUID
	Name  string
	Files []*models.File
}

var (
	errIDIsNill              = errors.New("file id is nil")
	errFileNotInPlaylist     = errors.New("this id is not part of this playlist, i can't remove it")
	errFileAlreadyInPlaylist = errors.New("this file is already part of this playlist")
	errPlaylistEmpty         = errors.New("playlist is empty")
	errOutOfBounds           = errors.New("index out of bounds")
	errFileIsNil             = errors.New("file is found but nil")
)

func NewPlaylist(name string) (*Playlist, error) {
	p := &Playlist{Name: name, Files: make([]*models.File, 0), ID: uuid.New()}
	return p, nil
}

func (p *Playlist) ContainsFile(id uuid.UUID) bool {
	for _, v := range p.Files {
		if v.ID == id {
			return true
		}
	}

	return false
}

func (p *Playlist) AddFile(file *models.File) error {
	if file == nil {
		return errIDIsNill
	}

	if p.ContainsFile(file.ID) {
		return errFileAlreadyInPlaylist
	}

	p.Files = append(p.Files, file)

	return nil
}

func (p *Playlist) RemoveFile(id uuid.UUID) error {
	if id == uuid.Nil {
		return errIDIsNill
	}

	if !p.ContainsFile(id) {
		return errFileNotInPlaylist
	}

	for i, v := range p.Files {
		if v.ID == id {
			p.Files = append(p.Files[:i], p.Files[i+1:]...) // removes item at position i
			break
		}
	}

	return nil
}

func (p *Playlist) ClearFiles() error {
	p.Files = make([]*models.File, 0)
	return nil
}

func (p *Playlist) GetFileByIndex(index int) (*models.File, error) {
	if len(p.Files) == 0 {
		return nil, errPlaylistEmpty
	}

	if index >= len(p.Files) {
		return nil, errOutOfBounds
	}

	if p.Files[index] == nil {
		return nil, errFileIsNil
	}

	return p.Files[index], nil
}

func (p *Playlist) GetRandomFile() (*models.File, error) {
	if len(p.Files) == 0 {
		return nil, errPlaylistEmpty
	}

	return p.Files[rand.Intn(len(p.Files))], nil
}

func (p *Playlist) Valid() bool {
	return len(p.Files) > 0 && p.Name != ""
}
