package alarms

import (
	"context"

	"github.com/lakrizz/sleepi/internal/domain/alarms/entities"
)

func (auc *AlarmsUseCases) DeleteAlarm(ctx context.Context, id entities.AlarmID) error {
	return auc.AlarmsRepository.DeleteAlarm(ctx, id)
}
