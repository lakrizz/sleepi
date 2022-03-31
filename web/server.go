package web

import (
	"fmt"
	"log"
	"net/http"
	"os"
	"path"
	"time"

	"github.com/gorilla/mux"
	"github.com/unrolled/render"
	"krizz.org/sleepi/config"
	"krizz.org/sleepi/web/app"
)

var ren *render.Render

func Serve() error {
	dir, err := os.Getwd()
	if err != nil {
		log.Fatal(err)
	}

	ren = render.New(render.Options{
		Directory:     path.Join(dir, "web", "templates"),
		IsDevelopment: true,
		Layout:        "layout",
	})

	m := mux.NewRouter()
	app.InitRoutes(m, ren)

	m.PathPrefix("/static/").Handler(http.StripPrefix("/static/", http.FileServer(http.Dir("./web/static"))))

	conf, err := config.GetConfig()
	if err != nil {
		panic(err)
	}

	server := &http.Server{
		Handler:      m,
		Addr:         fmt.Sprintf("%v:%v", conf.HTTP_HOST, conf.HTTP_PORT),
		WriteTimeout: 15 * time.Second,
		ReadTimeout:  15 * time.Second,
	}

	log.Printf("[frontend] listening on %v:%v", conf.HTTP_HOST, conf.HTTP_PORT)
	return server.ListenAndServe()
}
