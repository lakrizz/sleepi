package app

import (
	"errors"

	"github.com/gorilla/mux"
	"github.com/unrolled/render"
	"krizz.org/sleepi/internal/manager"
	"krizz.org/sleepi/pkg/audioplayer"
	"krizz.org/sleepi/pkg/library"
)

type App struct {
	alarms    *manager.AlarmManager
	playlists *manager.PlaylistManager
	library   *library.Library
	player    *audioplayer.Audioplayer
}

func InitApp(man *manager.Managers) (*App, error) {
	if man == nil {
		return nil, errors.New("managers not initialized")
	}

	app := &App{}
	app.alarms = man.Alarms
	app.playlists = man.Playlists
	app.library = man.Library
	app.player = man.AudioPlayer
	return app, nil
}

func (a *App) InitRoutes(mux *mux.Router, ren *render.Render) error {
	r := &Routes{a, mux, ren}
	routes := []func() error{
		// r.addAlarmRoutes,
		r.addLibraryRoutes,
		r.addPlayerRoutes,
	}

	for _, r := range routes {
		err := r()
		if err != nil {
			return err
		}
	}

	if err := r.Debug(); err != nil {
		return err
	}
	return nil
}
