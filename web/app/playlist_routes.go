package app

import (
	"errors"
	"fmt"
	"log"
	"net/http"

	"github.com/google/uuid"
	"github.com/k0kubun/pp"
	"krizz.org/sleepi/pkg/playlist"
)

func (r *Routes) addPlaylistRoutes() error {
	if r == nil {
		return errors.New("routes is null")
	}
	prefix := "/playlists"
	routes := map[string]func(http.ResponseWriter, *http.Request){
		"/":       r.PlaylistIndex,
		"/new":    r.PlaylistNew,
		"/create": r.PlaylistCreate,
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
	routes.ren.HTML(w, http.StatusOK, "playlists/new", params)
}

func (routes *Routes) PlaylistCreate(w http.ResponseWriter, r *http.Request) {
	if r.Method != "POST" {
		routes.ren.Data(w, http.StatusMethodNotAllowed, []byte("this is not a post request"))
		return
	}

	err := r.ParseForm()
	if err != nil {
		routes.ren.Data(w, http.StatusMethodNotAllowed, []byte(err.Error()))
		return
	}

	pl, err := playlist.NewPlaylist(r.PostFormValue("playlist-title"))
	if err != nil {
		routes.ren.Data(w, http.StatusMethodNotAllowed, []byte(err.Error()))
		return
	}
	for _, v := range []string(r.PostForm["order"]) {
		id, err := uuid.Parse(v)
		if err != nil {
			routes.ren.Data(w, http.StatusMethodNotAllowed, []byte(err.Error()))
			return
		}
		f := routes.api.library.Files[id]
		pl.Add(f)
	}

	pp.Println(pl)
}
