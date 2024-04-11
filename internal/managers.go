package internal

import (
	"github.com/google/uuid"

	"github.com/lakrizz/sleepi/pkg/alarm"
	"github.com/lakrizz/sleepi/pkg/models"
	"github.com/lakrizz/sleepi/pkg/playlist"
)

type PlaylistService interface {
	AddPlaylist(*playlist.Playlist) error
	GetPlaylistById(uuid.UUID) (*playlist.Playlist, error)
	GetAllPlaylists() ([]*playlist.Playlist, error)

	Save() error
}

type AlarmService interface {
	AddAlarm(*alarm.Alarm) error
	GetAlarm(uuid.UUID) (*alarm.Alarm, error)
	UpdateTimings() error
	GetAllAlarms() ([]*alarm.Alarm, error)

	Save() error
}

type LibraryService interface {
	AddFile([]byte, string) error
	RemoveByID(uuid.UUID) error
	GetFileByID(uuid.UUID) (*models.File, error)
	GetFiles() []*models.File

	Save() error
}
