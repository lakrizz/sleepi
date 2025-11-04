package handlers

import (
	"context"
	"log/slog"

	connect_go "github.com/bufbuild/connect-go"
	"github.com/k0kubun/pp"

	sleepiv1 "github.com/lakrizz/sleepi/api/protobuf/gen/v1"
	"github.com/lakrizz/sleepi/api/protobuf/gen/v1/sleepiv1connect"
	"github.com/lakrizz/sleepi/internal/infra/grpc"
	"github.com/lakrizz/sleepi/internal/usecases/alarms"
)

type AlarmHandler struct {
	usecases *alarms.AlarmsUseCases
}

func RegisterAlarmHandler(server *grpc.Server, usecases *alarms.AlarmsUseCases) error {
	h := &AlarmHandler{
		usecases: usecases,
	}

	// directly register on the concrete server
	server.RegisterHandler(sleepiv1connect.NewAlarmServiceHandler(h))
	slog.Info("registered alarmhandler")

	return nil
}

var _ sleepiv1connect.AlarmServiceHandler = (*AlarmHandler)(nil)

func (al *AlarmHandler) ListAlarms(_ context.Context, req *connect_go.Request[sleepiv1.ListAlarmsRequest]) (*connect_go.Response[sleepiv1.ListAlarmsResponse], error) {
	pp.Println(req)
	slog.Info("listalarms called", "req", req)
	return nil, nil
}

func (al *AlarmHandler) GetAlarm(_ context.Context, _ *connect_go.Request[sleepiv1.GetAlarmRequest]) (*connect_go.Response[sleepiv1.GetAlarmResponse], error) {
	panic("not implemented") // TODO: Implement
}

func (al *AlarmHandler) CreateAlarm(_ context.Context, _ *connect_go.Request[sleepiv1.CreateAlarmRequest]) (*connect_go.Response[sleepiv1.CreateAlarmResponse], error) {
	panic("not implemented") // TODO: Implement
}

func (al *AlarmHandler) UpdateAlarm(_ context.Context, _ *connect_go.Request[sleepiv1.UpdateAlarmRequest]) (*connect_go.Response[sleepiv1.UpdateAlarmResponse], error) {
	panic("not implemented") // TODO: Implement
}

func (al *AlarmHandler) DeleteAlarm(_ context.Context, _ *connect_go.Request[sleepiv1.DeleteAlarmRequest]) (*connect_go.Response[sleepiv1.DeleteAlarmResponse], error) {
	panic("not implemented") // TODO: Implement
}

func (al *AlarmHandler) ToggleAlarm(_ context.Context, _ *connect_go.Request[sleepiv1.ToggleAlarmRequest]) (*connect_go.Response[sleepiv1.ToggleAlarmResponse], error) {
	panic("not implemented") // TODO: Implement
}
