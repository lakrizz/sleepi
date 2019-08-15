package youtube

import (
	"net/http"

	modules ".."
	"../../modules/youtube"
	"github.com/k0kubun/pp"
	"github.com/unrolled/render"
)

type youTubeModule struct {
	routes []modules.Route
	r      *render.Render
	yt     *youtube.YouTube
}

func YouTubeModule(ren *render.Render, yt *youtube.YouTube) *youTubeModule {
	y := &youTubeModule{}
	y.r = ren
	y.yt = yt
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
	y.routes = append(y.routes,
		[]modules.Route{
			modules.Route{"/", y.MainRoute},
			modules.Route{"", y.MainRoute},
			modules.Route{"/search", y.SearchRoute},
			modules.Route{"/search/", y.SearchRoute},
		}...)
}

func (y *youTubeModule) MainRoute(r http.ResponseWriter, req *http.Request) {
	// this is the main screen, we will add a search button here and also use some websocket magic
	// for communication, there will be no other route here, we need a library controller then.
	y.r.HTML(r, http.StatusOK, "youtube/main", nil)
}

func (y *youTubeModule) SearchRoute(r http.ResponseWriter, req *http.Request) {
	err := req.ParseForm()
	if err != nil {
		y.r.Text(r, http.StatusBadRequest, err.Error())
	}

	searchterm := req.FormValue("term")
	result, err := y.yt.Search.SearchVideos(searchterm)

	if err != nil {
		y.r.Text(r, http.StatusBadRequest, err.Error())
	}
	pp.Println(result)
	y.r.HTML(r, http.StatusOK, "youtube/result", result)
}
