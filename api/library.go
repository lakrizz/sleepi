package api

import (
	"errors"
	"io/ioutil"
	"mime/multipart"
	"os"
	"path"
)

func (a *Api) AddSongToLibrary(file *multipart.FileHeader) error {
	target_name := path.Join(a.Library.Root, file.Filename)

	// sanity check: does the file already exist?
	if _, err := os.Stat(target_name); os.IsExist(err) {
		return errors.New("file already exists!")
	}

	fh, err := file.Open()
	if err != nil {
		return err
	}

	content, err := ioutil.ReadAll(fh)
	if err != nil {
		return err
	}

	err = ioutil.WriteFile(target_name, content, 0777)

	if err != nil {
		return err
	}

	return a.Library.Refresh()
}
