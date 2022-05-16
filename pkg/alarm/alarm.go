package alarm

import (
	"errors"
	"time"

	"github.com/google/uuid"
	"krizz.org/sleepi/pkg/effects"
)

type Alarm struct {
	Id         uuid.UUID
	ActiveDays []time.Weekday
	WakeHour   int
	WakeMinute int
	Playlist   *uuid.UUID

	// effects go here
	VolumeWarmup *effects.VolumeWarmup
}

func CreateAlarm(playlist_id *uuid.UUID, activedays []time.Weekday, wakehour, wakeminute int) (*Alarm, error) {
	a := &Alarm{ActiveDays: activedays, WakeHour: wakehour, WakeMinute: wakeminute, Playlist: playlist_id}
	id := uuid.New()
	a.Id = id
	return a, nil
}

func (a *Alarm) AddVolumeWarmup(fx *effects.VolumeWarmup) error {
	if fx == nil {
		return errors.New("effect shouldn't be null")
	}

	a.VolumeWarmup = fx
	return nil
}

func (a *Alarm) DurationUntilNextAlarm() time.Duration {
	now := time.Now()
	checkdate := time.Date(now.Year(), now.Month(), now.Day(), a.WakeHour, a.WakeMinute, 0, 0, now.Location())

	// special case when the alarm is still going off on this particular day
	if (a.WakeHour*60 + a.WakeMinute) > (now.Hour()*60 + now.Minute()) {
		return checkdate.Sub(now)
	}

	// otherwise add days until we look at an active weekday
	checkdate = checkdate.Add(24 * time.Hour)
	for !a.isActiveDay(checkdate.Weekday()) {
		checkdate = checkdate.Add(24 * time.Hour)
	}

	return time.Until(checkdate)
}

func (a *Alarm) isActiveDay(day time.Weekday) bool {
	for _, v := range a.ActiveDays {
		if v == day {
			return true
		}
	}
	return false
}

func (a *Alarm) Trigger() error {
	return nil
}
