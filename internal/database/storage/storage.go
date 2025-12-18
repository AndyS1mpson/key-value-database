package storage

import (
	"context"
	"errors"

	"go.uber.org/zap"

	"github.com/AndyS1mpson/key-value-database/internal/common"
	"github.com/AndyS1mpson/key-value-database/internal/database/compute/models"
	"github.com/AndyS1mpson/key-value-database/internal/database/storage/generator"
	walLog "github.com/AndyS1mpson/key-value-database/internal/database/storage/wal/log"
)

var (
	ErrorNotFound = errors.New("not found")
)

// Storage coordinates between the storage engine and WAL (Write-Ahead Log).
// It ensures data durability by writing to WAL before applying changes to the engine.
type Storage struct {
	engine    engine
	wal       wal
	generator generator.IDGenerator

	logger *zap.Logger
}

// NewStorage creates a new Storage instance with the given engine and logger.
// It recovers data from WAL if available and initializes the transaction ID generator.
func NewStorage(engine engine, logger *zap.Logger, options ...StorageOption) *Storage {
	storage := &Storage{
		engine: engine,
	}

	for _, option := range options {
		option(storage)
	}

	var lastLSN int64
	if storage.wal != nil {
		logs, err := storage.wal.Recover()
		if err != nil {
			logger.Error("failed to recover data from WAL", zap.Error(err))
		} else {
			lastLSN = storage.applyData(logs)
		}
	}

	storage.generator = *generator.NewIDGenerator(lastLSN)

	return storage
}

// Get retrieves a value by key from the storage engine.
func (s *Storage) Get(ctx context.Context, key string) (string, error) {
	txID := s.generator.Generate()
	ctx = common.ContextWithTxID(ctx, txID)

	value, found := s.engine.Get(ctx, key)
	if !found {
		return "", ErrorNotFound
	}

	return value, nil
}

// Set stores a key-value pair. It writes to WAL first (if enabled) for durability,
// then applies the change to the storage engine.
func (s *Storage) Set(ctx context.Context, key string, value string) error {
	txID := s.generator.Generate()
	ctx = common.ContextWithTxID(ctx, txID)

	if s.wal != nil {
		futureResponse := s.wal.Set(ctx, key, value)
		if err := futureResponse.Get(); err != nil {
			return err
		}
	}

	s.engine.Set(ctx, key, value)

	return nil
}

// Del removes a key-value pair. It writes to WAL first (if enabled) for durability,
// then applies the deletion to the storage engine.
func (s *Storage) Del(ctx context.Context, key string) error {
	txID := s.generator.Generate()
	ctx = common.ContextWithTxID(ctx, txID)

	if s.wal != nil {
		futureResponse := s.wal.Del(ctx, key)
		if err := futureResponse.Get(); err != nil {
			return err
		}
	}

	s.engine.Del(ctx, key)

	return nil
}

// applyData replays WAL logs to restore the in-memory storage state.
// Returns the highest LSN (Log Sequence Number) found in the logs.
func (s *Storage) applyData(logs []walLog.Log) int64 {
	var lastLSN int64
	for _, log := range logs {
		lastLSN = max(lastLSN, log.LSN)
		ctx := common.ContextWithTxID(context.Background(), log.LSN)

		switch log.Query.Command {
		case models.CommandSet:
			s.engine.Set(ctx, log.Query.Args[0], log.Query.Args[1])
		case models.CommandDel:
			s.engine.Del(ctx, log.Query.Args[0])
		}
	}

	return lastLSN
}
