package app

import (
	"database/sql"
	"fmt"
	"log/slog"
	"os"
	"time"

	"github.com/lmittmann/tint"

	"github.com/lakrizz/sleepi/internal/infra/grpc"
	"github.com/lakrizz/sleepi/internal/infra/grpc/handlers"
)

type App struct {
	Server *grpc.Server
	DB     *sql.DB
}

func New() (*App, error) {
	app := &App{}

	// here we initialize everything
	// logger first, we set a default slog so we don't care about passing instances
	lg := slog.New(
		tint.NewHandler(os.Stdout, &tint.Options{
			AddSource:  true,
			Level:      slog.LevelDebug,
			TimeFormat: time.Kitchen,
		}),
	)
	slog.SetDefault(lg)

	server, err := grpc.New()
	if err != nil {
		return nil, err
	}

	app.Server = server

	// database stuff
	err = app.initDatabase()
	if err != nil {
		return nil, err
	}

	// instantiate services

	// create usecase instances

	// instantiate handlers
	handlers.RegisterAlarmHandler(app.Server)

	// debug?
	app.Server.DebugRoutes()

	return app, nil
}

func (a *App) Start() error {
	if a.Server == nil {
		return fmt.Errorf("server is nil")
	}

	return a.Server.Start()
}
