package main

import (
	"html/template"
	"log"
	"time"

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
		Funcs: []template.FuncMap{
			template.FuncMap{
				"join": func(days []time.Weekday) string {
					rets := ""
					for _, d := range days {
						rets = rets + ", " + d.String()[:3]
					}
					return rets[1:]
				},
			},
		},
	})
	s.AddModule(alarm.AlarmModule(render))

	// Start Server
	log.Panic(s.Start())
}
