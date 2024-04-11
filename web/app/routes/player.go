package routes

import (
	"errors"
	"fmt"
	"log"
	"net/http"

	"github.com/google/uuid"
	"github.com/gorilla/mux"
)

func (r *Routes) AddPlayerRoutes() error {
	if r == nil {
		return errors.New("routes is null")
	}
	prefix := "/player"
	routes := map[string]func(http.ResponseWriter, *http.Request){
		"/play/{id}": r.PlayerPlay,
	}
	for url, fn := range routes {
		u := fmt.Sprintf("%v%v", prefix, url)
		r.m.HandleFunc(u, fn)
	}
	return nil
}

func (routes *Routes) PlayerPlay(w http.ResponseWriter, r *http.Request) {
	id, err := uuid.Parse(mux.Vars(r)["id"])
	if err != nil {
		log.Println(err)
		routes.ren.Data(w, http.StatusBadRequest, []byte(err.Error()))
		return
	}

	file, err := routes.rt.Library.GetFileByID(id)
	if err != nil {
		return
	}

	routes.rt.Player.Clear()

	err = routes.rt.Player.Queue(file)
	if err != nil {
		log.Println("player/add", err)
		routes.ren.Data(w, http.StatusBadRequest, []byte(err.Error()))
		return
	}

	go routes.rt.Player.Play()
	http.Redirect(routes.withoutFrontendCache(w), r, r.Referer(), http.StatusPermanentRedirect)
}
