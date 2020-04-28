package app

import (
	"github.com/gorilla/mux"
	base_api "github.com/lakrizz/sleepi/api"
	"github.com/unrolled/render"
)

var api *base_api.Api
var ren *render.Render

func Init(router *mux.Router, renderer *render.Render, api_v *base_api.Api) {
	api = api_v
	ren = renderer
	addRoutes(router)
}

func addRoutes(m *mux.Router) {
	m.HandleFunc("/alarms", AlarmsHome)
	m.HandleFunc("/alarms/", AlarmsHome)

	m.HandleFunc("/playlists", PlaylistsHome)
	m.HandleFunc("/playlists/", PlaylistsHome)
}
