package alarm

import (
	"log"
	"net/http"
	"strconv"
	"strings"
	"time"

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
			modules.Route{"put", a.PutAlarmRoute},
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

func (a *alarmModule) PutAlarmRoute(r http.ResponseWriter, req *http.Request) {
	err := req.ParseForm()
	if err != nil {
		a.r.JSON(r, http.StatusBadRequest, "Kapotttt!")
		return
	}
	// now we can use the form values
	log.Println("postform:", req.PostForm)

	days := make([]time.Weekday, 0)
	pf := strings.Split(req.PostFormValue("weekday"), " ")
	log.Println(pf)
	for _, d := range pf {
		log.Println(d)
		cv, err := strconv.Atoi(d)
		if err != nil {
			log.Println(err.Error())
		} else {
			days = append(days, time.Weekday(cv))
		}
	}
	log.Println(days)

	plname := req.PostFormValue("playlist")
	// validity check if pl exists

	alarm_hh, err := strconv.Atoi(req.PostFormValue("time_hh"))
	alarm_mm, err := strconv.Atoi(req.PostFormValue("time_mm"))

	err = a.AddAlarm(CreateAlarm(a.mc, plname, []time.Weekday(days), alarm_hh, alarm_mm, 0, 100, 15))
	if err != nil {
		log.Println(err.Error())
	}
	err = a.Save()
	if err != nil {
		log.Println(err.Error())
	}

}
