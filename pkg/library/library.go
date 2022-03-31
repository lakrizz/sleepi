package library

import (
	"encoding/json"
	"errors"
	"io/ioutil"
	"log"
	"os"
	"path/filepath"

	"github.com/adrg/xdg"
	"github.com/google/uuid"
	"krizz.org/sleepi/pkg/helper"
)

type Library struct {
	media_folder string
	Files        map[uuid.UUID]*File
}

var config_name string = "library.json"

// GetLibrary creates a library which is nothing but a bunch of files
// stored in a collection, more like a fs, it's an abstraction layer
// for any file consumer to access files without having the need
// to know about the location
func GetLibrary() (*Library, error) {
	file, err := helper.GetFullConfigPath(config_name)
	if err != nil {
		return nil, err
	}

	dat, err := ioutil.ReadFile(file)
	if err != nil {
		return nil, err
	}

	if len(dat) == 0 { // file is new or empty
		media_folder := xdg.UserDirs.Documents
		mm := make(map[uuid.UUID]*File)
		if m, err := walkFolder(media_folder); err == nil {
			mm = m
		} else {
			log.Println(err)
		}
		return &Library{media_folder: media_folder, Files: mm}, nil
	}

	var lib *Library
	err = json.Unmarshal(dat, &lib)
	if err != nil {
		return nil, err
	}

	return lib, nil
}

func walkFolder(folder string) (map[uuid.UUID]*File, error) {
	if _, err := os.Stat(folder); os.IsNotExist(err) {
		return nil, err
	}
	files, err := ioutil.ReadDir(folder)
	if err != nil {
		return nil, err
	}

	for _, v := range files {
		if v.Name() == config_name {
			// found it, read it!
			dat, err := ioutil.ReadFile(filepath.Join(folder, v.Name()))
			if err != nil {
				return nil, err
			}
			filelist := make(map[uuid.UUID]*File)
			err = json.Unmarshal(dat, &filelist)
			if err != nil {
				return nil, err
			}
			return filelist, nil
		}
	}
	return nil, errors.New("no config found")
}

func (l *Library) AddFile(data []byte, name string) error {

	err := ioutil.WriteFile(l.media_folder, data, 0777)
	if err != nil {
		return err
	}

	id := uuid.New()
	f := &File{Location: l.media_folder, Id: id}
	l.Files[f.Id] = f
	l.save()
	return nil
}

func (l *Library) save() error {
	fn, err := helper.GetFullConfigPath(config_name)
	if err != nil {
		return err
	}

	dat, err := json.Marshal(l.Files)
	if err != nil {
		return err
	}
	err = ioutil.WriteFile(fn, dat, 0777)
	return err
}

func (l *Library) GetAllFiles() []*File {
	r := make([]*File, 0)
	for _, v := range l.Files {
		r = append(r, v)
	}
	return r
}
