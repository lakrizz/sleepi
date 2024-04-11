package pkg

import "time"

type Alarm interface {
	TimeUntilNextTrigger() time.Duration
	IsActiveForDay(time.Weekday) bool
	Trigger() error
}

type Effect interface {
}
