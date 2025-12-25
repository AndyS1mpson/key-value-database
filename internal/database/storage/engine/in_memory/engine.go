package in_memory

import (
	"hash/fnv"

	"go.uber.org/zap"

	"github.com/AndyS1mpson/key-value-database/internal/utils/concurrency"
)

// Engine provides an in-memory key-value storage implementation using a thread-safe hash table.
type Engine struct {
	logger *zap.Logger

	partitions []*concurrency.HashTable
}

// NewEngine creates a new in-memory storage engine instance.
func NewEngine(logger *zap.Logger, options ...Option) *Engine {
	engine := &Engine{
		logger: logger,
	}

	for _, option := range options {
		option(engine)
	}

	if len(engine.partitions) == 0 {
		engine.partitions = make([]*concurrency.HashTable, 1)
		engine.partitions[0] = concurrency.NewHashTable()
	}

	return engine
}

func (e *Engine) partitionIdx(key string) int {
	hash := fnv.New32a()
	_, _ = hash.Write([]byte(key))
	return int(hash.Sum32()) % len(e.partitions)
}
