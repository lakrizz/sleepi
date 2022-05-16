package app

import (
	"errors"
	"html/template"

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
		r.addAlarmRoutes,
		r.addLibraryRoutes,
		r.addPlayerRoutes,
		r.addPlaylistRoutes,
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

func (a *App) GetFuncMap() template.FuncMap {
	return template.FuncMap(map[string]interface{}{
		"TopFiles": func(slice []*library.File, num int) []*library.File {
			if len(slice) > num {
				return slice[:num]
			}
			return slice
		},
		"Cut": func(s string, i int) string {
			if len(s) <= i {
				return s
			}
			return s[:i]
		},
	})
}
