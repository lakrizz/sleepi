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

	pl, _ := lib.CreatePlaylist()
	pl.AddVideo("https://www.youtube.com/watch?v=u5CVsCnxyXg")

	Manager.AddPlaylist(pl)
	fmt.Println(Manager)
	fmt.Println(Manager.playlists[0])

}
