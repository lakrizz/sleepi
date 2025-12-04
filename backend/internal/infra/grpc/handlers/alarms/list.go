package alarms

import (
	"context"

	connect_go "github.com/bufbuild/connect-go"

	sleepiv1 "github.com/lakrizz/sleepi/api/protobuf/gen/v1"
	"github.com/lakrizz/sleepi/internal/mapper"
)

func (al *AlarmHandler) ListAlarms(ctx context.Context, _ *connect_go.Request[sleepiv1.ListAlarmsRequest]) (*connect_go.Response[sleepiv1.ListAlarmsResponse], error) {
	alarms, err := al.usecases.ListAlarms(ctx)
	if err != nil {
		return nil, err
	}

	returnAlarms := make([]*sleepiv1.Alarm, len(alarms))
	for i, v := range alarms {

		returnAlarms[i] = mapper.AlarmDomainToAPI(v)
	}

	resp := connect_go.NewResponse(&sleepiv1.ListAlarmsResponse{
		Alarms: returnAlarms,
	})

	return resp, nil
}
