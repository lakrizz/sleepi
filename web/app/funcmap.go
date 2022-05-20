package app

import (
	"fmt"
	"html/template"
	"time"

	"krizz.org/sleepi/pkg/alarm"
	"krizz.org/sleepi/pkg/library"
)

func (a *App) GetFuncMap() template.FuncMap {
	return template.FuncMap(map[string]interface{}{
		"TopFiles": func(slice []*library.File, num int) []*library.File {
			if len(slice) > num {
				return slice[:num]
			}
			return slice
		},
		"Cut": func(s string, i int) string {
			if len(s) <= i {
				return s
			}
			return s[:i]
		},
		"IsActiveDay": func(day time.Weekday, alarm *alarm.Alarm) bool {
			return alarm.IsActiveDay(day)
		},
		"formatDurationWithDays": func(dur time.Duration) string {
			minutes := int64(dur/time.Minute) % int64(60)
			days := int64(dur/(24*time.Hour)) % int64(60*24)
			hours := int64(dur/time.Hour) % int64(24)
			return fmt.Sprintf("%v %s, %v %s, %v %s", days, "days", hours, "hours", minutes, "minutes")
		},
	})
}
