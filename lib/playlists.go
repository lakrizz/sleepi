package lib

import (
	"errors"
	"strings"

	"github.com/nu7hatch/gouuid"
)

type video struct {
	id   *uuid.UUID
	url  string
	file string
}

type Playlist struct {
	id     *uuid.UUID
	videos []*video
}

func (p *Playlist) AddVideo(url string) (*video, error) {
	if !YoutubeDLInstalled() {
		return nil, errors.New("youtube-dl is not installed, can't do it, mate!")
	}

	title, err := GetTitle(url)
	if err != nil {
		return nil, err
	}
	title = strings.TrimRight(title, " \n\b")

	uuid, err := uuid.NewV4()
	if err != nil {
		return nil, err
	}

	fl, err := DownloadVideo(url)
	if err != nil {
		return nil, err
	}
	path, err := CpFile(fl, title)
	if err != nil {
		return nil, err
	}

	return &video{uuid, url, path}, nil
}

func CreatePlaylist() (*Playlist, error) {
	uuid, err := uuid.NewV4()
	if err != nil {
		return nil, err
	}
	return &Playlist{uuid, make([]*video, 0)}, nil
}
