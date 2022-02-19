package playlist

import (
	"errors"
	"fmt"
	"math/rand"
	"time"

	"github.com/google/uuid"
	"krizz.org/sleepi/pkg/library"
)

// a playlist is nothing more than a loose collection of files that reside in "the" library
// we are NOT expecting a library-reference in the struct but more so a list of files
// thus the playlist is NOT managed from within and is not aware of itself

type Playlist struct {
	Id    uuid.UUID
	Name  string
	Files []*library.File
}

func NewPlaylist(name string) (*Playlist, error) {
	p := &Playlist{Name: name, Files: make([]*library.File, 0)}
	return p, nil
}

func (p *Playlist) contains(id uuid.UUID) bool {
	for _, v := range p.Files {
		if v.Id == id {
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

	p.Files = append(p.Files, file)
	return nil
}

func (p *Playlist) Remove(id uuid.UUID) error {
	if id == uuid.Nil {
		return errors.New("id is nil, i don't want to remove this")
	}

	if !p.contains(id) {
		return errors.New("this id is not part of this playlist, i can't remove it")
	}

	for i, v := range p.Files {
		if v.Id == id {
			p.Files = append(p.Files[:i], p.Files[i+1:]...) // removes item at position i
			break
		}
	}

	return nil
}

func (p *Playlist) GetFileByIndex(index int) (*library.File, error) {
	if len(p.Files) == 0 {
		return nil, errors.New("there are no files here, can't return anything")
	}

	if index >= len(p.Files) {
		return nil, fmt.Errorf("index out of bounds: %v", index)
	}

	if p.Files[index] == nil {
		return nil, fmt.Errorf("index was found but is still nil / somethings terribly fucked here: %v", index)
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
