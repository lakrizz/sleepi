package alarms

import (
	"context"
	"fmt"

	"github.com/google/uuid"

	connect_go "github.com/bufbuild/connect-go"

	sleepiv1 "github.com/lakrizz/sleepi/api/protobuf/gen/v1"
	"github.com/lakrizz/sleepi/internal/domain/shared"
	"github.com/lakrizz/sleepi/internal/usecases/dto"
)

func (al *AlarmHandler) CreateAlarm(ctx context.Context, req *connect_go.Request[sleepiv1.CreateAlarmRequest]) (*connect_go.Response[sleepiv1.CreateAlarmResponse], error) {
	if req.Msg.Alarm == nil {
		return nil, fmt.Errorf("could not create alarm: no parameters given")
	}

	timeOfDay, err := shared.ParseTimeOfDay(req.Msg.Alarm.Time)
	if err != nil {
		return nil, fmt.Errorf("could not parse time of day: %w", err)
	}

	playableId, err := uuid.Parse(req.Msg.Alarm.PlayableId)
	if err != nil {
		return nil, fmt.Errorf("could not parse playable uuid: %w", err)
	}

	cmd := dto.AlarmAddCommand{
		Label:          req.Msg.Alarm.Label,
		TimeOfDay:      timeOfDay,
		Enabled:        req.Msg.Alarm.Enabled,
		WarmupDuration: req.Msg.Alarm.WarmupDuration,
		LEDTarget:      shared.RGB{},
		Playable:       playableId,
		Repeat:         []shared.Weekday{},
	}

	alarm, err := al.usecases.AddAlarm(ctx, cmd)
	if err != nil {
		return nil, fmt.Errorf("could not add alarm: %w", err)
	}

	req.Msg.Alarm.Id = shared.ToPtr(string(alarm.ID))

	return connect_go.NewResponse(&sleepiv1.CreateAlarmResponse{Alarm: req.Msg.Alarm}), nil
}
