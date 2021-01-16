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
	m.HandleFunc("/alarms/new", AlarmsNew)
	m.HandleFunc("/alarms/new/", AlarmsNew)
	m.HandleFunc("/alarms/create", AlarmsCreate)
	m.HandleFunc("/alarms/create/", AlarmsCreate)
	m.HandleFunc("/alarms/{id}", AlarmsView)
	m.HandleFunc("/alarms/{id}/disable", AlarmsDisable)
	m.HandleFunc("/alarms/{id}/enable", AlarmsEnable)
	m.HandleFunc("/alarms/{id}/edit", AlarmsEdit)

	m.HandleFunc("/playlists", PlaylistsHome)
	m.HandleFunc("/playlists/", PlaylistsHome)
	m.HandleFunc("/playlists/new", PlaylistsNew)
	m.HandleFunc("/playlists/create", PlaylistsNew)
	m.HandleFunc("/playlists/{id}", PlaylistsView)
	m.HandleFunc("/playlists/{id}/deletesongs", PlaylistDeleteSongs)

	m.HandleFunc("/library", LibraryHome)
	m.HandleFunc("/library/", LibraryHome)
	m.HandleFunc("/library/delete/{id}", LibraryDelete)
}
