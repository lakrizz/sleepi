package app

import (
	"fmt"
	"net/http"
)

func addAlarmRoutes() {
	prefix := "/alarms"
	m.HandleFunc(fmt.Sprintf("%v", prefix), AlarmIndex)
}

func AlarmIndex(w http.ResponseWriter, r *http.Request) {
	vars := make(map[string]interface{})
	vars["alarms"] = alarmManager.Alarms
	ren.HTML(w, http.StatusOK, "alarms/index", vars)
}
