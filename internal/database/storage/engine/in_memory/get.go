package in_memory

import (
	"context"

	"github.com/AndyS1mpson/key-value-database/internal/common"
	"go.uber.org/zap"
)

// Get get data from database by key
func (e *Engine) Get(ctx context.Context, key string) (string, bool) {
	partitionIdx := e.getPartitionIdx(key)

	value, found := e.partitions[partitionIdx].Get(key)

	txID := common.GetTxIDFromContext(ctx)
	e.logger.Debug("successfull get query", zap.Int64("tx", txID))
	return value, found
}
