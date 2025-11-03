package main

import (
	"log/slog"

	"github.com/lakrizz/sleepi/internal/app"
)

func main() {
	a, err := app.New()
	if err != nil {
		slog.Error("could not initialize app", "error", err)
	}

	if err := a.Start(); err != nil {
		slog.Error("ungraceful shutdown", "reason", err)
	}

	slog.Info("graceful shutdown completed")
}
