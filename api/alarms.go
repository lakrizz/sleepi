package api

import (
	"errors"

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

func (a *Api) GetAlarm() {} // by id
