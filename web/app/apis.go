package app

import (
	"errors"

	"github.com/gorilla/mux"
	"github.com/unrolled/render"
	"krizz.org/sleepi/internal/manager"
	"krizz.org/sleepi/pkg/library"
)

type App struct {
	alarms    *manager.AlarmManager
	playlists *manager.PlaylistManager
	library   *library.Library
}

func InitApp(man *manager.Managers) (*App, error) {
	if man == nil {
		return nil, errors.New("managers not initialized")
	}

	app := &App{}
	app.alarms = man.Alarms
	app.playlists = man.Playlists
	app.library = man.Library
	return app, nil
}

func (a *App) InitRoutes(mux *mux.Router, ren *render.Render) error {
	r := &Routes{a, mux, ren}

	err := r.addAlarmRoutes()
	if err != nil {
		return err
	}

	err = r.addLibraryRoutes()
	if err != nil {
		return err
	}

	if err = r.Debug(); err != nil {
		return err
	}
	return nil
}
