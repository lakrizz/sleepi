package main

import (
	"log"
	"time"

	"github.com/k0kubun/pp"
	"krizz.org/sleepi/internal/manager"
	"krizz.org/sleepi/pkg/alarm"
)

func main() {
	al, _ := alarm.CreateAlarm(echo, []time.Weekday{time.Sunday}, 7, 30)
	al2, _ := alarm.CreateAlarm(echo, []time.Weekday{time.Sunday, time.Saturday}, 8, 30)
	am, err := manager.GetAlarmManager([]*alarm.Alarm{al, al2})
	next, err := am.GetClosestAlarm()
	pp.Println(am, err, next)
}

func echo() {
	log.Println("ALARM!")
}
