package mpd

import (
	"github.com/lakrizz/sleepi/pkg/library"
	"github.com/lakrizz/sleepi/pkg/player"
)

type MPDPlayer struct{}

func (mp *MPDPlayer) Play() error {
	panic("not implemented") // TODO: Implement
}

func (mp *MPDPlayer) Stop() error {
	panic("not implemented") // TODO: Implement
}

func (mp *MPDPlayer) Pause() error {
	panic("not implemented") // TODO: Implement
}

func (mp *MPDPlayer) Skip() error {
	panic("not implemented") // TODO: Implement
}

func (mp *MPDPlayer) SetVolume(_ int) error {
	panic("not implemented") // TODO: Implement
}

func (mp *MPDPlayer) GetVolume() (int, error) {
	panic("not implemented") // TODO: Implement
}

func (mp *MPDPlayer) Queue(_ *library.File) error {
	panic("not implemented") // TODO: Implement
}

func (mp *MPDPlayer) QueueMany(_ []*library.File, _ bool) error {
	panic("not implemented") // TODO: Implement
}

func (mp *MPDPlayer) GetCurrentSong() (player.Song, error) {
	panic("not implemented") // TODO: Implement
}
