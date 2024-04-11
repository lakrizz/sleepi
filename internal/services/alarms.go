package services

import (
	"encoding/json"
	"errors"
	"fmt"
	"os"
	"slices"
	"time"

	"github.com/google/uuid"

	"github.com/lakrizz/sleepi/config"
	"github.com/lakrizz/sleepi/pkg/alarm"
)

type AlarmService struct {
	Alarms     []*alarm.Alarm
	next       *alarm.Alarm
	alarmTimer *time.Timer

	cfg *config.Config
}

var (
	errAlarmNotFound = errors.New("alarm not found")
)

func NewAlarmService(cfg *config.Config) (*AlarmService, error) {
	dat, err := os.ReadFile(cfg.PlaylistsFileName)
	if err != nil {
		return &AlarmService{
			Alarms: make([]*alarm.Alarm, 0),
			cfg:    cfg,
		}, nil
	}

	am := &AlarmService{cfg: cfg}
	err = json.Unmarshal(dat, &am)
	if err != nil {
		return nil, err
	}

	am.UpdateTimings()  // we surpress the error here since it's not relevant whether there's actually alarms
	if am.next != nil { // if there's already an active alarm, start the listener
		go am.listen()
	}

	return am, nil
}

func (am *AlarmService) getClosestAlarm() (*alarm.Alarm, error) {
	if len(am.Alarms) == 0 {
		return nil, errors.New("no alarms")
	}

	if len(am.Alarms) == 1 {
		return am.Alarms[0], nil
	}

	var closest *alarm.Alarm
	for _, a := range am.Alarms[0:] {
		if a.Enabled && (closest == nil || a.DurationUntilNextTrigger() < closest.DurationUntilNextTrigger()) {
			closest = a
		}
	}

	return closest, nil
}

func (am *AlarmService) listen() {
	for {
		select {
		case <-am.alarmTimer.C:
			go am.next.Trigger()
			am.UpdateTimings()
		default:
			// log.Println(am.next.DurationUntilNextAlarm().String())
			// time.Sleep(150 * time.Millisecond)
			continue
		}
	}
}

func (am *AlarmService) AddAlarm(alarm *alarm.Alarm) error {
	if _, err := am.GetAlarm(alarm.ID); err == nil {
		return fmt.Errorf("could not add alarm with id %v - it's already in list", alarm.ID.String())
	}

	am.Alarms = append(am.Alarms, alarm)
	err := am.UpdateTimings()
	if err != nil {
		return err
	}

	return am.Save()
}

func (am *AlarmService) GetAlarm(id uuid.UUID) (*alarm.Alarm, error) {
	idx := slices.IndexFunc(am.Alarms, func(a *alarm.Alarm) bool { return a.ID == id })
	if idx == -1 {
		return nil, errAlarmNotFound
	}

	return am.Alarms[idx], nil
}

func (am *AlarmService) UpdateTimings() error {
	if len(am.Alarms) == 0 {
		return errors.New("there are no alarms")
	}

	next, err := am.getClosestAlarm()
	if err != nil {
		return err
	}

	if next == nil {
		// see if there was a timer running and stop it
		if am.alarmTimer != nil {
			am.alarmTimer.Stop()
		}
		// stop all timers
		return nil
	}
	am.next = next
	am.alarmTimer = time.NewTimer(am.next.DurationUntilNextTrigger())
	return nil
}

func (am *AlarmService) Save() error {
	dat, err := json.Marshal(am)
	if err != nil {
		return err
	}

	err = os.WriteFile(am.cfg.AlarmsFileName, dat, 07777)
	return err
}

func (al *AlarmService) GetAllAlarms() ([]*alarm.Alarm, error) {
	return al.Alarms, nil
}
