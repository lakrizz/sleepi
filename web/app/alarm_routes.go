package app

import (
	"errors"
	"fmt"
	"net/http"

	"github.com/k0kubun/pp"
)

func (r *Routes) addAlarmRoutes() error {
	if r == nil {
		return errors.New("routes is null")
	}
	prefix := "/alarms"
	routes := map[string]func(http.ResponseWriter, *http.Request){
		"/": r.AlarmIndex,
	}
	for url, fn := range routes {
		u := fmt.Sprintf("%v%v", prefix, url)
		r.m.HandleFunc(u, fn)
	}
	return nil
}

func (routes *Routes) AlarmIndex(w http.ResponseWriter, r *http.Request) {
	vars := make(map[string]interface{})
	vars["alarms"] = routes.api.alarms.Alarms
	pp.Println(vars)
	routes.ren.HTML(w, http.StatusOK, "alarms/index", vars)
}
