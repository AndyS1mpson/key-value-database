package container

import (
	"github.com/AndyS1mpson/key-value-database/internal/utils/container"
	"github.com/AndyS1mpson/key-value-database/internal/utils/dotenv"
)

// Container is a dependency injection container that manages service lifecycle and dependencies.
type Container struct {
	config AppConfig
	*container.Container
}

// NewContainer creates a new service container with the given configuration.
// Returns the container and a graceful shutdown function.
func NewContainer(config AppConfig) (c *Container, gracefulShutdown func()) {
	dotenv.Load()

	sc, gracefulShutdown := container.NewContainer()

	c = &Container{
		config:    config,
		Container: sc,
	}

	return
}

// GetConfig returns the application configuration.
func (c *Container) GetConfig() AppConfig {
	return c.config
}
