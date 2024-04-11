package config

import (
	"path"

	"github.com/adrg/xdg"
)

type Config struct {
	MPDProtocol string
	MPDHost     string

	HTTPHost string
	HTTPPort string

	LibraryFileName   string
	PlaylistsFileName string
	AlarmsFileName    string

	MusicFolderFileName string
}

func GetConfig() (*Config, error) {
	sleepiDocumentFolder := path.Join(xdg.UserDirs.Documents, "sleepi")

	return &Config{
		MPDProtocol:         "tcp",
		MPDHost:             "localhost:6600",
		HTTPHost:            "localhost",
		HTTPPort:            "8080",
		MusicFolderFileName: path.Join(sleepiDocumentFolder, "music"),
		LibraryFileName:     path.Join(sleepiDocumentFolder, "library.json"),
		PlaylistsFileName:   path.Join(sleepiDocumentFolder, "playlists.json"),
		AlarmsFileName:      path.Join(sleepiDocumentFolder, "alarms.json"),
	}, nil
}
