package alarms

import (
	"context"

	connect_go "github.com/bufbuild/connect-go"

	sleepiv1 "github.com/lakrizz/sleepi/api/protobuf/gen/v1"
	"github.com/lakrizz/sleepi/internal/domain/alarms/entities"
)

func (al *AlarmHandler) DeleteAlarm(ctx context.Context, req *connect_go.Request[sleepiv1.DeleteAlarmRequest]) (*connect_go.Response[sleepiv1.DeleteAlarmResponse], error) {
	err := al.usecases.DeleteAlarm(ctx, entities.AlarmID(req.Msg.Id))
	if err != nil {
		return nil, err
	}

	return connect_go.NewResponse(&sleepiv1.DeleteAlarmResponse{}), nil
}
