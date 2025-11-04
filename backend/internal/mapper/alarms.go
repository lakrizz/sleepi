package mapper

import (
	"database/sql"
	"encoding/json"
	"log/slog"
	"time"

	"github.com/google/uuid"

	"github.com/lakrizz/sleepi/internal/domain/alarms/entities"
	"github.com/lakrizz/sleepi/internal/domain/shared"
	"github.com/lakrizz/sleepi/internal/infra/db"
)

func DomainAlarmToDatabase(domainAlarm *entities.Alarm) *db.Alarm {
	if domainAlarm == nil {
		return nil
	}

	// serialize JSON fields
	ledTargetBytes, err := json.Marshal(domainAlarm.LEDTarget)
	if err != nil {
		slog.Error("could not marshal led target", "id", domainAlarm.ID, "error", err)
		return nil
	}

	repeatDaysBytes, err := json.Marshal(domainAlarm.Repeat)
	if err != nil {
		slog.Error("could not marshal weekdays", "id", domainAlarm.ID, "error", err)
		return nil
	}

	return &db.Alarm{
		ID:      string(domainAlarm.ID),
		Label:   domainAlarm.Label,
		Time:    domainAlarm.TimeOfDay.String(),
		Enabled: domainAlarm.Enabled,
		WarmupDuration: sql.NullInt64{
			Int64: int64(domainAlarm.WarmupDuration),
			Valid: true,
		},
		LedTarget: sql.NullString{
			String: string(ledTargetBytes),
			Valid:  true,
		},
		PlayableID: sql.NullString{
			String: domainAlarm.Playable.String(),
			Valid:  true,
		},
		Weekdays: sql.NullString{
			String: string(repeatDaysBytes),
			Valid:  true,
		},
		CreatedAt: sql.NullTime{
			Time:  domainAlarm.CreatedAt,
			Valid: true,
		},
		UpdatedAt: sql.NullTime{
			Time:  domainAlarm.UpdatedAt,
			Valid: true,
		},
	}
}

func AlarmDatabaseToDomain(dbAlarm *db.Alarm) *entities.Alarm {
	timeOfDay, err := shared.ParseTimeOfDay(dbAlarm.Time)
	if err != nil {
		slog.Error("could not parse timeofday", "id", dbAlarm.ID, "error", err)
		return nil
	}

	warmupDuration := time.Duration(dbAlarm.WarmupDuration.Int64)
	ledTarget := &shared.RGB{}
	err = json.Unmarshal([]byte(dbAlarm.LedTarget.String), &ledTarget)
	if err != nil {
		slog.Error("could not unmarshal led target", "id", dbAlarm.ID, "error", err)
		return nil
	}

	playableId, err := uuid.Parse(dbAlarm.PlayableID.String)
	if err != nil {
		slog.Error("could not parse playable id", "id", dbAlarm.ID, "error", err)
		return nil
	}

	repeatDays := []shared.Weekday{}
	err = json.Unmarshal([]byte(dbAlarm.Weekdays.String), &repeatDays)
	if err != nil {
		slog.Error("could not unmarshal weekdays", "id", dbAlarm.ID, "error", err)
		return nil
	}

	return &entities.Alarm{
		ID:             entities.AlarmID(dbAlarm.ID),
		Label:          dbAlarm.Label,
		TimeOfDay:      timeOfDay,
		Enabled:        dbAlarm.Enabled,
		WarmupDuration: warmupDuration,
		LEDTarget:      ledTarget,
		Playable:       playableId,
		Repeat:         repeatDays,
		CreatedAt:      dbAlarm.CreatedAt.Time,
		UpdatedAt:      dbAlarm.UpdatedAt.Time,
	}
}
