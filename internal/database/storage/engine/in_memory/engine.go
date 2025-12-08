package in_memory

import (
	"go.uber.org/zap"
)

// Engine in-memory database mock
type Engine struct {
	logger *zap.Logger

	data map[string]string
}

func NewEngine(logger *zap.Logger) *Engine {
	return &Engine{
		logger: logger,
		data:   make(map[string]string),
	}
}
