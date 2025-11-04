package shared

import "github.com/google/uuid"

type Playable struct {
	ID uuid.UUID
}

// NewPlayable constructs a Playable from known data.
func NewPlayable(id uuid.UUID) Playable {
	return Playable{ID: id}
}

// IsZero checks if the Playable is uninitialized.
func (p Playable) IsZero() bool {
	return p.ID == uuid.Nil
}
