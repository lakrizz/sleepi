package app

import (
	"net/http"

	"github.com/google/uuid"
	"github.com/gorilla/mux"
)

func LibraryHome(w http.ResponseWriter, r *http.Request) {
	ren.HTML(w, http.StatusOK, "library/main", api.Library)
}

func LibraryDelete(w http.ResponseWriter, r *http.Request) {
	// delete the song here and redirect
	// first find the song, then delete it from all playlists
	vars := mux.Vars(r)

	for _, pl := range api.Playlists.Playlists {
		err := api.DeleteSongsFromPlaylist(pl.Id.String(), []string{vars["id"]})
		if err != nil {
			ren.Text(w, http.StatusBadRequest, err.Error())
			return
		}
	}

	// now all pls should be without the song, do the removal stuff
	uid, err := uuid.Parse(vars["id"])
	if err != nil {
		ren.Text(w, http.StatusBadRequest, err.Error())
		return
	}
	err = api.Library.RemoveFile(uid)

	if err != nil {
		ren.Text(w, http.StatusBadRequest, err.Error())
		return
	}

	http.Redirect(w, r, "library/main", http.StatusPermanentRedirect)
}
