package youtube

import (
	"net/http"

	modules "../"
	"github.com/unrolled/render"
)

type youTubeModule struct {
	routes []modules.Route
	r      *render.Render
}

func YouTubeModule(ren *render.Render) *youTubeModule {
	y := &youTubeModule{}
	y.r = ren
	y.initRoutes()
	return y
}

func (y *youTubeModule) GetName() string {
	return "youtube"
}

func (y *youTubeModule) GetRoutes() []modules.Route {
	return y.routes
}

func (y *youTubeModule) initRoutes() {
	y.routes = append(y.routes, modules.Route{"/", y.MainRoute})
	y.routes = append(y.routes, modules.Route{"", y.MainRoute})
}

func (y *youTubeModule) MainRoute(r http.ResponseWriter, req *http.Request) {
	// this is the main screen, we will add a search button here and also use some websocket magic
	// for communication, there will be no other route here, we need a library controller then.
	y.r.HTML(r, http.StatusOK, "youtube/main", nil)
}
