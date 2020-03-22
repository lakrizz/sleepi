package alarm

import (
	"encoding/json"
	"errors"
	"io/ioutil"
	"log"
	"time"

	"github.com/k0kubun/pp"
)

type Alarm struct {
	Name       string         `json:"Name"`
	WakeHour   int            `json:"WakeHour"`
	WakeMinute int            `json:"WakeMinute"`
	Days       []time.Weekday `json:"Days"`
	Playlist   string         `json:"Playlist"`
	WakeupTime string         `json:"WakeupTime"`
	Enabled    bool           `json:"Enabled"`
}

func LoadAlarms(filename string) ([]*Alarm, error) {
	alarms := make([]*Alarm, 0)
	dat, err := ioutil.ReadFile(filename)
	if err != nil {
		return nil, err
	}

	err = json.Unmarshal(dat, &alarms)
	if err != nil {
		return nil, err
	}
	return alarms, nil
}

func (a *Alarm) NextWake() (time.Time, error) {
	now := time.Now()
	// pp.Println("A:", a)
	pp.Println("Hour", now.Hour() <= a.WakeHour)
	pp.Println("Minute", now.Minute() <= a.WakeMinute)
	if now.Hour() < a.WakeHour || (now.Hour() == a.WakeHour && now.Minute() <= a.WakeMinute) && a.ringsOnWeekday(now.Weekday()) { // the alarm will go off today :O
		log.Println("alles cool!")
		return time.Date(now.Year(), now.Month(), now.Day(), a.WakeHour, a.WakeMinute, 0, 0, time.Local), nil
	}
	return now, errors.New("was auch immer")
}

func (a *Alarm) TimeTillNextWake() (time.Duration, error) {
	nextWake, err := a.NextWake()
	if err != nil {
		return 0, err
	}
	return time.Until(nextWake), nil
}

func (a *Alarm) ringsOnWeekday(weekday time.Weekday) bool {
	pp.Println("hi")
	for _, v := range a.Days {
		if v == weekday {
			pp.Println(v, weekday)
			return true
		}
	}
	return false
}
