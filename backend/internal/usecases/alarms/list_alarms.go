package alarms

import (
	"context"
	"fmt"

	"github.com/k0kubun/pp"

	"github.com/lakrizz/sleepi/internal/domain/alarms/entities"
)

func (auc *AlarmsUseCases) ListAlarms(ctx context.Context) ([]*entities.Alarm, error) {
	alarms, err := auc.AlarmsRepository.ListAlarms(ctx)
	if err != nil {
		return nil, fmt.Errorf("could not list alarms: %w", err)
	}

	pp.Println(alarms)

	return nil, nil
}
