package app

import (
	"errors"
	"fmt"
	"log"
	"net/http"

	"github.com/k0kubun/pp"
)

func (r *Routes) addPlaylistRoutes() error {
	if r == nil {
		return errors.New("routes is null")
	}
	prefix := "/playlists"
	routes := map[string]func(http.ResponseWriter, *http.Request){
		"/":    r.PlaylistIndex,
		"/new": r.PlaylistNew,
	}
	for url, fn := range routes {
		u := fmt.Sprintf("%v%v", prefix, url)
		r.m.HandleFunc(u, fn)
		log.Println(u)
	}
	return nil
}

func (routes *Routes) PlaylistIndex(w http.ResponseWriter, r *http.Request) {
	params := make(map[string]interface{})
	params["Playlists"] = routes.api.playlists.Playlists
	routes.ren.HTML(w, http.StatusOK, "playlists/main", params)
}

func (routes *Routes) PlaylistNew(w http.ResponseWriter, r *http.Request) {
	params := make(map[string]interface{})
	params["Songs"] = routes.api.library.GetAllFiles()
	pp.Println(params)
	routes.ren.HTML(w, http.StatusOK, "playlists/new", params)
}
