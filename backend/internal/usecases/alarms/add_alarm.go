package alarms

import (
	"context"
	"fmt"

	"github.com/lakrizz/sleepi/internal/domain/alarms/entities"
	"github.com/lakrizz/sleepi/internal/usecases/dto"
)

func (auc *AlarmsUseCases) AddAlarm(ctx context.Context, alarm dto.AlarmAddCommand) (*entities.Alarm, error) {
	alarmEntity, err := entities.NewAlarm(
		alarm.Label,
		alarm.TimeOfDay,
		alarm.Playable,
		alarm.Repeat,
		alarm.WarmupDuration,
		alarm.LEDTarget,
	)
	if err != nil {
		return nil, fmt.Errorf("could not create new alarm: %w", err)
	}

	domainAlarm, err := auc.AlarmsRepository.AddAlarm(ctx, alarmEntity)
	if err != nil {
		return nil, fmt.Errorf("could not add new alarm: %w", err)
	}

	return domainAlarm, nil
}
