package in_memory

import (
	"hash/fnv"

	"github.com/AndyS1mpson/key-value-database/internal/utils/datastructs"
	"go.uber.org/zap"
)

// Engine in-memory database mock
type Engine struct {
	partitions []*datastructs.HashTable
	logger     *zap.Logger
}

func NewEngine(logger *zap.Logger, config Config) *Engine {
	partitions := make([]*datastructs.HashTable, 0, config.PartitionsCount)

	for range max(1, config.PartitionsCount) {
		partitions = append(partitions, datastructs.NewHashTable())
	}

	return &Engine{
		logger:     logger,
		partitions: partitions,
	}
}

// getPartitionIdx calculate and return partition index for key
func (e *Engine) getPartitionIdx(key string) int {
	hash := fnv.New32a()
	_, _ = hash.Write([]byte(key))

	return int(hash.Sum32()) % len(e.partitions)
}
