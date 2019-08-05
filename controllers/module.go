package modules

import "net/http"

type Route struct {
	Path string
	Func func(http.ResponseWriter, *http.Request)
}

type Module interface {
	GetName() string
	GetRoutes() []Route
}
