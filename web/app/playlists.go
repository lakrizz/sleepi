package app

import (
	"fmt"
	"log"
	"net/http"

	"github.com/google/uuid"
	"github.com/gorilla/mux"
	"github.com/k0kubun/pp"
	"github.com/lakrizz/sleepi/pkg/playlist"
)

func PlaylistsHome(w http.ResponseWriter, r *http.Request) {
	playlists, err := api.GetPlaylists()
	if err != nil {
		// show error page
		return
	} else {
		ren.HTML(w, http.StatusOK, "playlists/main", playlists)
	}
}

func PlaylistsNew(w http.ResponseWriter, r *http.Request) {
	songs := api.Library.Files
	ren.HTML(w, http.StatusOK, "playlists/new", songs)
}

func PlaylistsCreate(w http.ResponseWriter, r *http.Request) {
	err := r.ParseForm()
	if err != nil {
		return
	}

	plname := r.Form["playlist_name"][0]

	ids := make([]uuid.UUID, 0)
	for _, v := range r.Form["songs"] {
		id, _ := uuid.Parse(v)
		ids = append(ids, id)
	}

	err = api.Playlists.CreatePlaylist(plname, ids)
	if err != nil {
		pp.Println(err.Error())
		return
	}

	http.Redirect(w, r, "/playlists", http.StatusPermanentRedirect)
}

func PlaylistsView(w http.ResponseWriter, r *http.Request) {
	// finding playlist
	vars := mux.Vars(r)

	pl, err := api.GetPlaylist(vars["id"])
	if err != nil {
		// raise error
		return
	}

	// some eyecandy
	songs := struct {
		Songs []struct {
			Name   string
			Id     string
			Status bool
		}
		Playlist *playlist.Playlist
	}{}

	for _, v := range pl.Songs {
		s := api.Library.GetFile(v)

		songs.Songs = append(songs.Songs, struct {
			Name, Id string
			Status   bool
		}{
			Name:   s.Filename,
			Id:     s.Id.String(),
			Status: s.Exists(),
		})
	}
	songs.Playlist = pl

	ren.HTML(w, http.StatusOK, "playlists/view", songs)
}

func PlaylistDeleteSongs(w http.ResponseWriter, r *http.Request) {
	err := r.ParseForm()
	if err != nil {
		// do stuff
		return
	}
	vars := mux.Vars(r)
	pp.Println(r.Form)
	err = api.DeleteSongsFromPlaylist(vars["id"], r.Form["songs"])
	if err != nil {
		log.Println(err)
		return
	}

	http.Redirect(w, r, fmt.Sprintf("/playlists/%s", vars["id"]), http.StatusPermanentRedirect)
}
