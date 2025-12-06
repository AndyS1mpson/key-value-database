package storage

import (
	"context"
	"errors"

	"github.com/AndyS1mpson/key-value-database/internal/common"
)

var (
	ErrorNotFound = errors.New("not found")
)

// Storage implement working with engine
type Storage struct {
	engine engine

	generator *IDGenerator
}

func NewStorage(engine engine, options ...StorageOption) *Storage {
	storage := &Storage{
		engine:    engine,
		generator: NewIDGenerator(0),
	}

	for _, option := range options {
		option(storage)
	}

	return storage
}

func (s *Storage) Get(ctx context.Context, key string) (string, error) {
	value, found := s.engine.Get(s.wrapTXId(ctx), key)
	if !found {
		return "", ErrorNotFound
	}

	return value, nil
}

func (s *Storage) Set(ctx context.Context, key string, value string) error {
	s.engine.Set(s.wrapTXId(ctx), key, value)

	return nil
}

func (s *Storage) Del(ctx context.Context, key string) error {
	s.engine.Del(s.wrapTXId(ctx), key)

	return nil
}

func (s *Storage) wrapTXId(ctx context.Context) context.Context {
	txID := s.generator.Generate()
	return common.ContextWithTxID(ctx, txID)
}
