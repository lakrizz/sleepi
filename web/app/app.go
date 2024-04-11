package app

import (
	"github.com/gorilla/mux"
	"github.com/unrolled/render"

	"github.com/lakrizz/sleepi/internal/runtime"
	"github.com/lakrizz/sleepi/web/app/routes"
)

type App struct {
	rt *runtime.Runtime
}

func InitApp(rt *runtime.Runtime) (*App, error) {
	app := &App{rt: rt}
	return app, nil
}

func (a *App) InitRoutes(mux *mux.Router, ren *render.Render) error {
	r, err := routes.CreateRouter(a.rt, mux, ren)
	if err != nil {
		return err
	}

	routes := []func() error{
		r.AddAlarmRoutes,
		r.AddLibraryRoutes,
		r.AddPlayerRoutes,
		r.AddPlaylistRoutes,
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
