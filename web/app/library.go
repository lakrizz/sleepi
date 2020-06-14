package app

import "net/http"

func LibraryHome(w http.ResponseWriter, r *http.Request) {
	ren.HTML(w, http.StatusOK, "library/main", api.Library)
}
