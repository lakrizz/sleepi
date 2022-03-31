package main

import (
	"log"

	"krizz.org/sleepi/internal/manager"
	"krizz.org/sleepi/web"
	"krizz.org/sleepi/web/app"
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
