package api

import (
	"errors"

	"github.com/google/uuid"
	"github.com/lakrizz/sleepi/pkg/alarm"
)

func (a *Api) AddAlarm() {}

func (a *Api) GetAlarms() ([]*alarm.Alarm, error) {
	if a.Alarms == nil {
		return nil, errors.New("there's no alarm manager, we cannot do anything :(")
	}
	return a.Alarms.Alarms, nil
}

func (a *Api) EditAlarm() {}

func (a *Api) RemoveAlarm() {}

func (a *Api) GetAlarm(id_s string) (*alarm.Alarm, error) {
	id, err := uuid.Parse(id_s)
	if err != nil {
		return nil, err
	}
	return a.Alarms.GetAlarm(id)

} // by id

func (a *Api) EnableAlarm(id_s string) error {
	id, err := uuid.Parse(id_s)
	if err != nil {
		return err
	}

	alarm, err := a.Alarms.GetAlarm(id)
	if err != nil {
		return err
	}

	alarm.Enabled = true
	return a.Alarms.SaveAlarms()
}

func (a *Api) DisableAlarm(id_s string) error {
	id, err := uuid.Parse(id_s)
	if err != nil {
		return err
	}

	alarm, err := a.Alarms.GetAlarm(id)
	if err != nil {
		return err
	}

	alarm.Enabled = false
	return a.Alarms.SaveAlarms()
}
