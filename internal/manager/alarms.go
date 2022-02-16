package manager

import (
	"errors"

	"krizz.org/sleepi/pkg/alarm"
)

type AlarmManager struct {
	Alarms []*alarm.Alarm
}

func GetAlarmManager(alarms []*alarm.Alarm) (*AlarmManager, error) {
	am := &AlarmManager{Alarms: alarms}
	am.listen()
	return am, nil
}

func (am *AlarmManager) GetClosestAlarm() (*alarm.Alarm, error) {
	if len(am.Alarms) == 0 {
		return nil, errors.New("no alarms")
	}

	if len(am.Alarms) == 1 {
		return am.Alarms[0], nil
	}

	closest := am.Alarms[0]
	for _, a := range am.Alarms[0:] {
		if a.DurationUntilNextAlarm() < closest.DurationUntilNextAlarm() {
			closest = a
		}
	}

	return closest, nil
}

func (am *AlarmManager) listen() {
	for {

	}
}
