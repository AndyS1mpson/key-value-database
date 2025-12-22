package generator

import (
	"math"
	"sync/atomic"
)

// IDGenerator generates unique transaction IDs (LSN - Log Sequence Numbers) in a thread-safe manner.
type IDGenerator struct {
	counter atomic.Int64
}

// NewIDGenerator creates a new ID generator starting from the given previous ID.
func NewIDGenerator(previousID int64) *IDGenerator {
	generator := &IDGenerator{}
	generator.counter.Store(previousID)
	return generator
}

// Generate returns the next unique transaction ID.
// It wraps around to 0 when reaching MaxInt64.
func (g *IDGenerator) Generate() int64 {
	g.counter.CompareAndSwap(math.MaxInt64, 0)
	return g.counter.Add(1)
}
