package alarm

import (
	"fmt"
	"log"
	"time"
)

// the watcher takes an alarm and subscribes the next occurance with the manager
// it actually is the cronjob service behind the a single next occurance of an alarm
// and its 'execution' in the player context

type alarmWatcher struct {
	manager     *AlarmManager
	target      *Alarm
	refreshChan chan bool
}

func createWatcher(manager *AlarmManager) (*alarmWatcher, error) {
	aw := &alarmWatcher{manager: manager}
	aw.refreshChan = make(chan bool, 1)
	aw.refresh()

	return aw, nil
}

func (aw *alarmWatcher) refresh() {
	target, err := aw.manager.GetNextAlarm()
	if err != nil {
		log.Fatal(err)
	} else {
		aw.target = target
		aw.refreshChan <- true
	}
}

func (aw *alarmWatcher) run() {
	for {
		select {
		default:
			dur, err := aw.target.TimeTillNextWake()
			if err != nil {
				panic(err)
			}
			log.Println(fmt.Sprintf("next alarm in %s", dur.String()))
			time.Sleep(50 * time.Millisecond)
		}
	}
}
