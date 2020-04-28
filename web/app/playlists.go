package app

import (
	"net/http"
)

func PlaylistsHome(w http.ResponseWriter, r *http.Request) {
	playlists, err := api.GetPlaylists()
	if err != nil {
		// show error page
	} else {
		ren.HTML(w, http.StatusOK, "playlists/main", playlists)
	}
}
