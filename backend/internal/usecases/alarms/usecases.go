package alarms

import (
	"github.com/lakrizz/sleepi/internal/repositories"
)

type AlarmsUseCases struct {
	AlarmsRepository *repositories.AlarmsRepository
}

func NewAlarmUsecases(alarmsRepository *repositories.AlarmsRepository) (*AlarmsUseCases, error) {
	auc := &AlarmsUseCases{
		AlarmsRepository: alarmsRepository,
	}

	return auc, nil
}
