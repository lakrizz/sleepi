package alarm

import (
	"log"
	"time"

	"github.com/lakrizz/sleepi/pkg/player"
	"github.com/lakrizz/sleepi/pkg/playlist"
)

// the watcher takes an alarm and subscribes the next occurance with the manager
// it actually is the cronjob service behind the a single next occurance of an alarm
// and its 'execution' in the player context

type alarmWatcher struct {
	manager      *AlarmManager
	target       *Alarm
	target_timer *time.Timer
	player       *player.Player
	playlists    *playlist.PlaylistManager
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
			go aw.triggerAlarm()

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
			_, err := aw.target.TimeTillNextWake()
			if err != nil {
				panic(err)
			}
			// log.Println(fmt.Sprintf("next alarm in %s", dur.String()))
			time.Sleep(250 * time.Millisecond)
		}
	}
}

func (aw *alarmWatcher) triggerAlarm() {
	log.Println("triggering alarm!")
	if !aw.player.IsPlaying {
		pl, err := aw.playlists.GetPlaylist(aw.target.Playlist)
		if err != nil {
			log.Fatal(err)
		}
		if err := aw.player.QueueSongs(pl.Songs, false); err != nil {
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
