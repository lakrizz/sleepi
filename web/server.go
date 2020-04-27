package web

import (
	"log"
	"net/http"
	"os"
	"path"
	"time"

	"github.com/gorilla/mux"
	"github.com/unrolled/render"
)

var ren *render.Render

func Serve() {

	dir, err := os.Getwd()
	if err != nil {
		log.Fatal(err)
	}

	log.Println(dir)
	ren = render.New(render.Options{
		Directory:     path.Join(dir, "web", "templates"),
		IsDevelopment: true,
		Layout:        "layout",
	})

	m := mux.NewRouter()
	m.HandleFunc("/", Index)
	m.HandleFunc("", Index)
	m.PathPrefix("/static/").Handler(http.StripPrefix("/static/", http.FileServer(http.Dir("./web/static"))))

	server := &http.Server{
		Handler:      m,
		Addr:         ":8080",
		WriteTimeout: 15 * time.Second,
		ReadTimeout:  15 * time.Second,
	}

	log.Println("[frontend] listening on :8080")
	panic(server.ListenAndServe())
}

func Index(w http.ResponseWriter, r *http.Request) {
	ren.HTML(w, http.StatusOK, "alarms/main", nil)
}
