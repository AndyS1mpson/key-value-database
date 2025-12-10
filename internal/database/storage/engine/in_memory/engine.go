package in_memory

import (
	"github.com/AndyS1mpson/key-value-database/internal/utils/concurrency"
	"go.uber.org/zap"
)

// Engine in-memory database mock
type Engine struct {
	logger *zap.Logger

	data concurrency.HashTable
}

func NewEngine(logger *zap.Logger) *Engine {
	return &Engine{
		logger: logger,
		data: *concurrency.NewHashTable(),
	}
}
