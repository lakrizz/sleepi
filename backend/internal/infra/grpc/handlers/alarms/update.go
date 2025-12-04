package alarms

import (
	"context"

	connect_go "github.com/bufbuild/connect-go"

	sleepiv1 "github.com/lakrizz/sleepi/api/protobuf/gen/v1"
	"github.com/lakrizz/sleepi/internal/mapper"
)

func (al *AlarmHandler) UpdateAlarm(ctx context.Context, req *connect_go.Request[sleepiv1.UpdateAlarmRequest]) (*connect_go.Response[sleepiv1.UpdateAlarmResponse], error) {
	domainAlarm, err := mapper.AlarmAPIToDomain(req.Msg.Alarm)
	if err != nil {
		return nil, err
	}

	err = al.usecases.UpdateAlarm(ctx, domainAlarm)
	if err != nil {
		return nil, err
	}

	return connect_go.NewResponse(&sleepiv1.UpdateAlarmResponse{Alarm: req.Msg.Alarm}), nil
}
