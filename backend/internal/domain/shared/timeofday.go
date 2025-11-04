package shared

import "time"

// TimeOfDay represents a local clock time without a date.
type TimeOfDay struct {
	hour   int
	minute int
}

func ParseTimeOfDay(s string) (TimeOfDay, error) {
	t, err := time.Parse("15:04", s)
	if err != nil {
		return TimeOfDay{}, err
	}
	return TimeOfDay{hour: t.Hour(), minute: t.Minute()}, nil
}

func (t TimeOfDay) String() string {
	return time.Date(0, 0, 0, t.hour, t.minute, 0, 0, time.UTC).Format("15:04")
}

func (t TimeOfDay) Matches(now time.Time) bool {
	return now.Hour() == t.hour && now.Minute() == t.minute
}

func (t *TimeOfDay) Hour() int {
	return t.hour
}
func (t *TimeOfDay) Minute() int {
	return t.minute
}
