package routes

import (
	"fmt"
	"net/http"

	"github.com/gorilla/mux"
	"github.com/unrolled/render"

	"github.com/lakrizz/sleepi/internal/runtime"
)

type Routes struct {
	rt  *runtime.Runtime
	m   *mux.Router
	ren *render.Render
}

func CreateRouter(rt *runtime.Runtime, mux *mux.Router, renderer *render.Render) (*Routes, error) {
  r := &Routes{
  	rt:  rt,
  	m:   mux,
  	ren: renderer,
  }

  return r, nil
}

func (r *Routes) Debug() error {
	return r.m.Walk(func(route *mux.Route, router *mux.Router, ancestors []*mux.Route) error {
		t, err := route.GetPathTemplate()
		if err != nil {
			return err
		}
		fmt.Println(t)
		return nil
	})
}

func (r *Routes) withoutFrontendCache(w http.ResponseWriter) http.ResponseWriter {
	w.Header().Set("Cache-Control", "no-cache, no-store, must-revalidate") // HTTP 1.1.
	w.Header().Set("Pragma", "no-cache")                                   // HTTP 1.0.
	w.Header().Set("Expires", "0")                                         // Proxies.
	return w
}
