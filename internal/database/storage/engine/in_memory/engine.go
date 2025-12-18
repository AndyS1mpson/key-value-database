package in_memory

import (
	"go.uber.org/zap"

	"github.com/AndyS1mpson/key-value-database/internal/utils/concurrency"
)

// Engine provides an in-memory key-value storage implementation using a thread-safe hash table.
type Engine struct {
	logger *zap.Logger

	data concurrency.HashTable
}

// NewEngine creates a new in-memory storage engine instance.
func NewEngine(logger *zap.Logger) *Engine {
	return &Engine{
		logger: logger,
		data:   *concurrency.NewHashTable(),
	}
}
