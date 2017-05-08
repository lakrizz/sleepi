package testmodule

import "net/http"

type TestModule struct {
	routes []*Route
}

func TestModule() *TestModule {
	t := &TestModule{}
	t.initRoutes()
	return t
}

func (t *TestModule) GetName() string {
	return "testmodule"
}

func (t *TestModule) GetRoutes() []*Route {
	return t.routes
}

func (t *TestModule) initRoutes() {
	t.routes = append(t.routes, &Route{"/", t.MainRoute})
}

func (t *TestModule) MainRoute(r http.ResponseWriter, req *http.Request) {
	r.Write([]byte("Hola Amigo!"))
}
