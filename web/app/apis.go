package app

import (
	"krizz.org/sleepi/internal/manager"
	"krizz.org/sleepi/pkg/library"
)

var (
	alarmManager    *manager.AlarmManager
	playlistManager *manager.PlaylistManager
	lib             *library.Library
)

func AddApis(am *manager.AlarmManager, pm *manager.PlaylistManager, libr *library.Library) {
	alarmManager = am
	playlistManager = pm
	lib = libr
}
