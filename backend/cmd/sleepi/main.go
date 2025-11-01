package main

import (
	"log/slog"
	"os"
	"time"

	"github.com/lmittmann/tint"
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

}
