package runtime

import (
	"context"
	"log"

	"github.com/lakrizz/sleepi/config"
	"github.com/lakrizz/sleepi/internal"
	"github.com/lakrizz/sleepi/internal/services"
	"github.com/lakrizz/sleepi/pkg/player"
	"github.com/lakrizz/sleepi/pkg/player/mpd"
)

type Runtime struct {
	Alarms    internal.AlarmService
	Playlists internal.PlaylistService
	Library   internal.LibraryService
	Player    player.Player
}

func InitRuntime(ctx context.Context, cfg *config.Config) (*Runtime, error) {
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

	log.Println("initializing player...")
	player, err := mpd.NewMPDPlayer(ctx)
	if err != nil {
		return nil, err
	}

	rt := &Runtime{
		Alarms:    alarmService,
		Playlists: playlistService,
		Library:   libraryService,
		Player:    player,
	}

	return rt, nil
}
