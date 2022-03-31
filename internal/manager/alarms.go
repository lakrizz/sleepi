package manager

import (
	"errors"
	"fmt"
	"time"

	"krizz.org/sleepi/pkg/alarm"
)

type AlarmManager struct {
	Alarms      []*alarm.Alarm
	next        *alarm.Alarm
	alarm_timer *time.Timer
}

const (
	alarmManagerFileName = "alarms.json"
)

func GetAlarmManager() (*AlarmManager, error) {
	am := &AlarmManager{}
	closest, err := am.GetClosestAlarm()
	if err == nil {
		am.setNext(closest)
		go am.listen()
	}
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
		select {
		case <-am.alarm_timer.C:
			go am.next.AlarmFunction()
			closest, _ := am.GetClosestAlarm()
			am.setNext(closest)
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
	if new_next, _ := am.GetClosestAlarm(); am.next != new_next {
		am.setNext(new_next)
	}
	return nil
}

func (am *AlarmManager) isInAlarmList(alarm *alarm.Alarm) bool {
	for _, v := range am.Alarms {
		if v.Id == alarm.Id {
			return true
		}
	}
	return false
}

func (am *AlarmManager) setNext(alarm *alarm.Alarm) {
	if len(am.Alarms) == 0 {
		return
	}
	am.next, _ = am.GetClosestAlarm()
	am.alarm_timer = time.NewTimer(am.next.DurationUntilNextAlarm())
	fmt.Println("new next:", am.next.WakeHour, am.next.WakeMinute)
}
