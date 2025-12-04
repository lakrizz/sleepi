package alarms

import (
	"context"
	"reflect"
	"time"

	"github.com/lakrizz/sleepi/internal/domain/alarms/entities"
)

func (auc *AlarmsUseCases) UpdateAlarm(ctx context.Context, alarm *entities.Alarm) error {
	// first find the original alarm
	dbAlarm, err := auc.AlarmsRepository.GetAlarm(ctx, string(alarm.ID))
	if err != nil {
		return err
	}

	changed := false

	set := func(old, new any, apply func()) {
		if !reflect.DeepEqual(old, new) {
			apply()
			changed = true
		}
	}

	set(dbAlarm.Label, alarm.Label, func() { dbAlarm.Label = alarm.Label })
	set(dbAlarm.TimeOfDay, alarm.TimeOfDay, func() { dbAlarm.TimeOfDay = alarm.TimeOfDay })
	set(dbAlarm.Enabled, alarm.Enabled, func() { dbAlarm.Enabled = alarm.Enabled })
	set(dbAlarm.WarmupDuration, alarm.WarmupDuration, func() { dbAlarm.WarmupDuration = alarm.WarmupDuration })
	set(dbAlarm.LEDTarget, alarm.LEDTarget, func() { dbAlarm.LEDTarget = alarm.LEDTarget })
	set(dbAlarm.Playable, alarm.Playable, func() { dbAlarm.Playable = alarm.Playable })
	set(dbAlarm.Repeat, alarm.Repeat, func() { dbAlarm.Repeat = alarm.Repeat })
	// ID and CreatedAt typically never change

	if changed {
		dbAlarm.UpdatedAt = time.Now()
	}

	// persist it
	return auc.AlarmsRepository.UpdateAlarm(ctx, dbAlarm)
}
