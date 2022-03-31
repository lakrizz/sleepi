package app

import (
	"fmt"

	"github.com/gorilla/mux"
	"github.com/unrolled/render"
)

type Routes struct {
	api *App
	m   *mux.Router
	ren *render.Render
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
