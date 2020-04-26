package config

import (
	"encoding/json"
	"io/ioutil"
)

type Config struct {
	filename string   `json:"-"`
	Volumes  *volumes `json:"Volumes"`
}

type volumes struct {
	Silence float64 `json:"silence"`
	Normal  float64 `json:"normal"`
}

func LoadConfig(filename string) (*Config, error) { // we expect the config to be in the same dir
	dat, err := ioutil.ReadFile(filename)
	if err != nil {
		return nil, err
	}

	var c *Config
	err = json.Unmarshal(dat, &c)
	if err != nil {
		return nil, err
	}
	c.filename = filename
	return c, nil
}

func (c *Config) Save() error {
	dat, err := json.Marshal(c)
	if err != nil {
		return err
	}
	err = ioutil.WriteFile(c.filename, dat, 755)
	return err
}
