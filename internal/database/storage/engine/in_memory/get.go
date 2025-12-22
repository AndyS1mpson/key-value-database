package in_memory

import (
	"context"

	"go.uber.org/zap"

	"github.com/AndyS1mpson/key-value-database/internal/common"
)

// Get retrieves a value by key from the in-memory storage.
// Returns the value and a boolean indicating if the key was found.
func (e *Engine) Get(ctx context.Context, key string) (string, bool) {
	value, found := e.data.Get(key)

	txID := common.GetTxIDFromContext(ctx)
	e.logger.Debug("successful get query", zap.Int64("tx", txID))

	return value, found
}
