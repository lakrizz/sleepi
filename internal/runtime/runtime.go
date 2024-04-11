package runtime

import (
	"log"

	"github.com/lakrizz/sleepi/config"
	"github.com/lakrizz/sleepi/internal"
	"github.com/lakrizz/sleepi/internal/services"
)

type Runtime struct {
	Alarms    internal.AlarmService
	Playlists internal.PlaylistService
	Library   internal.LibraryService
}

func InitRuntime(cfg *config.Config) (*Runtime, error) {
	log.Println("initializing alarmservice...")
	alarmService, err := services.NewAlarmService(cfg)
	if err != nil {
		return nil, err
	}

	log.Println("initializing playlistservice...")
	playlistService, err := services.NewPlaylistService(cfg)
	if err != nil {
		return nil, err
	}

	log.Println("initializing libraryservice...")
	libraryService, err := services.NewLibraryService(cfg)
	if err != nil {
		return nil, err
	}

	rt := &Runtime{
		Alarms:    alarmService,
		Playlists: playlistService,
		Library:   libraryService,
	}

	return rt, nil
}
