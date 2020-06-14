package config

import (
	"encoding/json"
	"fmt"
	"io/ioutil"
	"path"

	"github.com/adrg/xdg"
)

type Config struct {
	filename string   `json:"-"`
	Volumes  *volumes `json:"Volumes"`
	Library  string   `json:"Library"`
}

type volumes struct {
	Silence float64 `json:"silence"`
	Normal  float64 `json:"normal"`
}

func createEmptyConfig(filename string) *Config {
	// we assume the file does not exist
	return &Config{
		Volumes: &volumes{
			Silence: -5.0,
			Normal:  2.0,
		},
		filename: filename,
		Library:  path.Join(xdg.UserDirs.Documents, "sleepi"),
	}
}

func LoadConfig() (*Config, error) {
	filename, err := xdg.ConfigFile("sleepi/config.json")
	if err != nil {
		return nil, err
	}

	dat, err := ioutil.ReadFile(filename)
	if err != nil {
		cfg := createEmptyConfig(filename)
		return cfg, cfg.Save()
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
	fmt.Println(string(dat))
	err = ioutil.WriteFile(c.filename, dat, 0755)
	return err
}
