package app

import (
	"fmt"
	"log"
	"net/http"

	"github.com/gorilla/mux"
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

func PlaylistsView(w http.ResponseWriter, r *http.Request) {
	// finding playlist
	vars := mux.Vars(r)

	pl, err := api.GetPlaylist(vars["id"])
	if err != nil {
		// raise error
		return
	}
	ren.HTML(w, http.StatusOK, "playlists/view", pl)
}

func PlaylistDeleteSongs(w http.ResponseWriter, r *http.Request) {
	err := r.ParseForm()
	if err != nil {
		// do stuff
		return
	}
	vars := mux.Vars(r)
	err = api.DeleteSongsFromPlaylist(vars["id"], r.Form["songs"])
	if err != nil {
		log.Println(err)
		return
	}

	http.Redirect(w, r, fmt.Sprintf("/playlists/%s", vars["id"]), http.StatusPermanentRedirect)
}
