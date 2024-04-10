package app

import (
	"errors"
	"fmt"
	"log"
	"net/http"

	"github.com/google/uuid"
	"github.com/gorilla/mux"
	"github.com/lakrizz/sleepi/pkg/library"
	"github.com/lakrizz/sleepi/pkg/playlist"
	"github.com/lakrizz/sleepi/pkg/util"
)

func (r *Routes) addPlaylistRoutes() error {
	if r == nil {
		return errors.New("routes is null")
	}
	prefix := "/playlists"
	routes := map[string]func(http.ResponseWriter, *http.Request){
		"/":          r.PlaylistMain,
		"/new":       r.PlaylistNew,
		"/create":    r.PlaylistCreate,
		"/{id}":      r.PlaylistEdit,
		"/{id}/save": r.PlaylistSave,
	}
	for url, fn := range routes {
		u := fmt.Sprintf("%v%v", prefix, url)
		r.m.HandleFunc(u, fn)
		log.Println(u)
	}
	return nil
}

func (routes *Routes) PlaylistMain(w http.ResponseWriter, r *http.Request) {
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
		pl.AddFile(f)
	}

	if pl.Valid() {
		routes.api.playlists.AddPlaylist(pl)
	}

	http.Redirect(routes.withoutFrontendCache(w), r, "/playlists/", http.StatusPermanentRedirect)
}

func (routes *Routes) PlaylistEdit(w http.ResponseWriter, r *http.Request) {
	if r.Method != "GET" {
		routes.ren.Data(w, http.StatusMethodNotAllowed, []byte("this is not a get request"))
		return
	}

	id_s := mux.Vars(r)["id"]
	id, err := uuid.Parse(id_s)
	if err != nil {
		routes.ren.Data(w, http.StatusMethodNotAllowed, []byte(err.Error()))
		return
	}

	pl, err := routes.api.playlists.GetPlaylistById(id)
	if err != nil {
		routes.ren.Data(w, http.StatusMethodNotAllowed, []byte(err.Error()))
		return
	}

	params := make(map[string]interface{})

	params["Songs"] = util.IntersectFiles(routes.api.library.GetAllFiles(), pl.Files, func(a, b *library.File) bool { return a.Id == b.Id })
	params["Playlist"] = pl

	routes.ren.HTML(w, http.StatusOK, "playlists/edit", params)
}

func (routes *Routes) PlaylistSave(w http.ResponseWriter, r *http.Request) {
	if r.Method != "POST" {
		routes.ren.Data(w, http.StatusMethodNotAllowed, []byte("this is not a post request"))
		return
	}

	err := r.ParseForm()
	if err != nil {
		routes.ren.Data(w, http.StatusMethodNotAllowed, []byte(err.Error()))
		return
	}

	id_s := mux.Vars(r)["id"]
	id, err := uuid.Parse(id_s)
	if err != nil {
		routes.ren.Data(w, http.StatusMethodNotAllowed, []byte(err.Error()))
		return
	}

	pl, err := routes.api.playlists.GetPlaylistById(id)
	if err != nil {
		routes.ren.Data(w, http.StatusMethodNotAllowed, []byte(err.Error()))
		return
	}

	err = pl.ClearFiles()
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
		pl.AddFile(f)
	}

	pl.Name = r.PostFormValue("playlist-title")

	if pl.Valid() {
		routes.api.playlists.UpdatePlaylist(pl)
	}

	http.Redirect(routes.withoutFrontendCache(w), r, "/playlists/", http.StatusPermanentRedirect)
}
