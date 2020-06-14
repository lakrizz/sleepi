package api

import (
	"github.com/lakrizz/sleepi/pkg/alarm"
	"github.com/lakrizz/sleepi/pkg/library"
	"github.com/lakrizz/sleepi/pkg/playlist"
)

type Api struct {
	Playlists *playlist.PlaylistManager
	Alarms    *alarm.AlarmManager
	Library   *library.Library
}
