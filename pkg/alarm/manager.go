package alarm

import (
	"errors"
	"fmt"

	"github.com/google/uuid"
)

func (a *AlarmManager) AddAlarm(alarm *Alarm) error {
	for _, v := range a.Alarms {
		if v.Name == alarm.Name {
			return errors.New(fmt.Sprintf("alarm with the name %s already exists", alarm.Name))
		}
	}

	if alarm.Id == uuid.Nil {
		id, err := uuid.NewRandom()
		if err != nil {
			return err
		}

		alarm.Id = id
	}
	a.Alarms = append(a.Alarms, alarm)
	return nil
}

func (a *AlarmManager) GetAlarm(id uuid.UUID) (*Alarm, error) {
	for _, v := range a.Alarms {
		if v.Id == id {
			return v, nil
		}
	}
	return nil, errors.New(fmt.Sprintf("could not find alarm with id %s", id.String()))
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
