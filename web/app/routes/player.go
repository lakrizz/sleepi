package routes

import (
	"errors"
	"fmt"
	"net/http"
)

func (r *Routes) AddPlayerRoutes() error {
	if r == nil {
		return errors.New("routes is null")
	}
	prefix := "/player"
	routes := map[string]func(http.ResponseWriter, *http.Request){
		"/play/{id}": r.PlayerPlay,
	}
	for url, fn := range routes {
		u := fmt.Sprintf("%v%v", prefix, url)
		r.m.HandleFunc(u, fn)
	}
	return nil
}

func (routes *Routes) PlayerPlay(w http.ResponseWriter, r *http.Request) {
	http.Redirect(routes.withoutFrontendCache(w), r, r.Referer(), http.StatusPermanentRedirect)
}
