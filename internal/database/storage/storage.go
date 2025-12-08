package storage

import (
	"context"
	"errors"
)

var (
	ErrorNotFound = errors.New("not found")
)

// Storage implement working with engine
type Storage struct {
	engine engine
}

func NewStorage(engine engine, options ...StorageOption) *Storage {
	storage := &Storage{
		engine: engine,
	}

	for _, option := range options {
		option(storage)
	}

	return storage
}

func (s *Storage) Get(ctx context.Context, key string) (string, error) {
	value, found := s.engine.Get(ctx, key)
	if !found {
		return "", ErrorNotFound
	}

	return value, nil
}

func (s *Storage) Set(ctx context.Context, key string, value string) error {
	s.engine.Set(ctx, key, value)

	return nil
}

func (s *Storage) Del(ctx context.Context, key string) error {
	s.engine.Del(ctx, key)

	return nil
}
