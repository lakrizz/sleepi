package main

import (
	"log"

	"github.com/unrolled/render"

	"./lib/http"
	"./modules/alarm"
)

func main() {
	// Initialize Server
	s, err := http.NewServer("8080")
	if err != nil {
		panic(err)
	}

	// Load Modules and add them
	render := render.New(render.Options{
		Directory: "./pub/",
		Layout:    "layout",
	})
	s.AddModule(alarm.AlarmModule(render))

	// Start Server
	log.Panic(s.Start())
}
