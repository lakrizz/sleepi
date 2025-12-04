package alarms

import (
	"context"
	"fmt"

	"github.com/lakrizz/sleepi/internal/domain/alarms/entities"
	"github.com/lakrizz/sleepi/internal/mapper"
)

func (auc *AlarmsUseCases) ListAlarms(ctx context.Context) ([]*entities.Alarm, error) {
	dbAlarms, err := auc.AlarmsRepository.ListAlarms(ctx)
	if err != nil {
		return nil, fmt.Errorf("could not list alarms: %w", err)
	}

	// now convert them to the domain
	domainAlarms := make([]*entities.Alarm, len(dbAlarms))
	for i, v := range dbAlarms {
		domainAlarms[i] = mapper.AlarmDatabaseToDomain(&v)
	}
	return domainAlarms, nil
}
