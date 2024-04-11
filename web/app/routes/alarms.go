package routes

import (
	"errors"
	"fmt"
	"log"
	"net/http"
	"strconv"
	"time"

	"github.com/google/uuid"
	"github.com/gorilla/mux"

	"github.com/lakrizz/sleepi/pkg/alarm"
	"github.com/lakrizz/sleepi/pkg/effects"
	"github.com/lakrizz/sleepi/pkg/util"
)

func (r *Routes) AddAlarmRoutes() error {
	if r == nil {
		return errors.New("routes is null")
	}
	prefix := "/alarms"
	routes := map[string]func(http.ResponseWriter, *http.Request){
		"/":                 r.AlarmIndex,
		"/new":              r.AlarmNew,
		"/create":           r.AlarmCreate,
		"/{id}/activate/":   r.AlarmActivate,
		"/{id}/deactivate/": r.AlarmDeactivate,
	}
	for url, fn := range routes {
		u := fmt.Sprintf("%v%v", prefix, url)
		r.m.HandleFunc(u, fn)
	}
	return nil
}

func (routes *Routes) AlarmIndex(w http.ResponseWriter, r *http.Request) {
	alarms, err := routes.rt.Alarms.GetAllAlarms()
	if err != nil {
		return
	}

	vars := map[string]any{
		"Alarms": alarms,
		"Days":   util.Weekdays(),
	}

	routes.ren.HTML(w, http.StatusOK, "alarms/main", vars)
}

func (routes *Routes) AlarmNew(w http.ResponseWriter, r *http.Request) {
	playlists, err := routes.rt.Playlists.GetAllPlaylists()
	if err != nil {
		return
	}

	vars := map[string]any{
		"Days":      util.Weekdays(),
		"Playlists": playlists,
	}
	routes.ren.HTML(w, http.StatusOK, "alarms/new", vars)
}

func (routes *Routes) AlarmActivate(w http.ResponseWriter, r *http.Request) {
	id, err := uuid.Parse(mux.Vars(r)["id"])
	if err != nil {
		log.Println(err)
		routes.ren.Data(w, http.StatusBadRequest, []byte(err.Error()))
		return
	}

	a, err := routes.rt.Alarms.GetAlarm(id)
	if err != nil {
		log.Println(err)
		routes.ren.Data(w, http.StatusBadRequest, []byte(err.Error()))
		return
	}

	a.Enabled = true
	err = routes.rt.Alarms.UpdateTimings()
	if err != nil {
		routes.ren.Data(w, http.StatusBadRequest, []byte(err.Error()))
		return
	}

	err = routes.rt.Alarms.Save()
	if err != nil {
		routes.ren.Data(w, http.StatusBadRequest, []byte(err.Error()))
		return
	}
	http.Redirect(routes.withoutFrontendCache(w), r, "/alarms/", http.StatusPermanentRedirect)
}

func (routes *Routes) AlarmDeactivate(w http.ResponseWriter, r *http.Request) {
	id, err := uuid.Parse(mux.Vars(r)["id"])
	if err != nil {
		routes.ren.Data(w, http.StatusBadRequest, []byte(err.Error()))
		return
	}

	a, err := routes.rt.Alarms.GetAlarm(id)
	if err != nil {
		routes.ren.Data(w, http.StatusBadRequest, []byte(err.Error()))
		return
	}

	a.Enabled = false
	err = routes.rt.Alarms.UpdateTimings()
	if err != nil {
		routes.ren.Data(w, http.StatusBadRequest, []byte(err.Error()))
		return
	}

	err = routes.rt.Alarms.Save()
	if err != nil {
		routes.ren.Data(w, http.StatusBadRequest, []byte(err.Error()))
		return
	}

	http.Redirect(routes.withoutFrontendCache(w), r, "/alarms/", http.StatusPermanentRedirect)
}

func (routes *Routes) AlarmCreate(w http.ResponseWriter, r *http.Request) {
	if r.Method != "POST" {
		routes.ren.Data(w, http.StatusMethodNotAllowed, []byte("this is not a post request"))
		return
	}

	err := r.ParseForm()
	if err != nil {
		routes.ren.Data(w, http.StatusMethodNotAllowed, []byte(err.Error()))
		return
	}

	id, err := uuid.Parse(r.PostFormValue("alarm-playlist"))
	if err != nil {
		routes.ren.Data(w, http.StatusBadRequest, []byte(err.Error()))
		return
	}

	days := make([]time.Weekday, 0)
	for _, v := range []string(r.PostForm["days"]) {
		vi, err := strconv.Atoi(v)
		if err != nil {
			routes.ren.Data(w, http.StatusBadRequest, []byte(err.Error()))
			return
		}
		days = append(days, time.Weekday(vi))
	}

	minute, err := strconv.Atoi(r.PostFormValue("alarm-minute"))
	if err != nil {
		routes.ren.Data(w, http.StatusBadRequest, []byte(err.Error()))
		return
	}
	hour, err := strconv.Atoi(r.PostFormValue("alarm-hour"))
	if err != nil {
		routes.ren.Data(w, http.StatusBadRequest, []byte(err.Error()))
		return
	}

	name := r.PostFormValue("alarm-name")

	a, err := alarm.CreateAlarm(&id, name, days, hour, minute)
	if err != nil {
		routes.ren.Data(w, http.StatusBadRequest, []byte(err.Error()))
		return
	}

	// check for volume warmup
	warmup_enabled := r.PostFormValue("cb_volumewarmup")
	if warmup_enabled != "" && warmup_enabled == "on" {
		warmup := &effects.VolumeWarmup{}
		wu_sv, err := strconv.Atoi(r.PostFormValue("warmup-start-volume"))
		if err != nil {
			routes.ren.Data(w, http.StatusBadRequest, []byte(err.Error()))
			return
		}
		warmup.StartVolume = wu_sv

		wu_ev, err := strconv.Atoi(r.PostFormValue("warmup-end-volume"))
		if err != nil {
			routes.ren.Data(w, http.StatusBadRequest, []byte(err.Error()))
			return
		}
		warmup.EndVolume = wu_ev

		wu_dur, err := time.ParseDuration(r.PostFormValue("warmup-duration"))
		if err != nil {
			routes.ren.Data(w, http.StatusBadRequest, []byte(err.Error()))
			return
		}
		warmup.Duration = &wu_dur
	}

	err = routes.rt.Alarms.AddAlarm(a)
	if err != nil {
		routes.ren.Data(w, http.StatusBadRequest, []byte(err.Error()))
		return
	}

	http.Redirect(routes.withoutFrontendCache(w), r, "/alarms/", http.StatusPermanentRedirect)
}
