package main

import (
	"errors"

	"./lib"
)

type manager struct {
	Alarms    []*lib.Alarm
	Playlists []*lib.Playlist
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

func (m *manager) GetPlaylist(uuid string) (*lib.Playlist, error) {
	id, err := uuid.Parse(uuid)

	if err != nil {
		return nil, err
	}

	for _, p := range m.Playlists {
		if p.id == id {
			return p, nil
		}
	}

	return nil, errors.New("playlist not found")
}
