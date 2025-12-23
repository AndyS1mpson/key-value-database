package container

import (
	"github.com/AndyS1mpson/key-value-database/internal/database"
	"github.com/AndyS1mpson/key-value-database/internal/database/compute/parser"
	"github.com/AndyS1mpson/key-value-database/internal/database/storage"
	"github.com/AndyS1mpson/key-value-database/internal/database/storage/engine/in_memory"
	"github.com/AndyS1mpson/key-value-database/internal/database/storage/replication"
	"github.com/AndyS1mpson/key-value-database/internal/utils/container"
	"github.com/samber/lo"
)

// getParser creates a QueryParser instance for parsing database queries.
func (c *Container) getParser() *parser.QueryParser {
	return container.MustOrGetNew(c.Container, func() *parser.QueryParser {
		return parser.New(c.GetLogger())
	})
}

// getInMemoryStorageEngine creates an in-memory storage engine instance.
func (c *Container) getInMemoryStorageEngine() *in_memory.Engine {
	return container.MustOrGetNew(c.Container, func() *in_memory.Engine {
		options := []in_memory.Option{
			in_memory.WithPartitions(lo.FromPtr(c.config.Engine.PartitionsNumber)),
		}

		return in_memory.NewEngine(c.GetLogger(), options...)
	})
}

// getStorage creates a Storage instance with WAL support configured.
func (c *Container) getStorage() *storage.Storage {
	return container.MustOrGetNew(c.Container, func() *storage.Storage {
		options := []storage.Option{
			storage.WithWAL(c.getWAL()),
		}

		if c.config.Replication.ReplicaType == replication.TypeMaster {
			options = append(options, storage.WithReplication(c.getMasterReplica()))
		} else {
			options = append(options, storage.WithReplication(c.getSlaveReplica()))
			options = append(options, storage.WithReplicationStream(c.getSlaveReplica().ReplicationStream()))

		}

		return storage.NewStorage(c.getInMemoryStorageEngine(), c.GetLogger(), options...)
	})
}

// GetDatabase creates and returns the main Database instance with all dependencies configured.
func (c *Container) GetDatabase() *database.Database {
	return container.MustOrGetNew(c.Container, func() *database.Database {
		return database.New(c.getParser(), c.getStorage(), c.GetLogger())
	})
}
