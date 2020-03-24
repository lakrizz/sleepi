package alarm

import (
	"fmt"
	"log"
	"time"

	"../player"
)

// the watcher takes an alarm and subscribes the next occurance with the manager
// it actually is the cronjob service behind the a single next occurance of an alarm
// and its 'execution' in the player context

type alarmWatcher struct {
	manager      *AlarmManager
	target       *Alarm
	target_timer *time.Timer
	player       *player.Player
	refreshChan  chan bool
}

func createWatcher(manager *AlarmManager) (*alarmWatcher, error) {
	aw := &alarmWatcher{manager: manager}
	aw.refreshChan = make(chan bool, 1)
	player, err := player.GetPlayer()
	if err != nil {
		return nil, err
	}
	aw.player = player
	aw.refresh()

	return aw, nil
}

func (aw *alarmWatcher) refresh() {
	target, err := aw.manager.GetNextAlarm()
	if err != nil {
		log.Fatal(err)
	} else {
		// this is also where we want to call the cronjob when to call the player ;)
		// but consider that this is not only called when an alarm is hit!
		// so we need to check that, and that's important!
		aw.target = target
		if aw.target_timer != nil {
			aw.target_timer.Stop()
		}
		dur, err := aw.target.TimeTillNextWake()
		if err != nil {
			log.Fatal(err)
		} else {
			aw.target_timer = time.NewTimer(dur)
			aw.refreshChan <- true
		}
	}
}

func (aw *alarmWatcher) run() {
	for {
		select {
		case <-aw.target_timer.C:
			go aw.TriggerAlarm()
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

func (aw *alarmWatcher) TriggerAlarm() {
	log.Println("triggering alarm!")
	if !aw.player.IsPlaying {
		if err := aw.player.LoadPlaylist(aw.target.Playlist, true); err != nil {
			log.Fatal(err)
		} else {
			log.Println("loaded playlist, starting to play now")
			err = aw.player.Play()
			log.Println("i played?!")
			if err != nil {
				panic(err)
			}
		}
	}
}
