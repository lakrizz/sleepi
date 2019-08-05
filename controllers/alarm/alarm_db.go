package alarm

import (
	"encoding/json"
	"io/ioutil"
)

const (
	db_file = "alarms.db"
)

func (a *alarmModule) Save() error {
	data, err := json.Marshal(a.alarms)
	if err != nil {
		return err
	}
	err = ioutil.WriteFile(db_file, data, 0x777)
	if err != nil {
		return err
	}
	return nil
}

func (a *alarmModule) Load() error {
	var alarms []*alarm

	data, err := ioutil.ReadFile(db_file)
	if err != nil {
		return err
	}
	err = json.Unmarshal(data, &alarms)
	if err != nil {
		return err
	}
	a.alarms = alarms
	return nil
}

func (a *alarmModule) AddAlarm(al *alarm) error {
	a.alarms = append(a.alarms, al)
	return a.Save()
}
