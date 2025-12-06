package container

import (
	"github.com/AndyS1mpson/key-value-database/internal/utils/container"
	"github.com/AndyS1mpson/key-value-database/internal/utils/dotenv"
)

// Container service DI-container
type Container struct {
	config AppConfig
	*container.Container
}

func NewContainer(config AppConfig) (c *Container, gracefulShutdown func()) {
	dotenv.Load()

	sc, gracefulShutdown := container.NewContainer()

	c = &Container{
		config:    config,
		Container: sc,
	}

	return
}

// GetConfig get filled configuraion
func (c *Container) GetConfig() AppConfig {
	return c.config
}
