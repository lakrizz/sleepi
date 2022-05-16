package app

import (
	"errors"
	"fmt"
	"net/http"

	"github.com/k0kubun/pp"
	"krizz.org/sleepi/pkg/util"
)

func (r *Routes) addAlarmRoutes() error {
	if r == nil {
		return errors.New("routes is null")
	}
	prefix := "/alarms"
	routes := map[string]func(http.ResponseWriter, *http.Request){
		"/":       r.AlarmIndex,
		"/new":    r.AlarmNew,
		"/create": r.AlarmCreate,
	}
	for url, fn := range routes {
		u := fmt.Sprintf("%v%v", prefix, url)
		r.m.HandleFunc(u, fn)
	}
	return nil
}

func (routes *Routes) AlarmIndex(w http.ResponseWriter, r *http.Request) {
	vars := make(map[string]interface{})
	vars["Alarms"] = routes.api.alarms.Alarms
	routes.ren.HTML(w, http.StatusOK, "alarms/main", vars)
}

func (routes *Routes) AlarmNew(w http.ResponseWriter, r *http.Request) {
	vars := make(map[string]interface{})
	vars["Days"] = util.Weekdays()
	vars["Playlists"] = routes.api.playlists.Playlists
	routes.ren.HTML(w, http.StatusOK, "alarms/new", vars)
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

	pp.Println(r.PostForm)

	http.Redirect(routes.withoutFrontendCache(w), r, "/playlists/", http.StatusPermanentRedirect)
}
