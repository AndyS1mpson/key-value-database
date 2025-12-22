package container

import (
	"context"
	"sync"

	"github.com/AndyS1mpson/key-value-database/internal/utils/datastructs"
	"github.com/AndyS1mpson/key-value-database/internal/utils/shutdown"
)

// ServiceName is a key type for resolving dependency instances in the container.
type ServiceName string

// Container is a dependency injection container that manages service instances and graceful shutdown.
type Container struct {
	appCtx context.Context

	services sync.Map
	shutdown *datastructs.Stack[func()]
}

// NewContainer creates a new container instance with graceful shutdown support.
func NewContainer() (c *Container, gracefulShutdown func()) {
	ctx, cancel := shutdown.WithCancel(context.Background())

	c = &Container{
		appCtx:   ctx,
		shutdown: &datastructs.Stack[func()]{},
	}

	gracefulShutdown = c.gracefulShutdown(cancel)

	return
}

// Ctx returns the application context that can be cancelled for graceful shutdown.
func (c *Container) Ctx() context.Context {
	return c.appCtx
}

// PushShutdown adds a shutdown function to the stack for execution during graceful shutdown.
func (c *Container) PushShutdown(f func()) {
	c.shutdown.Push(f)
}

// gracefulShutdown returns a function that executes all shutdown handlers in reverse order.
func (c *Container) gracefulShutdown(ctxCancel func()) func() {
	return func() {
		for !c.shutdown.IsEmpty() {
			c.shutdown.Pop()()
		}

		ctxCancel()
	}
}
