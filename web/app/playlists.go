package app

import (
	"fmt"
	"log"
	"net/http"
	"path"

	"github.com/gorilla/mux"
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
	ren.HTML(w, http.StatusOK, "playlists/new", nil)
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
			Path     string
			Name     string
			FullPath string
			Status   bool
		}
		Playlist *playlist.Playlist
	}{}

	for _, v := range pl.Songs {
		s := api.Library.GetFile(v)

		songs.Songs = append(songs.Songs, struct {
			Path, Name, FullPath string
			Status               bool
		}{
			Path:     s.Path,
			Name:     s.Filename,
			FullPath: path.Join(s.Path, s.Filename),
			Status:   s.Exists(),
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
	err = api.DeleteSongsFromPlaylist(vars["id"], r.Form["songs"])
	if err != nil {
		log.Println(err)
		return
	}

	http.Redirect(w, r, fmt.Sprintf("/playlists/%s", vars["id"]), http.StatusPermanentRedirect)
}
