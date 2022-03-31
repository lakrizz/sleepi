package library

import (
	"fmt"
	"io/ioutil"
	"os"

	"github.com/google/uuid"
)

type File struct {
	Path string
	Id   uuid.UUID
}

func (f *File) existsPhysically() bool {
	if _, err := os.Stat(f.Path); err != os.ErrNotExist {
		return true
	}
	return false
}

func (f *File) Read() ([]byte, error) {
	if !f.existsPhysically() {
		return nil, fmt.Errorf("%v has no valid physical location (was looking for %v)", f.Id, f.Path)
	}
	dat, err := ioutil.ReadFile(f.Path)
	if err != nil {
		return nil, err
	}
	return dat, nil
}
