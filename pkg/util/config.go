package util

import (
	"encoding/json"
	"fmt"
	"io/ioutil"
	"os"
	"path"

	"github.com/adrg/xdg"
)

// FindConfig returns the full qualified path for a given filename
// this function searches through all standard paths
// and creates a new config file if it's been empty
func GetFullConfigPath(name string) (string, error) {
	folder := path.Join(xdg.UserDirs.Documents, "sleepi")
	err := os.MkdirAll(folder, 0777)
	if err != nil {
		return "", err
	}

	filename := fmt.Sprintf("%v/%v", folder, name)
	if _, err = os.Stat(filename); os.IsNotExist(err) {
		_, err = os.Create(filename)
		if err != nil {
			return "", err
		}
	}

	return filename, nil
}

func ReadOrCreateConfigFile(filename string, obj any) error {
	dat, err := ioutil.ReadFile(filename)
	if err != nil {
		return err
	}

	if len(dat) == 0 {
		return nil // file is empty, but that's valid
	}

	err = json.Unmarshal(dat, &obj)
	if err != nil {
		return err
	}

	return nil
}
