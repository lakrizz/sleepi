package entities_test

import (
	"testing"
	"time"

	"github.com/google/uuid"
	"github.com/k0kubun/pp"
	"github.com/stretchr/testify/assert"

	"github.com/lakrizz/sleepi/internal/domain/alarms/entities"
	"github.com/lakrizz/sleepi/internal/domain/shared"
)

func TestTimeUntilNextTrigger(t *testing.T) {
	timeofday, err := shared.ParseTimeOfDay("10:30")
	assert.NoError(t, err)
	alarm, err := entities.NewAlarm("Foo", timeofday, shared.Playable{
		Name: "foo",
		ID:   uuid.New(),
	}, []shared.Weekday{shared.Tuesday, shared.Sunday}, 1*time.Hour, nil)
	assert.NoError(t, err)

	pp.Println(alarm.TimeUntilNextTrigger().String())
}
