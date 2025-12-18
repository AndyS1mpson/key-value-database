package storage

// StorageOption is a function type for configuring Storage instances.
type StorageOption func(*Storage)

// WithWAL configures Storage to use the provided WAL implementation.
func WithWAL(wal wal) StorageOption {
	return func(storage *Storage) {
		storage.wal = wal
	}
}
