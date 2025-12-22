package storage

import (
	"context"

	"github.com/AndyS1mpson/key-value-database/internal/database/storage/wal/log"
	"github.com/AndyS1mpson/key-value-database/internal/utils/concurrency"
)

//go:generate mockgen -source=deps.go -destination=./mocks/mock.go

// engine defines the interface for low-level storage operations.
type engine interface {
	Set(context.Context, string, string)
	Get(context.Context, string) (string, bool)
	Del(context.Context, string)
}

// wal defines the interface for Write-Ahead Logging operations.
type wal interface {
	Recover() ([]log.Log, error)
	Set(ctx context.Context, key, value string) concurrency.FutureError
	Del(ctx context.Context, key string) concurrency.FutureError
}
