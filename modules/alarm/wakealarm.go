package alarm

import (
	"strings"
	"time"

	"github.com/fhs/gompd/mpd"
	"github.com/hako/durafmt"
	"labix.org/v2/mgo/bson"
)

type alarm struct {
	ID                     bson.ObjectId
	PlaylistName           string
	Duration               time.Duration // the amount of time the bottom function runs
	StartVolume, EndVolume int           // mpd volume information, e.g. 30 and 100
	mc                     *mpd.Client
	Active                 bool
	WakeDays               []time.Weekday
	WakeHour               int // v2
	WakeMinute             int
	nextrun                time.Time
}

func CreateAlarm(client *mpd.Client, plName string, days []time.Weekday, hour, minute, startvolume, endvolume, duration_minutes int) *alarm {
	id := bson.NewObjectId()
	al := &alarm{
		ID:           id,
		mc:           client,
		PlaylistName: plName,
		StartVolume:  startvolume,
		EndVolume:    endvolume,
		Duration:     time.Duration(duration_minutes) * time.Minute,
		WakeDays:     days,
		WakeHour:     hour,
		WakeMinute:   minute,
	}

	return al
}

func (a *alarm) calcNextRun() {
	now := time.Now()

	// today, but later
	if inWeekdays(a.WakeDays, now.Weekday()) && now.Hour() <= a.WakeHour && now.Minute() < a.WakeMinute {
		a.nextrun = time.Date(now.Year(), now.Month(), now.Day(), a.WakeHour, a.WakeMinute, 0, 0, now.Location())
		return
	}

	// from tomorrow on
	for i := 0; i < 7; i++ {
		now = now.Add(time.Duration(24) * time.Hour)
		if inWeekdays(a.WakeDays, now.Weekday()) {
			a.nextrun = time.Date(now.Year(), now.Month(), now.Day(), a.WakeHour, a.WakeMinute, 0, 0, now.Location())
			return
		}
	}
}

func inWeekdays(h []time.Weekday, n time.Weekday) bool {
	for _, d := range h {
		if d == n {
			return true
		}
	}
	return false
}

func (a *alarm) UntilReadable() string {
	a.calcNextRun()
	duration := durafmt.Parse(time.Until(a.nextrun)).String()
	// dur := strings.Split(duration, " ")
	if len(duration) > 1 {
		return duration[:strings.Index(duration, "minutes")+len("minutes")]
	} else {
		return duration
	}
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
	// err := a.mc.Clear()
	// if err != nil {
	// 	log.Println(err.Error())
	// 	return
	// }
	// a.mc.PlaylistLoad(a.playlistName, -1, -1) // loads the playlist that we want to wake up to
	// a.mc.SetVolume(a.startVolume)             // start at stgartvolume
	// volStep := a.duration.Nanoseconds() / int64(a.endVolume-a.startVolume)
	// a.mc.Play(0)
	// a.mc.Random(true)
	// a.mc.Shuffle(-1, -1) // we don't want to have the same seed all the time, this might be an issue with mopidy
	// for i := a.startVolume; i < a.endVolume; i++ {
	// 	a.mc.SetVolume(i)
	// 	time.Sleep(time.Duration(volStep) * time.Nanosecond)
	// }
	// // and now we should be awake and beautiful :D
}
