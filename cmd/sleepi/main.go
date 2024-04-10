package main

import (
	"log"

	"github.com/lakrizz/sleepi/internal/manager"
	"github.com/lakrizz/sleepi/web"
	"github.com/lakrizz/sleepi/web/app"
)

func main() {
	if err := run(); err != nil {
		log.Println(err)
	}
}

func run() error {
	managers, err := manager.GetManagers()
	if err != nil {
		return err
	}
	app, err := app.InitApp(managers)
	if err != nil {
		return err
	}
	return web.Serve(app)
}
