package main

import (
	"errors"

	"./lib"
)

type manager struct {
	alarms    []*lib.Alarm
	playlists []*lib.Playlist
}

var Manager *manager

func Init() {
	Manager = &manager{make([]*lib.Alarm, 0), make([]*lib.Playlist, 0)}
}

func (m *manager) AddAlarm(alarm *lib.Alarm) error {
	if alarm == nil {
		return errors.New("alarm can't be null you sleazy man")
	}
	m.alarms = append(m.alarms, alarm)
	return nil
}

func (m *manager) AddPlaylist(playlist *lib.Playlist) error {
	if playlist == nil {
		return errors.New("playlist can't be null you sleazy man")
	}
	m.playlists = append(m.playlists, playlist)
	return nil
}

func Run() {
scheduler.
}
