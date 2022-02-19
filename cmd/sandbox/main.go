package main

import (
	"log"
	"time"

	"krizz.org/sleepi/internal/manager"
	"krizz.org/sleepi/pkg/alarm"
)

func main() {
	al, _ := alarm.CreateAlarm(echo, []time.Weekday{time.Sunday}, 13, 30)
	al2, _ := alarm.CreateAlarm(echo, []time.Weekday{time.Saturday}, 21, 30)
	al3, _ := alarm.CreateAlarm(echo, []time.Weekday{time.Sunday}, 8, 30)
	am, _ := manager.GetAlarmManager([]*alarm.Alarm{al, al2})

	time.Sleep(2 * time.Second)
	am.AddAlarm(al3)
	time.Sleep(2 * time.Second)

}

func echo() {
	log.Println("ALARM!")
}
