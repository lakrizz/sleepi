package app

import (
	"net/http"

	"github.com/gorilla/mux"
	"github.com/unrolled/render"
)

var ren *render.Render

func InitRoutes(router *mux.Router, renderer *render.Render) {
	ren = renderer
	addRoutes(router)
}

func addRoutes(m *mux.Router) {
	m.HandleFunc("/", IndexHandler)
}

func IndexHandler(w http.ResponseWriter, r *http.Request) {
	ren.HTML(w, http.StatusOK, "main/main", nil)
}
