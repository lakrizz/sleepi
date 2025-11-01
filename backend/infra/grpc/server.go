package grpc

import (
	"log/slog"
	"net/http"

	"github.com/lakrizz/sleepi/api/protobuf/gen/v1/sleepiv1connect"
	"github.com/lakrizz/sleepi/infra/grpc/handlers"
)

type Server struct {
}

func New() (*Server, error) {
	alarmHandler := &handlers.AlarmHandler{}
	mux := http.NewServeMux()
	path, handler := sleepiv1connect.NewAlarmServiceHandler(alarmHandler)

	mux.Handle(path, handler)
	p := new(http.Protocols)

	p.SetHTTP1(true)
	// Use h2c so we can serve HTTP/2 without TLS.
	p.SetUnencryptedHTTP2(true)
	s := http.Server{
		Addr:      "localhost:8080",
		Handler:   mux,
		Protocols: p,
	}

	slog.Info("listening on 8080")
	return nil, s.ListenAndServe()
}
