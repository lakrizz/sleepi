package entities

import "github.com/google/uuid"

type Playlist struct {
	ID   uuid.UUID
	Name string

	Items []*Playable
}
