package entities

import (
	"errors"
	"slices"
	"time"

	"github.com/google/uuid"

	"github.com/lakrizz/sleepi/internal/domain/shared"
)

// AlarmID is a strong typedef for clarity.
type AlarmID string

// Alarm represents a single configured alarm.
type Alarm struct {
	ID             AlarmID
	Label          string
	TimeOfDay      shared.TimeOfDay // e.g. 07:30
	Enabled        bool
	WarmupDuration time.Duration
	LEDTarget      *shared.RGB      // optional
	Playable       uuid.UUID        // playlist or file reference
	Repeat         []shared.Weekday // MONDAY..SUNDAY
	CreatedAt      time.Time
	UpdatedAt      time.Time
}

func NewAlarm(
	label string,
	timeOfDay shared.TimeOfDay,
	playable uuid.UUID,
	repeat []shared.Weekday,
	warmup time.Duration,
	ledTarget *shared.RGB,
) (*Alarm, error) {
	if label == "" {
		return nil, errors.New("label cannot be empty")
	}

	if playable == uuid.Nil || playable.Variant() == uuid.Invalid {
		return nil, errors.New("playable must be set and valid")
	}

	if timeOfDay.Hour() < 0 || timeOfDay.Hour() > 23 || timeOfDay.Minute() < 0 || timeOfDay.Minute() > 59 {
		return nil, errors.New("invalid time of day")
	}

	now := time.Now()

	return &Alarm{
		ID:             AlarmID(uuid.NewString()),
		Label:          label,
		TimeOfDay:      timeOfDay,
		Enabled:        true,
		WarmupDuration: warmup,
		LEDTarget:      ledTarget,
		Playable:       playable,
		Repeat:         repeat,
		CreatedAt:      now,
		UpdatedAt:      now,
	}, nil
}

func (a *Alarm) TriggersAt(t time.Time) bool {
	if !a.Enabled {
		return false
	}

	if len(a.Repeat) == 0 {
		return a.TimeOfDay.Matches(t)
	}

	if slices.Contains(a.Repeat, shared.Weekday(t.Weekday())) {
		return a.TimeOfDay.Matches(t)
	}

	return false
}

func (a *Alarm) TimeUntilNextTrigger() time.Duration {
	if !a.Enabled {
		return -1
	}

	now := time.Now()
	// start today
	checkDate := time.Date(now.Year(), now.Month(), now.Day(), a.TimeOfDay.Hour(), a.TimeOfDay.Minute(), 0, 0, now.Location())

	for checkDate.Before(time.Now()) || !a.TriggersAt(checkDate) {
		// add a day
		checkDate = checkDate.Add(24 * time.Hour)
	}

	return time.Until(checkDate)
}
