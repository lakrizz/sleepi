package playlist

import (
	"log"
	"net/http"

	modules ".."

	"github.com/fhs/gompd/mpd"
	"github.com/unrolled/render"
)

type playlistModule struct {
	routes []modules.Route
	mc     *mpd.Client
	r      *render.Render
}

func PlaylistModule(ren *render.Render) *playlistModule {
	log.Println("ohai i am playlist module")
	mod := &playlistModule{}

	c, err := mpd.Dial("tcp", "127.0.0.1:6600")
	if err != nil {
		log.Println("ERROR: ", err.Error())
		return nil
	}

	mod.initRoutes()
	mod.mc = c
	mod.r = ren
	return mod

}

func (p *playlistModule) GetName() string {
	return "playlist"
}

func (p *playlistModule) GetRoutes() []modules.Route {
	return p.routes
}

func (p *playlistModule) initRoutes() {
	p.routes = append(p.routes,
		[]modules.Route{
			modules.Route{"/", p.MainRoute},
		}...,
	)
}

type playlst struct {
	Name      string
	Songs     []string
	SongCount int
}

func (p *playlistModule) MainRoute(r http.ResponseWriter, req *http.Request) {
	dm := make([]*playlst, 0) // a slice of playlists with its name as key and a slice of songs as values, kchch
	if p.mc != nil {
		pl, err := p.mc.ListPlaylists()
		if err != nil {
			log.Println(err.Error())
			return
		}
		for _, cpl := range pl[1:] {
			curpl := make([]string, 0)
			songs, err := p.mc.PlaylistContents(cpl["playlist"])
			if err != nil {
				log.Println(err.Error())
				return
			}

			for _, v := range songs {
				curpl = append(curpl, v["Artist"]+" - "+v["Title"])
			}

			dm = append(dm, &playlst{cpl["playlist"], curpl, len(curpl)})
		}

		p.r.HTML(r, http.StatusOK, "playlist/main", dm)
	}
}
