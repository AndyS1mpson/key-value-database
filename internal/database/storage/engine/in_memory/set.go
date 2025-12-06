package in_memory

import (
	"context"

	"github.com/AndyS1mpson/key-value-database/internal/common"
	"go.uber.org/zap"
)

// Set save data to database
func (e *Engine) Set(ctx context.Context, key, value string) {
	partitionIdx := e.getPartitionIdx(key)
	e.partitions[partitionIdx].Set(key, value)

	txID := common.GetTxIDFromContext(ctx)


	e.logger.Debug("successfull set query", zap.Int64("tx", txID))
}
