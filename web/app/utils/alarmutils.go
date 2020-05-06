package utils

import (
	"time"

	"github.com/lakrizz/sleepi/pkg/alarm"
)

type WakeDay struct {
	Id       int
	Readable string
	Enabled  bool
}

func CreateWakeDayMap(a *alarm.Alarm) []*WakeDay {
	days := []string{"Sunday", "Monday", "Tuesday", "Wednesday", "Thursday", "Friday", "Saturday"}
	return_map := make([]*WakeDay, 0)
	for i := 0; i < 7; i++ {
		return_map = append(return_map, &WakeDay{i, days[i], dayInSlice(time.Weekday(i), a.Days)})
	}
	return return_map
}

func dayInSlice(needle time.Weekday, haystack []time.Weekday) bool {
	for _, v := range haystack {
		if v == needle {
			return true
		}
	}
	return false
}
