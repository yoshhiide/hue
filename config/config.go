package config

import "os"

type (
	Config struct {
		BridgeURL string
		Token     string
	}
)

var (
	config *Config
)

func Load() (*Config, error) {
	if config == nil {
		config = &Config{
			BridgeURL: os.Getenv("HUE_BRIDGE_URL"),
			Token:     os.Getenv("HUE_BRIDGE_TOKEN"),
		}
	}

	return config, nil
}
