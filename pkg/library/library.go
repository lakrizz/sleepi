package library

import (
	"errors"
	"os"
	"path"
	"path/filepath"

	"github.com/google/uuid"
)

type Library struct {
	filename string              `json:"-"`
	Root     string              `json:"root"`
	Files    map[uuid.UUID]*File `json:"Files"` // key is the subdir and the element is the filename
	scanning bool                `json:"-"`
}

func (l *Library) Refresh() error {
	if !l.scanning {
		return l.walkdir(l.Root)
	}

	return errors.New("library is already scanning")
}

var extAllowlist []string = []string{"mp3"} // currently only mp3s are supported, eh?

func (l *Library) walkdir(dir string) error {
	err := filepath.Walk(dir, func(walkpath string, info os.FileInfo, err error) error {
		if err != nil {
			return err
		}

		if info.IsDir() {
			return l.walkdir(info.Name())
		}

		if isAllowedExtension(path.Ext(info.Name()), extAllowlist) {
			// this is a valid file, yay
			if !l.fileIndexed(info.Name()) {
				newid, err := uuid.NewRandom()
				if err != nil {
					return err
				}
				l.Files[newid] = &File{newid, dir, info.Name()}
			}
		}
		return nil
	})
	return err
}

func (l *Library) fileIndexed(dir string) bool {
	for _, v := range l.Files {
		if v.Path == dir {
			return true
		}
	}
	return false
}

func isAllowedExtension(ext string, allowlist []string) bool {
	for _, v := range allowlist {
		if ext == v {
			return true
		}
	}
	return false
}

func (c *Library) GetFile(id uuid.UUID) *File {
	for _, v := range c.Files {
		if v.Id == id {
			return v
		}
	}
	return nil
}
