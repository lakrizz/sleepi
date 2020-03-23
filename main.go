package main

import (
	"./alarm"
	"./config"
	"./player"
	"github.com/k0kubun/pp"
)

func main() {
	// let's start anew, but this time let it grow gradually instead of chopping it all at once
	// s := gin.Default()
	am, err := alarm.CreateAlarmManager("alarms.json")
	if err != nil {
		panic(err)
	}

	v, err := am.GetNextAlarm()
	if err != nil {
		panic(err)
	}
	pp.Println(v)
	return

	// nw, err := am[0].TimeTillNextWake()
	// pp.Println(nw.String())
	// return

	c, err := config.LoadConfig("config.json")
	if err != nil {
		panic(err)
	}

	pp.Println(c)
	p, err := player.GetPlayer(c.Volumes.Silence, c.Volumes.Normal)
	if err != nil {
		panic(err)
	}

	err = p.LoadPlaylist("pl.json", false)
	if err != nil {
		panic(err)
	}

	err = p.Play()
	panic(err)
}
