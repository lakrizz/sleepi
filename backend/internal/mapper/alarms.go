package mapper

import (
	"database/sql"
	"encoding/json"
	"fmt"
	"log/slog"

	"github.com/google/uuid"

	sleepiv1 "github.com/lakrizz/sleepi/api/protobuf/gen/v1"
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
		WarmupDuration: sql.NullString{
			String: domainAlarm.WarmupDuration,
			Valid:  true,
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

func AlarmDomainToAPI(domainAlarm *entities.Alarm) *sleepiv1.Alarm {
	parsedUuid, err := uuid.Parse(string(domainAlarm.ID))
	if err != nil {
		slog.Error("could not parse uuid", "error", err)
		return nil
	}

	repeatDays := make([]sleepiv1.Weekday, len(domainAlarm.Repeat))
	for i, v := range domainAlarm.Repeat {
		repeatDays[i] = sleepiv1.Weekday(v)
	}

	a := &sleepiv1.Alarm{
		Id:             shared.ToPtr(string(domainAlarm.ID)),
		Label:          domainAlarm.Label,
		Time:           domainAlarm.TimeOfDay.String(),
		RepeatDays:     repeatDays,
		Enabled:        domainAlarm.Enabled,
		WarmupDuration: domainAlarm.WarmupDuration,
		LedTarget:      &sleepiv1.RGB{R: uint32(domainAlarm.LEDTarget.R), G: uint32(domainAlarm.LEDTarget.G), B: uint32(domainAlarm.LEDTarget.B)},
		PlayableId:     parsedUuid.String(),
	}

	return a
}

func AlarmDatabaseToDomain(dbAlarm *db.Alarm) *entities.Alarm {
	timeOfDay, err := shared.ParseTimeOfDay(dbAlarm.Time)
	if err != nil {
		slog.Error("could not parse timeofday", "id", dbAlarm.ID, "error", err)
		return nil
	}

	ledTarget := shared.RGB{}
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
		WarmupDuration: dbAlarm.WarmupDuration.String,
		LEDTarget:      ledTarget,
		Playable:       playableId,
		Repeat:         repeatDays,
		CreatedAt:      dbAlarm.CreatedAt.Time,
		UpdatedAt:      dbAlarm.UpdatedAt.Time,
	}
}

func AlarmAPIToDomain(apiAlarm *sleepiv1.Alarm) (*entities.Alarm, error) {
	if apiAlarm == nil {
		return nil, fmt.Errorf("nil apiAlarm")
	}

	// Parse ID
	id := entities.AlarmID(*apiAlarm.Id)

	// Parse PlayableId (uuid)
	_, err := uuid.Parse(apiAlarm.PlayableId)
	if err != nil {
		return nil, fmt.Errorf("invalid playableId: %w", err)
	}

	// Repeat days
	repeat := make([]int, len(apiAlarm.RepeatDays))
	for i, v := range apiAlarm.RepeatDays {
		repeat[i] = int(v)
	}

	// Time of day
	tod, err := shared.ParseTimeOfDay(apiAlarm.Time)
	if err != nil {
		return nil, fmt.Errorf("invalid time-of-day: %w", err)
	}

	// LED
	var led shared.RGB
	if apiAlarm.LedTarget != nil {
		led = shared.RGB{
			R: uint8(apiAlarm.LedTarget.R),
			G: uint8(apiAlarm.LedTarget.G),
			B: uint8(apiAlarm.LedTarget.B),
		}
	}

	playableId, err := uuid.Parse(apiAlarm.PlayableId)
	if err != nil {
		return nil, err
	}

	return &entities.Alarm{
		ID:             id,
		Label:          apiAlarm.Label,
		TimeOfDay:      tod,
		Repeat:         shared.ParseWeekdays(repeat),
		Enabled:        apiAlarm.Enabled,
		WarmupDuration: apiAlarm.WarmupDuration,
		LEDTarget:      led,
		Playable:       playableId,
	}, nil
}
