package alarm

import (
	"encoding/json"
	"errors"
	"fmt"
	"io/ioutil"

	"github.com/k0kubun/pp"
)

type AlarmManager struct {
	Alarms []*Alarm
}

func CreateAlarmManager(filename string) (*AlarmManager, error) {
	alarms, err := loadAlarms(filename)
	if err != nil {
		return nil, err
	}

	am := &AlarmManager{Alarms: alarms}
	pp.Println(am)
	return am, nil
}

func loadAlarms(filename string) ([]*Alarm, error) {
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

func (a *AlarmManager) AddAlarm(alarm *Alarm) error {
	for _, v := range a.Alarms {
		if v.Name == alarm.Name {
			return errors.New(fmt.Sprintf("alarm with the name %s is already registered", alarm.Name))
		}
	}
	a.Alarms = append(a.Alarms, alarm)
	return nil
}

func (a *AlarmManager) GetNextAlarm() (*Alarm, error) {
	// baseline is the first alarm
	next, err := a.Alarms[0].TimeTillNextWake()
	if err != nil {
		return nil, err
	}
	idx := 0
	for i, v := range a.Alarms[0:] {
		n, err := v.TimeTillNextWake()
		if err != nil {
			return nil, err
		}

		if n < next {
			next = n
			idx = i
		}
	}

	return a.Alarms[idx], nil

}

func (a *AlarmManager) Run() error {
	for {
		// need a endless loop because we need to watch all alarms
	}
}
