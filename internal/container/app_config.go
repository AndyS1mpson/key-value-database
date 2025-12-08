package container

import (
	"github.com/AndyS1mpson/key-value-database/internal/utils/dotenv"
	"github.com/kelseyhightower/envconfig"
)

type AppConfig struct{}

// LoadConfig load configuration from environment.
func LoadConfig() AppConfig {
	dotenv.Load()

	var appConfig AppConfig

	envconfig.MustProcess("", &appConfig)

	return appConfig
}
