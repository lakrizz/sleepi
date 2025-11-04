package app

import "github.com/lakrizz/sleepi/internal/usecases/alarms"

func (a *App) initUsecases() error {
	alarmUsecases, err := alarms.NewAlarmUsecases(a.AlarmRepository)
	if err != nil {
		return err
	}

	a.AlarmUsecases = alarmUsecases
	return nil
}
