package in_memory

import (
	"context"

	"go.uber.org/zap"

	"github.com/AndyS1mpson/key-value-database/internal/common"
)

// Del removes a key-value pair from the in-memory storage by key.
func (e *Engine) Del(ctx context.Context, key string) {
	partitionIdx := 0
	if len(e.partitions) > 1 {
		partitionIdx = e.partitionIdx(key)
	}

	partition := e.partitions[partitionIdx]
	partition.Del(key)

	txID := common.GetTxIDFromContext(ctx)
	e.logger.Debug("successfull del query", zap.Int64("tx", txID))
}
