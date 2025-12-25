package in_memory

import "github.com/AndyS1mpson/key-value-database/internal/utils/concurrency"

// Option is a function type for configuring Engine instances.
type Option func(*Engine)

// WithPartitions configures the Engine to use the specified number of partitions for data distribution.
func WithPartitions(partitionsNumber uint) Option {
	return func(e *Engine) {
		e.partitions = make([]*concurrency.HashTable, partitionsNumber)
		for i := 0; i < int(partitionsNumber); i++ {
			e.partitions[i] = concurrency.NewHashTable()
		}
	}
}
