package container

import (
	"context"
	"sync"

	"github.com/AndyS1mpson/key-value-database/internal/utils/datastructs"
	"github.com/AndyS1mpson/key-value-database/internal/utils/shutdown"
)

// ServiceName service key for resolving existed dependency instance
type ServiceName string

// Container service DI-container
type Container struct {
	appCtx context.Context

	services sync.Map
	shutdown *datastructs.Stack[func()]
}

func NewContainer() (c *Container, gracefulShutdown func()) {
	ctx, cancel := shutdown.WithCancel(context.Background())

	c = &Container{
		appCtx:   ctx,
		shutdown: &datastructs.Stack[func()]{},
	}

	gracefulShutdown = c.gracefulShutdown(cancel)

	return
}

// Ctx get app context
func (c *Container) Ctx() context.Context {
	return c.appCtx
}

// PushShutdown add cancel function to stack
func (c *Container) PushShutdown(f func()) {
	c.shutdown.Push(f)
}

func (c *Container) gracefulShutdown(ctxCancel func()) func() {
	return func() {
		for !c.shutdown.IsEmpty() {
			c.shutdown.Pop()()
		}

		ctxCancel()
	}
}
