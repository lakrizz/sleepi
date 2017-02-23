package lib

import (
	"time"

	"github.com/nu7hatch/gouuid"
)

type Alarm struct {
	id         *uuid.UUID
	time       time.Time
	enabled    bool
	playlistid *uuid.UUID
}

func CreateAlarm(time time.Time, playlistid *uuid.UUID) (*Alarm, error) {
	u, err := uuid.NewV4()
	if err != nil {
		return nil, err
	}
	return &Alarm{u, time, true, playlistid}, nil
}

func Delete(id *uuid.UUID) {
	// call the manager :D
}
