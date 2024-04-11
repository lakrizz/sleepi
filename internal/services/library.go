package services

import (
	"encoding/json"
	"errors"
	"log"
	"os"
	"path"

	"github.com/google/uuid"

	"github.com/lakrizz/sleepi/config"
	"github.com/lakrizz/sleepi/pkg/models"
)

type LibraryService struct {
	config *config.Config

	Files map[uuid.UUID]*models.File
}

var (
	errFileNotFound = errors.New("file not found")
)

// GetLibraryService creates a library which is nothing but a bunch of files
// stored in a collection, more like a fs, it's an abstraction layer
// for any file consumer to access files without having the need
// to know about the location
func NewLibraryService(cfg *config.Config) (*LibraryService, error) {
	dat, err := os.ReadFile(cfg.LibraryFileName)
	if err != nil {
		return nil, err
	}

	lib := &LibraryService{config: cfg}
	err = json.Unmarshal(dat, &lib)
	if err != nil {
		return nil, err
	}

	return lib, nil
}

func (l *LibraryService) AddFile(data []byte, name string) error {
	targetFile := path.Join(l.config.MusicFolderFileName, name)

	err := os.WriteFile(targetFile, data, 0777)
	if err != nil {
		return err
	}

	f := &models.File{Path: targetFile, ID: uuid.New()}
	l.Files[f.ID] = f // add to the map

	return l.Save()
}

func (l *LibraryService) RemoveByID(id uuid.UUID) error {
	if _, ok := l.Files[id]; !ok {
		return errors.New("file not found")
	}

	f := l.Files[id]
	if !f.Exists() {
		delete(l.Files, id)
		return nil
	}

	err := os.Remove(f.Path)
	if err != nil {
		return err
	}
	delete(l.Files, id)
	return l.Save()
}

func (l *LibraryService) Save() error {
	dat, err := json.Marshal(l)
	if err != nil {
		return err
	}
	err = os.WriteFile(l.config.LibraryFileName, dat, 0777)
	if err == nil {
		log.Println("wrote library file:", l.config.LibraryFileName)
	}
	return err
}

func (l *LibraryService) GetFiles() []*models.File {
	r := make([]*models.File, 0)
	for _, v := range l.Files {
		r = append(r, v)
	}
	return r
}

func (li *LibraryService) GetFileByID(id uuid.UUID) (*models.File, error) {
	v, ok := li.Files[id]
	if !ok {
		return nil, errFileNotFound
	}
	return v, nil
}
