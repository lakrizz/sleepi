package domain

import (
	"github.com/lakrizz/sleepi/internal/playback/entities"
)

type Player interface {
	Load(entities.Playable) error
	Play() error
	Pause() error
	Stop() error
}
