package main

import (
	"krizz.org/sleepi/internal/manager"
	"krizz.org/sleepi/pkg/library"
	"krizz.org/sleepi/web"
	"krizz.org/sleepi/web/app"
)

func main() {

	alarmManager, err := manager.GetAlarmManager()
	if err != nil {
		panic(err)
	}

	playlistManager, err := manager.GetPlaylistManager()
	if err != nil {
		panic(err)
	}

	library, err := library.GetLibrary()
	if err != nil {
		panic(err)
	}

	app.AddApis(alarmManager, playlistManager, library)
	web.Serve()

	for {
	}
}
