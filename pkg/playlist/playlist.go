package playlist

import (
	"errors"
	"math/rand"
	"time"

	"github.com/gofrs/uuid"
	"krizz.org/sleepi/pkg/library"
)

// a playlist is nothing more than a loose collection of files that reside in "the" library
// we are NOT expecting a library-reference in the struct but more so a list of files
// thus the playlist is NOT managed from within and is not aware of itself

type Playlist struct {
	Name  string
	Files []*library.File
}

func NewPlaylist(name string) (*Playlist, error) {
	p := &Playlist{Name: name, Files: make([]*library.File, 0)}
	return p, nil
}

func (p *Playlist) contains(id uuid.UUID) bool {
	for _, v := range p.Files {
		if v == id {
			return true
		}
	}
	return false

}

func (p *Playlist) Add(file *library.File) error {
	if file == nil {
		return errors.New("cannot add nil-file")
	}

	if p.contains(file.Id) {
		return errors.New("this file is already part of this playlist")
	}

	p.Files = append(p.files, file)
	return nil
}

func (p *Playlist) Remove(id uuid.UUID) error {
	if id == uuid.NullUUID {
		return errors.New("id is nil, i don't want to remove this")
	}

	if !p.contains(id) {
		return errors.New("this id is not part of this playlist, i can't remove it")
	}

	id := -1
	for i, v := range p.Files {
		if v.Id == id {
			id = i
			break
		}
	}

	p.Files = append(p.Files[:id], p.Files[id+1:]...) // removes item at position i
	return nil
}

func (p *Playlist) GetFileByIndex(index int) (*library.File, error) {
	if len(p.Files) == 0 {
		return nil, errors.New("there are no files here, can't return anything")
	}

	if index >= len(p.Files) {
		return nil, errors.New("index out of bounds", index)
	}

	if p.Files[index] == nil {
		return nil, errors.New("index was found but is still nil / somethings terribly fucked here", index)
	}

	return p.Files[index], nil
}

func (p *Playlist) GetRandomFile() (*library.File, error) {
	if len(p.Files) == 0 {
		return nil, errors.New("there are no files here, can't return anything")
	}

	rand.Seed(time.Now().UnixMicro())
	i := rand.Intn(len(p.Files))
	return p.Files[i], nil
}
