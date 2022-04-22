package app

import (
	"errors"
	"fmt"
	"log"
	"net/http"

	"github.com/google/uuid"
	"github.com/gorilla/mux"
)

func (r *Routes) addPlayerRoutes() error {
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
	id := mux.Vars(r)["id"]
	uu_id, err := uuid.Parse(id)
	if err != nil {
		log.Println(err)
		routes.ren.Data(w, http.StatusBadRequest, []byte(err.Error()))
		return
	}

	file := routes.api.library.Files[uu_id]
	routes.api.player.Clear()
	err = routes.api.player.Add(file)
	if err != nil {
		log.Println(err)
		routes.ren.Data(w, http.StatusBadRequest, []byte(err.Error()))
		return
	}
	routes.api.player.Play()
	http.Redirect(routes.withoutFrontendCache(w), r, r.Referer(), http.StatusPermanentRedirect)
}
