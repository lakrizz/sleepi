package player

import (
	"errors"

	"github.com/lakrizz/sleepi/pkg/models"
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
	Clear() error

	SetVolume(int) error
	GetVolume() (int, error)

	Queue(*models.File) error
	QueueMany([]*models.File, bool) error

	GetCurrentSong() (Song, error)
}

type PlayerOption func(Player) error
