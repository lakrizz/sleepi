package library

import (
	"fmt"
	"os"
	"path"

	"github.com/google/uuid"
)

type File struct {
	Id       uuid.UUID `json:"id"`
	Path     string    `json:"path"`
	Filename string    `json:"filename"`
}

func (f *File) FullPath() string {
	return path.Join(f.Path, f.Filename)
}

func (f *File) Exists() bool {
	_, err := os.Stat(f.Path)
	if err != nil {
		fmt.Println(err)
		return !os.IsNotExist(err)
	}
	return true
}
