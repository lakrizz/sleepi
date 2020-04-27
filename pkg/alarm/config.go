package alarm

import (
	"encoding/json"
	"io/ioutil"
	"os"

	"github.com/adrg/xdg"
)

func loadAlarms() ([]*Alarm, error) {
	filename, err := xdg.ConfigFile("sleepi/config.json")
	if err != nil {
		return nil, err
	}

	alarms := make([]*Alarm, 0)
	if _, err := os.Stat(filename); err == os.ErrExist {
		dat, err := ioutil.ReadFile(filename)
		if err != nil {
			return nil, err
		}

		err = json.Unmarshal(dat, &alarms)
		if err != nil {
			return nil, err
		}
	}

	return alarms, nil
}

func saveAlarms(alarms []*Alarm) error {
	filename, err := xdg.ConfigFile("sleepi/config.json")
	if err != nil {
		return err
	}

	dat, err := json.Marshal(alarms)

	if err != nil {
		return err
	}

	return ioutil.WriteFile(filename, dat, 0644)
}
