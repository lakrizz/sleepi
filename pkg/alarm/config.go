package alarm

import (
	"encoding/json"
	"fmt"
	"io/ioutil"
	"log"

	"github.com/adrg/xdg"
)

type AlarmManager struct {
	Alarms  []*Alarm      `json:"Alarms"`
	watcher *alarmWatcher `json:"-"`
}

var filename string = "alarms.json"

func LoadAlarmManager() (*AlarmManager, error) {
	fname, err := xdg.ConfigFile(fmt.Sprintf("sleepi/%s", filename))
	if err != nil {
		return nil, err
	}

	dat, err := ioutil.ReadFile(fname)
	if err != nil {
		am := &AlarmManager{Alarms: []*Alarm{}}
		am.SaveAlarms()
		return am, nil
	}

	log.Println("data", string(dat))

	var am *AlarmManager
	err = json.Unmarshal(dat, &am)

	if err != nil {
		return nil, err
	}

	if len(am.Alarms) > 0 { // if there's no alarms, there's nothing to watch
		aw, err := createWatcher(am)
		if err != nil {
			return nil, err
		}
		am.watcher = aw
		go am.watcher.run()
	}

	return am, nil
}

func (am *AlarmManager) SaveAlarms() error {
	fname, err := xdg.ConfigFile(fmt.Sprintf("sleepi/%s", filename))
	if err != nil {
		return err
	}

	dat, err := json.Marshal(am)

	if err != nil {
		return err
	}

	return ioutil.WriteFile(fname, dat, 0755)
}
