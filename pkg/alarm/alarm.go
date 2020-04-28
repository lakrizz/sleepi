package alarm

import (
	"errors"
	"time"

	"github.com/google/uuid"
)

type Alarm struct {
	Name       string         `json:"Name"`
	Id         uuid.UUID      `json:"Id"`
	WakeHour   int            `json:"WakeHour"`
	WakeMinute int            `json:"WakeMinute"`
	Days       []time.Weekday `json:"Days"`
	Playlist   uuid.UUID      `json:"Playlist"`
	WakeupTime string         `json:"WakeupTime"`
	Enabled    bool           `json:"Enabled"`
}

func (a *Alarm) NextWake() (time.Time, error) {
	// tomorrow?
	for i := 0; i < 7; i++ { // amount of added days
		date, err := a.checkDate(i)
		if err == nil {
			return date, nil
		}
	}
	return time.Now(), errors.New("it's fucked")
}

func (a *Alarm) checkDate(add_days int) (time.Time, error) {
	now := time.Now().AddDate(0, 0, add_days)
	if (now.Hour() < a.WakeHour || (now.Hour() == a.WakeHour && now.Minute() < a.WakeMinute) || add_days > 0) && a.ringsOnWeekday(now.Weekday()) {
		return time.Date(now.Year(), now.Month(), now.Day(), a.WakeHour, a.WakeMinute, 0, 0, time.Local), nil
	}

	return time.Now(), errors.New("not today!")
}

func (a *Alarm) TimeTillNextWake() (time.Duration, error) {
	nextWake, err := a.NextWake()
	if err != nil {
		return 0, err
	}
	dur, err := time.ParseDuration(a.WakeupTime)
	if err != nil {
		return 0, err
	}

	return time.Until(nextWake.Add(-dur)), nil
}

func (a *Alarm) TimeTillNextWakeWithoutWarmup() (time.Duration, error) {
	nextWake, err := a.NextWake()
	if err != nil {
		return 0, err
	}
	return time.Until(nextWake), nil
}

func (a *Alarm) ringsOnWeekday(weekday time.Weekday) bool {
	for _, v := range a.Days {
		if v == weekday {
			return true
		}
	}
	return false
}
