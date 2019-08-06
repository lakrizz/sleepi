package main

import (
	"html/template"
	"log"
	"time"

	"github.com/unrolled/render"

	"./controllers/alarm"
	"./controllers/playlist"
	"./lib/http"
	"./modules/youtube"
)

func main() {

	yt, err := youtube.CreateYouTubeWrapper()
	if err != nil {
		panic(err)
	}

	list, err := yt.Search.SearchVideos("Bachelors Of Science - Song For Lovers")
	if err != nil {
		panic(err)
	}

	err = yt.Downloader.Download(list[0])
	if err != nil {
		panic(err)
	}

	panic("lol")

	// Initialize Server
	s, err := http.NewServer("8080")
	if err != nil {
		panic(err)
	}

	// Load Modules and add them
	render := render.New(render.Options{
		Directory:     "./pub/",
		Layout:        "layout",
		IsDevelopment: true,
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
	s.AddModule(playlist.PlaylistModule(render))

	// Start Server
	log.Panic(s.Start())
}
