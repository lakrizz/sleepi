package app

import (
	"net/http"

	"github.com/gorilla/mux"
	"github.com/unrolled/render"
)

var ren *render.Render
var m *mux.Router

func InitRoutes(router *mux.Router, renderer *render.Render) {
	ren = renderer
	m = router
	addRoutes()
	addAlarmRoutes()
}

func addRoutes() {
	m.HandleFunc("/", IndexHandler)
}

func IndexHandler(w http.ResponseWriter, r *http.Request) {
	ren.HTML(w, http.StatusOK, "main/main", nil)
}
