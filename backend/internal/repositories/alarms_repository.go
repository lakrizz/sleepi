package repositories

import (
	"context"
	"fmt"

	"github.com/lakrizz/sleepi/internal/domain/alarms/entities"
	"github.com/lakrizz/sleepi/internal/infra/db"
	"github.com/lakrizz/sleepi/internal/mapper"
)

type AlarmsRepository struct {
	Queries *db.Queries
}

func NewAlarmsRepository(queries *db.Queries) (*AlarmsRepository, error) {
	return &AlarmsRepository{
		Queries: queries,
	}, nil
}

func (ar *AlarmsRepository) ListAlarms(ctx context.Context) ([]*entities.Alarm, error) {
	dbAlarms, err := ar.Queries.ListAlarms(ctx)
	if err != nil {
		return nil, fmt.Errorf("could not retrieve alarms from db: %w", err)
	}

	// now convert them to the domain
	domainAlarms := make([]*entities.Alarm, len(dbAlarms))
	for _, v := range dbAlarms {
		domainAlarms = append(domainAlarms, mapper.AlarmDatabaseToDomain(&v))
	}

	return domainAlarms, nil
}
