package alarms

import (
	"log/slog"

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
