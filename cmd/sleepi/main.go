package main

import (
	"log"

	"github.com/lakrizz/sleepi/config"
	"github.com/lakrizz/sleepi/internal/runtime"
	"github.com/lakrizz/sleepi/web"
	"github.com/lakrizz/sleepi/web/app"
)

func main() {
	if err := run(); err != nil {
		log.Println(err)
	}
}

func run() error {
	log.Println("loading config...")
	cfg, err := config.GetConfig()
	if err != nil {
		return err
	}

	log.Println("initializing runtime...")
	rt, err := runtime.InitRuntime(cfg)
	if err != nil {
		return err
	}

	app, err := app.InitApp(rt)
	if err != nil {
		return err
	}

	return web.Serve(app, cfg)
}
