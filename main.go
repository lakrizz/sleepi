package main

import (
	"fmt"
	"time"

	"./lib"
)

func main() {
	Init()
	a, e := lib.CreateAlarm(time.Now(), nil)
	if e != nil {
		panic(e)
	}

	Manager.AddAlarm(a)

	fmt.Println(Manager)

}
