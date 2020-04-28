package main

import (
	"github.com/lakrizz/sleepi/api"
	"github.com/lakrizz/sleepi/pkg/alarm"
	"github.com/lakrizz/sleepi/pkg/config"
	"github.com/lakrizz/sleepi/pkg/player"
	"github.com/lakrizz/sleepi/pkg/playlist"
	"github.com/lakrizz/sleepi/web"
)

func main() {
	// let's start anew, but this time let it grow gradually instead of chopping it all at once
	// s := gin.Default()
	c, err := config.LoadConfig()
	if err != nil {
		panic(err)
	}

	err = player.InitPlayer(c.Volumes.Silence, c.Volumes.Normal)
	if err != nil {
		panic(err)
	}

	alarmmanager, err := alarm.LoadAlarmManager()
	if err != nil {
		panic(err)
	}

	playlistmanager, err := playlist.LoadConfig()
	if err != nil {
		panic(err)
	}

	// 	pl_id, _ := uuid.Parse("ee1689a9-563f-4886-a95c-1cfbde97b0f6")

	// 	i, _ := uuid.NewRandom()
	// 	na := &alarm.Alarm{
	// 		Name:       "foobar",
	// 		Id:         i,
	// 		WakeHour:   7,
	// 		WakeMinute: 30,
	// 		Days:       []time.Weekday{0, 1, 2, 3, 4, 5},
	// 		Playlist:   pl_id,
	// 		WakeupTime: "30m",
	// 		Enabled:    true,
	// 	}
	// 	err = alarmmanager.AddAlarm(na)
	// 	alarmmanager.SaveAlarms()

	api := &api.Api{
		Playlists: playlistmanager,
		Alarms:    alarmmanager,
	}

	// this starts the loop, nothing more needed
	web.Serve(api)
}
