package library

import (
	"errors"
	"fmt"
	"io"
	"io/ioutil"
	"os"
	"path/filepath"

	"github.com/google/uuid"
	"github.com/k0kubun/pp"
)

type File struct {
	Path string
	Id   uuid.UUID
}

func (f *File) Exists() bool {
	if _, err := os.Stat(f.Path); os.IsNotExist(err) {
		pp.Println(err)
		return false
	}
	return true
}

func (f *File) Read() ([]byte, error) {
	if !f.Exists() {
		return nil, fmt.Errorf("%v has no valid physical location (was looking for %v)", f.Id, f.Path)
	}
	dat, err := ioutil.ReadFile(f.Path)
	if err != nil {
		return nil, err
	}
	return dat, nil
}

func (f *File) Name() string {
	return filepath.Base(f.Path)
}

func (f *File) GetReadCloser() (io.ReadCloser, error) {
	if !f.Exists() {
		return nil, errors.New("does not exist")
	}

	fh, err := os.Open(f.Path)
	if err != nil {
		return nil, err
	}

	return fh, nil
}
