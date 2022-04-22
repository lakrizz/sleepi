package library

import (
	"errors"
	"fmt"
	"io/ioutil"
	"os"

	"github.com/google/uuid"
	"github.com/mikkyang/id3-go"
)

type File struct {
	Path string
	Id   uuid.UUID
	Meta *metaInformation
}

type metaInformation struct {
	Album  string
	Title  string
	Artist string
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

func (f *File) readID3() error {
	if !f.existsPhysically() {
		return errors.New("file does not exist, can't read id3 tags")
	}

	mp3File, err := id3.Open(f.Path)
	if err != nil {
		return err
	}
	defer mp3File.Close()

	f.Meta = &metaInformation{Title: mp3File.Title(), Artist: mp3File.Artist(), Album: mp3File.Album()}
	return nil
}
