package manager

import (
	"krizz.org/sleepi/pkg/audioplayer"
	"krizz.org/sleepi/pkg/library"
)

type Managers struct {
	Alarms      *AlarmManager
	Playlists   *PlaylistManager
	Library     *library.Library
	AudioPlayer *audioplayer.Audioplayer
}

func GetManagers() (*Managers, error) {
	managers := &Managers{}
	pl, err := getPlaylistManager()
	if err != nil {
		return nil, err
	}
	managers.Playlists = pl

	alarms, err := getAlarmManager()
	if err != nil {
		return nil, err
	}
	managers.Alarms = alarms

	lib, err := library.GetLibrary()
	if err != nil {
		return nil, err
	}
	managers.Library = lib

	audioplayer, err := audioplayer.GetAudioplayer()
	if err != nil {
		return nil, err
	}

	managers.AudioPlayer = audioplayer

	return managers, nil
}
