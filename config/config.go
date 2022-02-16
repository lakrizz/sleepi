package config

type config struct {
	MPD_PROTOCOL string
	MPD_HOST     string

	HTTP_HOST string
	HTTP_PORT string
}

const (
	CONFIG_PATH = "./config/"
)

func GetConfig() (*config, error) {
	cfg := &config{}
	cfg.MPD_HOST = "localhost:6600"
	cfg.MPD_PROTOCOL = "tcp"

	cfg.HTTP_PORT = "8080"
	cfg.HTTP_HOST = "localhost"

	return cfg, nil
}
