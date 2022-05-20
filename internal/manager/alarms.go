package manager

import (
	"encoding/json"
	"errors"
	"fmt"
	"io/ioutil"
	"time"

	"github.com/google/uuid"
	"krizz.org/sleepi/pkg/alarm"
	"krizz.org/sleepi/pkg/util"
)

type AlarmManager struct {
	Alarms      []*alarm.Alarm
	next        *alarm.Alarm
	alarm_timer *time.Timer
}

const (
	alarmManagerFileName = "alarms.json"
)

func getAlarmManager() (*AlarmManager, error) {
	am := &AlarmManager{}
	path, err := util.GetFullConfigPath(alarmManagerFileName)
	if err != nil {
		return nil, err
	}

	err = util.ReadOrCreateConfigFile(path, am)
	if err != nil {
		return nil, err
	}

	err = am.setNext()
	if err != nil {
		return nil, err
	}

	if am.next != nil { // only do this if there's actually an enabled alarm
		go am.listen()
	}
	return am, nil
}

func (am *AlarmManager) UpdateNextAlarm() error {
	return am.setNext()
}

func (am *AlarmManager) getClosestAlarm() (*alarm.Alarm, error) {
	if len(am.Alarms) == 0 {
		return nil, errors.New("no alarms")
	}

	if len(am.Alarms) == 1 {
		return am.Alarms[0], nil
	}

	var closest *alarm.Alarm
	for _, a := range am.Alarms[0:] {
		if a.Enabled && (closest == nil || a.DurationUntilNextAlarm() < closest.DurationUntilNextAlarm()) {
			closest = a
		}
	}

	return closest, nil
}

func (am *AlarmManager) listen() {
	for {
		select {
		case <-am.alarm_timer.C:
			go am.next.Trigger()
			am.setNext()
		default:
			// log.Println(am.next.DurationUntilNextAlarm().String())
			// time.Sleep(150 * time.Millisecond)
			continue
		}
	}
}

func (am *AlarmManager) AddAlarm(alarm *alarm.Alarm) error {
	if am.isInAlarmList(alarm) {
		return fmt.Errorf("could not add alarm with id %v - it's already in list", alarm.Id.String())
	}
	am.Alarms = append(am.Alarms, alarm)
	am.setNext()

	return am.Save()
}

func (am *AlarmManager) GetAlarm(id uuid.UUID) (*alarm.Alarm, error) {
	for _, v := range am.Alarms {
		if v.Id == id {
			return v, nil
		}
	}

	return nil, errors.New("alarm not found")
}

func (am *AlarmManager) isInAlarmList(alarm *alarm.Alarm) bool {
	for _, v := range am.Alarms {
		if v.Id == alarm.Id {
			return true
		}
	}
	return false
}

func (am *AlarmManager) setNext() error {
	if len(am.Alarms) == 0 {
		return errors.New("there are no alarms")
	}

	next, err := am.getClosestAlarm()
	if err != nil {
		return err
	}

	if next == nil {
		// see if there was a timer running and stop it
		if am.alarm_timer != nil {
			am.alarm_timer.Stop()
		}
		// stop all timers
		return nil
	}
	am.next = next
	am.alarm_timer = time.NewTimer(am.next.DurationUntilNextAlarm())
	return nil
}

func (am *AlarmManager) Save() error {
	dat, err := json.Marshal(am)
	if err != nil {
		return err
	}

	fname, err := util.GetFullConfigPath(alarmManagerFileName)
	if err != nil {
		return err
	}

	return ioutil.WriteFile(fname, dat, 0777)
}
