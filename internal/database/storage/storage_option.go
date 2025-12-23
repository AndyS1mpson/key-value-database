package storage

import walLog "github.com/AndyS1mpson/key-value-database/internal/database/storage/wal/log"

// Option is a function type for configuring Storage instances.
type Option func(*Storage)

// WithWAL configures Storage to use the provided WAL implementation.
func WithWAL(wal wal) Option {
	return func(storage *Storage) {
		storage.wal = wal
	}
}

// WithReplication configures Storage to use the provided replication implementation.
func WithReplication(replica Replica) Option {
	return func(storage *Storage) {
		storage.replica = replica
	}
}

// WithReplicationStream configures Storage to receive replicated WAL logs from a master node.
func WithReplicationStream(stream <-chan []walLog.Log) Option {
	return func(storage *Storage) {
		storage.stream = stream
	}
}
