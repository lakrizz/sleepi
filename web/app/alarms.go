package app

import (
	"fmt"
	"log"
	"net/http"
	"strconv"
	"time"

	"github.com/google/uuid"
	"github.com/gorilla/mux"
	"github.com/lakrizz/sleepi/pkg/alarm"
	"github.com/lakrizz/sleepi/pkg/playlist"
)

func AlarmsHome(w http.ResponseWriter, r *http.Request) {
	alarms, err := api.GetAlarms()
	if err != nil {
		// show error page
	} else {
		ren.HTML(w, http.StatusOK, "alarms/main", alarms)
	}
}

func AlarmsView(w http.ResponseWriter, r *http.Request) {
	vars := mux.Vars(r)

	a, err := api.GetAlarm(vars["id"])
	if err != nil {
		// show error page
	} else {
		// get playlist for alarm
		p, err := api.Playlists.GetPlaylist(a.Playlist)
		if err != nil {
			// show error page
		}

		all_playlists := api.Playlists.Playlists

		ret := struct {
			Alarm        *alarm.Alarm
			Playlist     *playlist.Playlist
			AllPlaylists []*playlist.Playlist
		}{
			a,
			p,
			all_playlists,
		}
		ren.HTML(w, http.StatusOK, "alarms/view", ret)
	}
}

func AlarmsDisable(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Cache-Control", "no-cache, no-store, must-revalidate") // HTTP 1.1.
	w.Header().Set("Pragma", "no-cache")                                   // HTTP 1.0.
	w.Header().Set("Expires", "0")                                         // Proxies.
	vars := mux.Vars(r)
	err := api.DisableAlarm(vars["id"])
	if err != nil {
		// error
		return
	}

	http.Redirect(w, r, fmt.Sprintf("/alarms/%s", vars["id"]), http.StatusPermanentRedirect)
}
func AlarmsEnable(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Cache-Control", "no-cache, no-store, must-revalidate") // HTTP 1.1.
	w.Header().Set("Pragma", "no-cache")                                   // HTTP 1.0.
	w.Header().Set("Expires", "0")                                         // Proxies.
	vars := mux.Vars(r)
	err := api.EnableAlarm(vars["id"])
	if err != nil {
		// error
		return
	}

	http.Redirect(w, r, fmt.Sprintf("/alarms/%s", vars["id"]), http.StatusPermanentRedirect)
}

func AlarmsEdit(w http.ResponseWriter, r *http.Request) {
	vars := mux.Vars(r)
	alarm, err := api.GetAlarm(vars["id"])
	if err != nil {
		// error
		log.Println(err.Error())
		return
	}

	err = r.ParseForm()
	if err != nil {
		// error
		log.Println(err.Error())
		return
	}

	alarm_wakehour, err := strconv.Atoi(r.FormValue("alarm_wakehour"))
	if err != nil {
		// error
		log.Println(err.Error())
		return
	}

	alarm_wakeminute, err := strconv.Atoi(r.FormValue("alarm_wakeminute"))
	if err != nil {
		// error
		log.Println(err.Error())
		return
	}

	alarm_playlist, err := uuid.Parse(r.FormValue("alarm_playlist"))
	if err != nil {
		// error
		log.Println(err.Error())
		return
	}

	_, err = time.ParseDuration(r.FormValue("alarm_waketime"))
	if err != nil {
		// error
		log.Println(err.Error())
		return
	}

	fmt.Println(r.Form)
	alarm.Name = r.FormValue("alarm_name")
	alarm.WakeHour = alarm_wakehour
	alarm.WakeMinute = alarm_wakeminute
	alarm.Playlist = alarm_playlist
	alarm.ShufflePlaylist = r.FormValue("shuffle_playlist") == "on"
	alarm.WakeupTime = r.FormValue("alarm_waketime")
	err = api.Alarms.SaveAlarms()
	if err != nil {
		// error
		log.Println(err.Error())
		return
	}

	http.Redirect(w, r, fmt.Sprintf("/alarms/%s", vars["id"]), http.StatusPermanentRedirect)
}
