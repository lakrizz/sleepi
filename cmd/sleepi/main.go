package main

import (
	"fmt"

	"github.com/lakrizz/sleepi/pkg/alarm"
	"github.com/lakrizz/sleepi/pkg/config"
	"github.com/lakrizz/sleepi/pkg/player"
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

	am, err := alarm.CreateAlarmManager()
	if err != nil {
		panic(err)
	}

	fmt.Println("all alarms:", am.Alarms)

	// this starts the loop, nothing more needed
	web.Serve()
}
