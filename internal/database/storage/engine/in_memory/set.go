package in_memory

import (
	"context"

	"go.uber.org/zap"

	"github.com/AndyS1mpson/key-value-database/internal/common"
)

// Set stores a key-value pair in the in-memory storage.
func (e *Engine) Set(ctx context.Context, key, value string) {
	partitionIdx := 0
	if len(e.partitions) > 1 {
		partitionIdx = e.partitionIdx(key)
	}

	partition := e.partitions[partitionIdx]
	partition.Set(key, value)

	txID := common.GetTxIDFromContext(ctx)
	e.logger.Debug("successfull set query", zap.Int64("tx", txID))
}
