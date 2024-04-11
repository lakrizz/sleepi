package models

import (
	"errors"
	"fmt"
	"io"
	"os"
	"path/filepath"

	"github.com/google/uuid"
)

type File struct {
	Path string
	ID   uuid.UUID
}

// Exists method figures out if the file exists in the filesystem
func (f *File) Exists() bool {
	if _, err := os.Stat(f.Path); os.IsNotExist(err) {
		return false
	}
	return true
}

// Read method reads the whole file and returns the contents
func (f *File) Read() ([]byte, error) {
	if !f.Exists() {
		return nil, fmt.Errorf("%v has no valid physical location (was looking for %v)", f.ID, f.Path)
	}
	dat, err := os.ReadFile(f.Path)
	if err != nil {
		return nil, err
	}
	return dat, nil
}

// Name method returns the pure filename of the file
func (f *File) Name() string {
	return filepath.Base(f.Path)
}

// GetReadCloser method returns a readcloser for the file
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
