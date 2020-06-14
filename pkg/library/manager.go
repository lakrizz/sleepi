package library

import (
	"encoding/json"
	"fmt"
	"io/ioutil"
	"path"

	"github.com/adrg/xdg"
	"github.com/google/uuid"
)

func createEmptyLibrary(filename string) *Library {
	// we assume the file does not exist
	files := make(map[uuid.UUID]*File)
	return &Library{
		filename: filename,
		Root:     path.Join(xdg.UserDirs.Documents, "sleepi"),
		Files:    files,
	}
}

func LoadLibrary() (*Library, error) {
	filename, err := xdg.ConfigFile("sleepi/library.json")

	dat, err := ioutil.ReadFile(filename)
	if err != nil {
		cfg := createEmptyLibrary(filename)
		return cfg, cfg.Save()
	}

	var c *Library
	err = json.Unmarshal(dat, &c)
	if err != nil {
		return nil, err
	}
	c.filename = filename
	return c, nil
}

func (c *Library) Save() error {
	dat, err := json.Marshal(c)
	if err != nil {
		return err
	}
	fmt.Println(string(dat))
	err = ioutil.WriteFile(c.filename, dat, 0755)
	return err
}
