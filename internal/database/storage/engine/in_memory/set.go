package in_memory

import (
	"context"

	"go.uber.org/zap"

	"github.com/AndyS1mpson/key-value-database/internal/common"
)

// Set stores a key-value pair in the in-memory storage.
func (e *Engine) Set(ctx context.Context, key, value string) {
	e.data.Set(key, value)

	txID := common.GetTxIDFromContext(ctx)
	e.logger.Debug("successfull set query", zap.Int64("tx", txID))
}
