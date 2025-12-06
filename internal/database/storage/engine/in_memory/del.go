package in_memory

import (
	"context"

	"github.com/AndyS1mpson/key-value-database/internal/common"
	"go.uber.org/zap"
)

// Del remove data by ker from database
func (e *Engine) Del(ctx context.Context, key string) {
	partitionIdx := e.getPartitionIdx(key)
	e.partitions[partitionIdx].Del(key)

	txID := common.GetTxIDFromContext(ctx)
	e.logger.Debug("successfull del query", zap.Int64("tx", txID))
}
