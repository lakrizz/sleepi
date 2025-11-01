package main

import (
	"log/slog"
	"os"
	"time"

	"github.com/lmittmann/tint"

	"github.com/lakrizz/sleepi/infra/grpc"
)

func main() {
	lg := slog.New(
		tint.NewHandler(os.Stdout, &tint.Options{
			AddSource:  true,
			Level:      slog.LevelDebug,
			TimeFormat: time.Kitchen,
		}),
	)
	slog.SetDefault(lg)

	slog.Info("HI!")

	grpc.New()
}
