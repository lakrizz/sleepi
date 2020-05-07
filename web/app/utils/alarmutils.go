package utils

import (
	"errors"
	"strconv"
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
		return_map = append(return_map, &WakeDay{i, days[i], a != nil && dayInSlice(time.Weekday(i), a.Days)})
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

func StringToWeekdaySlice(str []string) ([]time.Weekday, error) {
	days := make([]time.Weekday, 0)

	for _, v := range str {
		// we want to know whati's inside, eh? :D
		d, err := strconv.Atoi(v)
		if err != nil {
			return nil, err
		}
		days = append(days, time.Weekday(d))
	}

	if len(days) == 0 {
		return nil, errors.New("must at least choose one day, eh?")
	}
	return days, nil
}
