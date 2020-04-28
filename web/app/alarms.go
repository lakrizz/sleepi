package app

import (
	"net/http"
)

func AlarmsHome(w http.ResponseWriter, r *http.Request) {
	alarms, err := api.GetAlarms()
	if err != nil {
		// show error page
	} else {
		ren.HTML(w, http.StatusOK, "alarms/main", alarms)
	}
}
