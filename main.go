package main

import (
	"./alarm"
	"./config"
	"./player"
)

func main() {
	// let's start anew, but this time let it grow gradually instead of chopping it all at once
	// s := gin.Default()
	c, err := config.LoadConfig("config.json")
	if err != nil {
		panic(err)
	}

	err = player.InitPlayer(c.Volumes.Silence, c.Volumes.Normal)
	if err != nil {
		panic(err)
	}

	_, err = alarm.CreateAlarmManager("alarms.json")
	if err != nil {
		panic(err)
	}

	for {
	}
}
