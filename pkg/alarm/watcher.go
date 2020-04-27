package alarm

import (
	"fmt"
	"log"
	"time"

	"github.com/lakrizz/sleepi/pkg/player"
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
	target, err := aw.manager.GetNextAlarm()
	if err != nil {
		return nil, err
	}
	aw.target = target
	dur, err := aw.target.TimeTillNextWake()
	if err != nil {
		return nil, err
	}
	aw.target_timer = time.NewTimer(dur)

	return aw, nil
}

func (aw *alarmWatcher) run() {
	for {
		select {
		case <-aw.target_timer.C:
			log.Println("triggering alarm")
			go aw.triggerAlarm(aw.target.Playlist)

			target, err := aw.manager.GetNextAlarm()
			if err != nil {
				log.Fatal(err)
			}
			aw.target = target
			log.Println("next target is", target)
			dur, err := aw.target.TimeTillNextWake()
			if err != nil {
				log.Fatal(err)
			}
			log.Println("resetting timer")
			aw.target_timer.Reset(dur)
		default:
			dur, err := aw.target.TimeTillNextWake()
			if err != nil {
				panic(err)
			}
			log.Println(fmt.Sprintf("next alarm in %s", dur.String()))
			time.Sleep(25 * time.Millisecond)
		}
	}
}

func (aw *alarmWatcher) triggerAlarm(playlist string) {
	log.Println("triggering alarm!")
	if !aw.player.IsPlaying {
		if err := aw.player.LoadPlaylist(playlist, false); err != nil {
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
