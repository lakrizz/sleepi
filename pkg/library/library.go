package library

import (
	"errors"
	"fmt"
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

var extAllowlist []string = []string{".mp3", ".m4a", ".flac"} // currently only mp3s are supported, eh?

func (l *Library) Refresh() error {
	if !l.scanning {
		return l.walkdir(l.Root)
	}

	return errors.New("library is already scanning")
}

func (l *Library) walkdir(basedir string) error {
	err := filepath.Walk(basedir, func(walkpath string, info os.FileInfo, err error) error {
		if err != nil {
			return err
		}

		if isAllowedExtension(path.Ext(info.Name()), extAllowlist) {
			// this is a valid file, yay
			if !l.fileIndexed(info.Name()) {
				newid, err := uuid.NewRandom()
				if err != nil {
					return err
				}
				l.Files[newid] = &File{newid, walkpath, info.Name()}
				log.Println("addiung file", info.Name())
			}
		}
		return nil
	})

	if err != nil {
		return err
	}

	return l.Save()
}

func (l *Library) fileIndexed(path string) bool {
	for _, v := range l.Files {
		if v.Filename == path {
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

func (c *Library) RemoveFile(id uuid.UUID) error {
	// we want to remove the physical file as well :D
	f := c.GetFile(id)
	err := os.Remove(f.Path)
	if err != nil {
		return err
	}

	delete(c.Files, id)
	return nil
}

func (c *Library) CheckFiles() {
	for _, v := range c.Files {
		if !v.Exists() {
			fmt.Println("deleting", v.Path)
			delete(c.Files, v.Id)
		}
	}
}
