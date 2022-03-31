package main

import (
	"log"
	"time"

	"krizz.org/sleepi/internal/manager"
	"krizz.org/sleepi/pkg/alarm"
	"krizz.org/sleepi/pkg/library"
	"krizz.org/sleepi/pkg/playlist"
)

func main() {

	pl, _ := playlist.NewPlaylist("wow")
	lib, _ := library.GetLibrary()

	for _, v := range lib.GetAllFiles() {
		pl.Add(v)
	}

	alarm.StartMPDPlaylist(pl, false)
	now := time.Now()
	al, _ := alarm.CreateAlarm(func() { alarm.StartMPDPlaylist(pl, true) }, []time.Weekday{time.Sunday}, 0, 15)
	al2, _ := alarm.CreateAlarm(echo, []time.Weekday{time.Saturday}, 22, 56)
	al3, _ := alarm.CreateAlarm(func() { alarm.StartMPDPlaylist(pl, true) }, []time.Weekday{time.Sunday}, now.Hour(), now.Minute()+1)
	am, _ := manager.GetAlarmManager()
	am.AddAlarm(al)
	am.AddAlarm(al2)
	am.AddAlarm(al3)

	time.Sleep(2 * time.Second)
	am.AddAlarm(al3)

	for {
	}
}

func echo() {
	log.Println("ALARM!")
}
