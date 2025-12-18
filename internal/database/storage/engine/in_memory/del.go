package in_memory

import (
	"context"

	"go.uber.org/zap"

	"github.com/AndyS1mpson/key-value-database/internal/common"
)

// Del removes a key-value pair from the in-memory storage by key.
func (e *Engine) Del(ctx context.Context, key string) {
	e.data.Del(key)

	txID := common.GetTxIDFromContext(ctx)
	e.logger.Debug("successfull del query", zap.Int64("tx", txID))
}
