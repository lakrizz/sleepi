package alarm

import (
	"log"
	"time"

	"github.com/google/uuid"
)

type Alarm struct {
	Id            uuid.UUID
	AlarmFunction func()
	ActiveDays    []time.Weekday
	WakeHour      int
	WakeMinute    int
}

func CreateAlarm(function func(), activedays []time.Weekday, wakehour, wakeminute int) (*Alarm, error) {
	a := &Alarm{ActiveDays: activedays, AlarmFunction: function, WakeHour: wakehour, WakeMinute: wakeminute}
	id := uuid.New()
	a.Id = id
	return a, nil
}

func (a *Alarm) DurationUntilNextAlarm() time.Duration {
	now := time.Now()
	checkdate := time.Date(now.Year(), now.Month(), now.Day(), a.WakeHour, a.WakeMinute, 0, 0, now.Location())

	// special case when the alarm is still going off on this particular day
	if (a.WakeHour*60 + a.WakeMinute) < (now.Hour()*60 + now.Minute()) {
		return checkdate.Sub(now)
	}

	// otherwise add days until we look at an active weekday
	for !a.isActiveDay(checkdate.Weekday()) {
		checkdate = checkdate.Add(24 * time.Hour)
	}

	return checkdate.Sub(now)
}

func (a *Alarm) isActiveDay(day time.Weekday) bool {
	for _, v := range a.ActiveDays {
		if v == day {
			log.Println(v, day)
			return true
		}
	}
	return false
}
