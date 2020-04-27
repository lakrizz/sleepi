package config

import (
	"encoding/json"
	"io/ioutil"

	"github.com/adrg/xdg"
)

type Config struct {
	filename string   `json:"-"`
	Volumes  *volumes `json:"Volumes"`
}

type volumes struct {
	Silence float64 `json:"silence"`
	Normal  float64 `json:"normal"`
}

func createEmptyConfig() *Config {
	// we assume the file does not exist
	return &Config{
		Volumes: &volumes{
			Silence: -5.0,
			Normal:  2.0,
		},
	}
}

func LoadConfig() (*Config, error) {
	filename, err := xdg.ConfigFile("sleepi/config.json")
	if err != nil {
		return nil, err
	}

	dat, err := ioutil.ReadFile(filename)
	if err != nil {
		return createEmptyConfig(), nil
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
