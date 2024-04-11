package main

import (
	"context"
	"log"

	"github.com/lakrizz/sleepi/config"
	"github.com/lakrizz/sleepi/internal/runtime"
	"github.com/lakrizz/sleepi/web"
	"github.com/lakrizz/sleepi/web/app"
)

func main() {
	ctx := context.Background()
	if err := run(ctx); err != nil {
		log.Println(err)
	}
}

func run(ctx context.Context) error {
	log.Println("loading config...")
	cfg, err := config.GetConfig()
	if err != nil {
		return err
	}

	log.Println("initializing runtime...")
	rt, err := runtime.InitRuntime(ctx, cfg)
	if err != nil {
		return err
	}

	app, err := app.InitApp(rt)
	if err != nil {
		return err
	}

	return web.Serve(app, cfg)
}
