package container

import (
	"github.com/AndyS1mpson/key-value-database/internal/database"
	"github.com/AndyS1mpson/key-value-database/internal/database/compute/parser"
	"github.com/AndyS1mpson/key-value-database/internal/database/storage"
	"github.com/AndyS1mpson/key-value-database/internal/database/storage/engine/in_memory"
	"github.com/AndyS1mpson/key-value-database/internal/utils/container"
)

func (c *Container) getParser() *parser.QueryParser {
	return container.MustOrGetNew(c.Container, func() *parser.QueryParser {
		return parser.New(c.GetLogger())
	})
}

func (c *Container) getInMemoryStorageEngine() *in_memory.Engine {
	return container.MustOrGetNew(c.Container, func() *in_memory.Engine {
		return in_memory.NewEngine(c.GetLogger())
	})
}

func (c *Container) getStorage() *storage.Storage {
	return container.MustOrGetNew(c.Container, func() *storage.Storage {
		return storage.NewStorage(c.getInMemoryStorageEngine())
	})
}

func (c *Container) GetDatabase() *database.Database {
	return container.MustOrGetNew(c.Container, func() *database.Database {
		return database.New(c.getParser(), c.getStorage(), c.GetLogger())
	})
}
