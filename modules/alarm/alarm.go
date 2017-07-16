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
	t.initRoutes()
	t.mc = c
	t.r = ren
	t.alarms = t.createAlarms(10)
	return t
}

func (a *alarmModule) createAlarms(amount int) []*alarm {
	// alarms := make([]*alarm, 0)
	// for i := 0; i <= amount; i++ {
	// }
	return nil
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
	// overview, huh?
	a.r.HTML(r, http.StatusOK, "alarm/main", a.alarms)
}

func (a *alarmModule) TestRoute(r http.ResponseWriter, req *http.Request) {
	_ = CreateAlarm(a.mc, "wakeywakey (by 1121749173)")

	// pl, err := a.mc.ListPlaylists()
	// if err != nil {
	// 	return
	// }

	// for _, i := range pl[1:] { // the first playlist is always ".plu8", so skip it :D
	// 	r.Write([]byte(i["playlist"] + "\n"))
	// }
}

func (a *alarmModule) AddAlarmRoute(r http.ResponseWriter, req *http.Request) {
	pl, err := a.mc.ListPlaylists()
	if err != nil {
		a.r.HTML(r, http.StatusOK, "alarm/add", nil)
		return
	}
	a.r.HTML(r, http.StatusOK, "alarm/add", pl[1:])
}
