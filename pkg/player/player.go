package player

import (
	"errors"

	"github.com/lakrizz/sleepi/pkg/library"
)

const (
	StateStopped = iota
	StatePlaying
	StatePaused
	StateUnknown
)

var (
	NoSong = Song{
		Length: -1,
		Pos:    -1,
	}

	ErrFileNotFound = errors.New("file not found, cannot play")
	ErrNoSongLoaded = errors.New("currently no song is loaded")
	ErrQueueEmpty   = errors.New("queue is empty")
)

type Player interface {
	Play() error
	Stop() error
	Pause() error

	Skip() error

	SetVolume(int) error
	GetVolume() (int, error)

	Queue(*library.File) error
	QueueMany([]*library.File, bool) error

	GetCurrentSong() (Song, error)
}

type PlayerOption func(Player) error
