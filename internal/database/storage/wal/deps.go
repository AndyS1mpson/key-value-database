package wal

import "github.com/AndyS1mpson/key-value-database/internal/database/storage/wal/log"

//go:generate mockgen -source=deps.go -destination=./mocks/mock.go

// logsWriter defines the interface for writing WAL log batches.
type logsWriter interface {
	Write([]log.WriteRequest)
}

// logsReader defines the interface for reading WAL logs.
type logsReader interface {
	Read() ([]log.Log, error)
}
