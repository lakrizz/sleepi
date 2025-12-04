package dto

import (
	"github.com/google/uuid"

	"github.com/lakrizz/sleepi/internal/domain/shared"
)

type AlarmAddCommand struct {
	Label          string
	TimeOfDay      shared.TimeOfDay // e.g. 07:30
	Enabled        bool
	WarmupDuration string
	LEDTarget      shared.RGB       // optional
	Playable       uuid.UUID        // playlist or file reference
	Repeat         []shared.Weekday // MONDAY..SUNDAY
}
