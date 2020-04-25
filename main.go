package main

import (
	"fmt"

	"github.com/lakrizz/sleepi/alarm"
	"github.com/lakrizz/sleepi/config"
	"github.com/lakrizz/sleepi/player"
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

	am, err := alarm.CreateAlarmManager("alarms.json")
	if err != nil {
		panic(err)
	}

	fmt.Println("all alarms:", am.Alarms)

	// panic(http.ListenAndServe(":8080", nil))
	for {
	}
}
