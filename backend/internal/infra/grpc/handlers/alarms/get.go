package alarms

import (
	"context"

	connect_go "github.com/bufbuild/connect-go"

	sleepiv1 "github.com/lakrizz/sleepi/api/protobuf/gen/v1"
	"github.com/lakrizz/sleepi/internal/mapper"
)

func (al *AlarmHandler) GetAlarm(ctx context.Context, req *connect_go.Request[sleepiv1.GetAlarmRequest]) (*connect_go.Response[sleepiv1.GetAlarmResponse], error) {
	alarm, err := al.usecases.AlarmsRepository.GetAlarm(ctx, req.Msg.Id)
	if err != nil {
		return nil, err
	}

	apiAlarm := mapper.AlarmDomainToAPI(alarm)
	resp := &sleepiv1.GetAlarmResponse{
		Alarm: apiAlarm,
	}
	return connect_go.NewResponse(resp), nil
}
