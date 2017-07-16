package alarm

import (
	"log"
	"time"

	"github.com/fhs/gompd/mpd"
	"github.com/jasonlvhit/gocron"
)

type alarm struct {
	playlistName           string
	duration               time.Duration // the amount of time the bottom function runs
	startVolume, endVolume int           // mpd volume information, e.g. 30 and 100
	mc                     *mpd.Client
	active                 bool
	wakeDay                time.Weekday
	// wakeTime               time.Time // v2
	// nextAlarm              time.Time // v2
}

func CreateAlarm(client *mpd.Client, plName string) *alarm {
	al := &alarm{mc: client, playlistName: plName, startVolume: 0, endVolume: 100, duration: (15 * time.Minute)}
	al.setAlarms()
	return al
}

func (a *alarm) setAlarms() {
	gocron.Every(1).Monday().At("06:45").Do(a.wake)
	gocron.Every(1).Tuesday().At("06:45").Do(a.wake)
	gocron.Every(1).Wednesday().At("06:45").Do(a.wake)
	gocron.Every(1).Thursday().At("06:45").Do(a.wake)
	gocron.Every(1).Friday().At("06:45").Do(a.wake)
	gocron.RunAll()
}

func inSlice(n string, h []string) bool {
	for _, v := range h {
		if v == n {
			return true
		}
	}
	return false
}

func (a *alarm) wake() {
	err := a.mc.Clear()
	if err != nil {
		log.Println(err.Error())
		return
	}
	a.mc.PlaylistLoad(a.playlistName, -1, -1) // loads the playlist that we want to wake up to
	a.mc.SetVolume(a.startVolume)             // start at stgartvolume
	volStep := a.duration.Nanoseconds() / int64(a.endVolume-a.startVolume)
	a.mc.Play(0)
	a.mc.Random(true)
	a.mc.Shuffle(-1, -1) // we don't want to have the same seed all the time, this might be an issue with mopidy
	for i := a.startVolume; i < a.endVolume; i++ {
		a.mc.SetVolume(i)
		time.Sleep(time.Duration(volStep) * time.Nanosecond)
	}
	// and now we should be awake and beautiful :D
}
