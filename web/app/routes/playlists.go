package routes

import (
	"errors"
	"fmt"
	"log"
	"net/http"

	"github.com/google/uuid"
	"github.com/gorilla/mux"

	"github.com/lakrizz/sleepi/pkg/models"
	"github.com/lakrizz/sleepi/pkg/playlist"
	"github.com/lakrizz/sleepi/pkg/util"
)

func (r *Routes) AddPlaylistRoutes() error {
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
	playlists, err := routes.rt.Playlists.GetAllPlaylists()
	if err != nil {
		return
	}

	params := map[string]any{
		"Playlists": playlists,
	}

	routes.ren.HTML(w, http.StatusOK, "playlists/main", params)
}

func (routes *Routes) PlaylistNew(w http.ResponseWriter, r *http.Request) {
	params := map[string]any{
		"Songs": routes.rt.Library.GetFiles(),
	}

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
		f, err := routes.rt.Library.GetFileByID(id)
		if err != nil {
			return
		}
		pl.AddFile(f)
	}

	if pl.Valid() {
		routes.rt.Playlists.AddPlaylist(pl)
	}

	http.Redirect(routes.withoutFrontendCache(w), r, "/playlists/", http.StatusPermanentRedirect)
}

func (routes *Routes) PlaylistEdit(w http.ResponseWriter, r *http.Request) {
	if r.Method != "GET" {
		routes.ren.Data(w, http.StatusMethodNotAllowed, []byte("this is not a get request"))
		return
	}

	id, err := uuid.Parse(mux.Vars(r)["id"])
	if err != nil {
		routes.ren.Data(w, http.StatusMethodNotAllowed, []byte(err.Error()))
		return
	}

	pl, err := routes.rt.Playlists.GetPlaylistById(id)
	if err != nil {
		routes.ren.Data(w, http.StatusMethodNotAllowed, []byte(err.Error()))
		return
	}

	params := make(map[string]any)

	params["Songs"] = util.IntersectFiles(routes.rt.Library.GetFiles(), pl.Files, func(a, b *models.File) bool { return a.ID == b.ID })
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

	id, err := uuid.Parse(mux.Vars(r)["id"])
	if err != nil {
		routes.ren.Data(w, http.StatusMethodNotAllowed, []byte(err.Error()))
		return
	}

	pl, err := routes.rt.Playlists.GetPlaylistById(id)
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
		f, err := routes.rt.Library.GetFileByID(id)
		if err != nil {
			return
		}

		pl.AddFile(f)
	}

	pl.Name = r.PostFormValue("playlist-title")

	http.Redirect(routes.withoutFrontendCache(w), r, "/playlists/", http.StatusPermanentRedirect)
}
