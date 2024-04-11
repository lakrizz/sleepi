package alarm

import (
	"errors"
	"slices"
	"time"

	"github.com/google/uuid"

	"github.com/lakrizz/sleepi/pkg"
)

var (
	errNoActiveDays = errors.New("no active days set")
)

type Alarm struct {
	Enabled    bool
	name       string
	ID         uuid.UUID
	activeDays []time.Weekday
	wakeHour   int
	wakeMinute int
	playlist   *uuid.UUID

	// effects go here
	effects []pkg.Effect
}

func CreateAlarm(playlistID *uuid.UUID, name string, activedays []time.Weekday, wakehour, wakeminute int) (*Alarm, error) {
	a := &Alarm{
		Enabled:    true,
		activeDays: activedays,
		wakeHour:   wakehour,
		wakeMinute: wakeminute,
		playlist:   playlistID,
		name:       name,
		ID:         uuid.New(),
	}
	return a, nil
}

func (a *Alarm) AddEffects(fx ...pkg.Effect) error {
	a.effects = append(a.effects, fx...)
	return nil
}

func (a *Alarm) DurationUntilNextTrigger() time.Duration {
	if len(a.activeDays) == 0 {
		return time.Duration(-1)
	}

	now := time.Now()
	checkdate := time.Date(now.Year(), now.Month(), now.Day(), a.wakeHour, a.wakeMinute, 0, 0, now.Location())

	// special case when the alarm is still going off on this particular day
	if (a.wakeHour*60 + a.wakeMinute) > (now.Hour()*60 + now.Minute()) {
		return checkdate.Sub(now)
	}

	// otherwise add days until we look at an active weekday
	checkdate = checkdate.Add(24 * time.Hour)
	for !a.IsActiveDay(checkdate.Weekday()) {
		checkdate = checkdate.Add(24 * time.Hour)
	}

	return time.Until(checkdate)
}

// IsActiveDay method returns a boolean value on whether the alarm will be
// triggered on a given weekday
func (a *Alarm) IsActiveDay(day time.Weekday) bool {
	return slices.Contains(a.activeDays, day)
}

func (a *Alarm) Trigger() error {
	return nil
}
