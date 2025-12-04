package repositories

import (
	"context"
	"fmt"

	"github.com/lakrizz/sleepi/internal/domain/alarms/entities"
	"github.com/lakrizz/sleepi/internal/infra/db"
	"github.com/lakrizz/sleepi/internal/mapper"
)

type AlarmsRepository struct {
	queries *db.Queries
}

func NewAlarmsRepository(queries *db.Queries) (*AlarmsRepository, error) {
	return &AlarmsRepository{
		queries: queries,
	}, nil
}

func (ar *AlarmsRepository) ListAlarms(ctx context.Context) ([]db.Alarm, error) {
	dbAlarms, err := ar.queries.ListAlarms(ctx)
	if err != nil {
		return nil, fmt.Errorf("could not retrieve alarms from db: %w", err)
	}

	return dbAlarms, nil
}

func (ar *AlarmsRepository) AddAlarm(ctx context.Context, alarm *entities.Alarm) (*entities.Alarm, error) {
	dbAlarm := mapper.DomainAlarmToDatabase(alarm)

	params := db.CreateAlarmParams{
		ID:             dbAlarm.ID,
		Label:          dbAlarm.Label,
		Time:           dbAlarm.Time,
		Enabled:        dbAlarm.Enabled,
		WarmupDuration: dbAlarm.WarmupDuration,
		LedTarget:      dbAlarm.LedTarget,
		PlayableID:     dbAlarm.PlayableID,
		Weekdays:       dbAlarm.Weekdays,
	}
	err := ar.queries.CreateAlarm(ctx, params)
	if err != nil {
		return nil, fmt.Errorf("could not create alarm in db: %w", err)
	}

	return alarm, nil
}

func (ar AlarmsRepository) GetAlarm(ctx context.Context, id string) (*entities.Alarm, error) {
	dbAlarm, err := ar.queries.GetAlarm(ctx, id)
	if err != nil {
		return nil, err
	}

	return mapper.AlarmDatabaseToDomain(&dbAlarm), nil
}

func (ar *AlarmsRepository) UpdateAlarm(ctx context.Context, alarm *entities.Alarm) error {
	dbAlarm := mapper.DomainAlarmToDatabase(alarm)

	params := db.UpdateAlarmParams{
		Label:          dbAlarm.Label,
		Time:           dbAlarm.Time,
		Enabled:        dbAlarm.Enabled,
		WarmupDuration: dbAlarm.WarmupDuration,
		LedTarget:      dbAlarm.LedTarget,
		PlayableID:     dbAlarm.PlayableID,
		Weekdays:       dbAlarm.Weekdays,
		ID:             dbAlarm.ID,
	}

	return ar.queries.UpdateAlarm(ctx, params)
}

func (ar *AlarmsRepository) DeleteAlarm(ctx context.Context, id entities.AlarmID) error {
	return ar.queries.DeleteAlarm(ctx, string(id))
}
