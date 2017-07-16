package testmodule

import (
	"net/http"

	"../"
)

type testModule struct {
	routes []modules.Route
}

func TestModule() *testModule {
	t := &testModule{}
	t.initRoutes()
	return t
}

func (t *testModule) GetName() string {
	return "TestModule"
}

func (t *testModule) GetRoutes() []modules.Route {
	return t.routes
}

func (t *testModule) initRoutes() {
	t.routes = append(t.routes, modules.Route{"/", t.MainRoute})
}

func (t *testModule) MainRoute(r http.ResponseWriter, req *http.Request) {
	r.Write([]byte("Hola Amigo!"))
}
