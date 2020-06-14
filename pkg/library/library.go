package library

import (
	"errors"
	"log"
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

var extAllowlist []string = []string{"mp3"} // currently only mp3s are supported, eh?

func (l *Library) Refresh() error {
	if !l.scanning {
		return l.walkdir(l.Root)
	}

	return errors.New("library is already scanning")
}

func (l *Library) walkdir(basedir string) error {
	err := filepath.Walk(basedir, func(walkpath string, info os.FileInfo, err error) error {
		log.Println("walking", walkpath)
		if err != nil {
			return err
		}

		if info.IsDir() {
			return l.walkdir(path.Join(walkpath, info.Name()))
		}

		if isAllowedExtension(path.Ext(info.Name()), extAllowlist) {
			// this is a valid file, yay
			if !l.fileIndexed(info.Name()) {
				newid, err := uuid.NewRandom()
				if err != nil {
					return err
				}
				l.Files[newid] = &File{newid, basedir, info.Name()}
				log.Println("addiung file", info.Name())
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
