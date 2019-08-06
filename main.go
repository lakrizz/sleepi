package main

import (
	"html/template"
	"log"
	"time"

	"github.com/unrolled/render"

	"./controllers/alarm"
	"./controllers/playlist"
	"./lib/http"
	"./models"
	"./modules/youtube"
)

func main() {

	yt, err := youtube.CreateYouTubeWrapper()
	if err != nil {
		panic(err)
	}
	go listen(yt.Messages)

	for _, song := range []string{"Bachelors Of Science - Song For Lovers", "Tomte - Korn und Sprite", "Kettcar - 48 Stunden", "American Football - Never Meant", "Pendulum - Tarantula"} {

		go func(song string) {
			log.Println("moin", song)
			list, err := yt.Search.SearchVideos(song)
			if err != nil {
				panic(err)
			}

			go yt.Downloader.Download(list[0])
		}(song)
	}

	for {
		// endless loop for reasons Oo
	}

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

func listen(msgs chan *models.Message) {
	for {
		select {
		case v := <-msgs:
			switch v.Status {
			case models.MSG_SEARCH:
				log.Printf("searching for: %s\n", v.Payload)
				break
			case models.MSG_SEARCH_DONE:
				log.Printf("found results for: %s\n", v.Payload)
				break
			case models.MSG_DOWNLOAD_STARTING:
				log.Printf("starting download for: %s\n", v.Payload)
				break
			case models.MSG_DOWNLOAD_FINISHED:
				log.Printf("finished download for: %s\n", v.Payload)
				break
			case models.MSG_DOWNLOAD_ERROR:
				log.Printf("error at download for: %s\n", v.Payload)
				break
			}

		}

	}

}
