package alarm

import (
	"log"
	"net/http"

	"../"
	"github.com/fhs/gompd/mpd"
	"github.com/unrolled/render"
)

type alarmModule struct {
	routes []modules.Route
	mc     *mpd.Client
	alarms []*alarm
	r      *render.Render
}

func AlarmModule(ren *render.Render) *alarmModule {
	log.Println("ohai")
	c, err := mpd.Dial("tcp", "127.0.0.1:6600")
	if err != nil {
		log.Println(err.Error())
		return nil
	}
	c.ListPlaylists()

	t := &alarmModule{}
	err = t.Load()
	if err != nil {
		log.Println(err.Error())
	}
	t.initRoutes()
	t.mc = c
	t.r = ren
	return t
}

func (a *alarmModule) GetName() string {
	return "alarm"
}

func (a *alarmModule) GetRoutes() []modules.Route {
	return a.routes
}

func (a *alarmModule) initRoutes() {
	a.routes = append(a.routes,
		[]modules.Route{
			modules.Route{"test", a.TestRoute},
			modules.Route{"", a.MainRoute},
			modules.Route{"add", a.AddAlarmRoute},
		}...,
	)
}

func (a *alarmModule) MainRoute(r http.ResponseWriter, req *http.Request) {
	// a.AddAlarm(CreateAlarm(
	// 	a.mc,
	// 	"wakeywakey (by 1121749173)",
	// 	[]time.Weekday{time.Monday, time.Tuesday, time.Wednesday},
	// 	6,  // *06*:30
	// 	30, // 06:_30_
	// 	0,
	// 	100,
	// 	30, // minutes
	// ))
	a.r.HTML(r, http.StatusOK, "alarm/main", a.alarms)
}

func (a *alarmModule) TestRoute(r http.ResponseWriter, req *http.Request) {
}

func (a *alarmModule) AddAlarmRoute(r http.ResponseWriter, req *http.Request) {
	pl, err := a.mc.ListPlaylists()
	if err != nil {
		log.Println(err.Error())
		a.r.HTML(r, http.StatusOK, "alarm/add", nil)
		return
	}
	a.r.HTML(r, http.StatusOK, "alarm/add", pl[1:])
}
