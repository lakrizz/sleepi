package players

import (
	"github.com/fhs/gompd/v2/mpd"

	"example.com/sleepi/internal/playback/domain"
)

type MPDPlayer struct {
	client *mpd.Client
}

func InitMPDPlayer() (*MPDPlayer, error) {
	client, err := mpd.Dial("tcp", "127.0.0.1:6600")
	if err != nil {
		return nil, err
	}

	m := &MPDPlayer{
		client: client,
	}

	return m, nil
}

func (mp *MPDPlayer) Load(_ domain.Source) error {
	panic("not implemented") // TODO: Implement
}

func (mp *MPDPlayer) Play() error {
	panic("not implemented") // TODO: Implement
}

func (mp *MPDPlayer) Pause() error {
	panic("not implemented") // TODO: Implement
}

func (mp *MPDPlayer) Stop() error {
	panic("not implemented") // TODO: Implement
}
