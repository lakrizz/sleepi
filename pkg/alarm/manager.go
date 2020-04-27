package alarm

import (
	"errors"
	"fmt"
)

type AlarmManager struct {
	Alarms  []*Alarm
	watcher *alarmWatcher
}

func CreateAlarmManager() (*AlarmManager, error) {
	alarms, err := loadAlarms()
	if err != nil {
		return nil, err
	}

	am := &AlarmManager{Alarms: alarms}
	if len(alarms) > 0 { // if there's no alarms, there's nothing to watch
		aw, err := createWatcher(am)
		if err != nil {
			return nil, err
		}
		am.watcher = aw
		go am.watcher.run()
	}
	return am, nil
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
	if len(a.Alarms) == 0 {
		return nil, errors.New("there are no alarms")
	}

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
