package library

import (
	"encoding/json"
	"errors"
	"io/ioutil"
	"log"
	"os"
	"path"
	"path/filepath"

	"github.com/adrg/xdg"
	"github.com/google/uuid"
)

type library struct {
	base_folder string
	Files       map[uuid.UUID]*File
}

var config_name string = "library.json"

// GetLibrary creates a library which is nothing but a bunch of files
// stored in a collection, more like a fs, it's an abstraction layer
// for any file consumer to access files without having the need
// to know about the location
func GetLibrary() (*library, error) {
	folder := path.Join(xdg.UserDirs.Documents, "sleepi")
	if _, err := os.Stat(folder); os.IsNotExist(err) {
		err := os.Mkdir(folder, 0777)
		if err != nil {
			return nil, err
		}
	}

	// check if there is a file list inside this folder
	mm := make(map[uuid.UUID]*File)
	if m, err := searchConfig(folder); err == nil {
		mm = m
	} else {
		log.Println(err)
	}
	return &library{base_folder: folder, Files: mm}, nil
}

func searchConfig(folder string) (map[uuid.UUID]*File, error) {
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

func (l *library) AddFile(data []byte, name string) error {
	loc := filepath.Join(l.base_folder, name)
	err := ioutil.WriteFile(loc, data, 0777)
	if err != nil {
		return err
	}

	id := uuid.New()
	f := &File{Location: loc, Id: id}
	l.Files[f.Id] = f
	l.save()
	return nil
}

func (l *library) save() error {
	fn := filepath.Join(l.base_folder, config_name)
	dat, err := json.Marshal(l.Files)
	if err != nil {
		return err
	}
	err = ioutil.WriteFile(fn, dat, 0777)
	return err
}

func (l *library) GetAllFiles() []*File {
	r := make([]*File, 0)
	for _, v := range l.Files {
		r = append(r, v)
	}
	return r
}
