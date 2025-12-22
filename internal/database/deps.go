package database

import (
	"context"

	"github.com/AndyS1mpson/key-value-database/internal/database/compute/models"
)

//go:generate mockgen -source=deps.go -destination=./mocks/mock.go

// computeLayer defines the interface for parsing raw query strings into structured queries.
type computeLayer interface {
	Parse(string) (*models.Query, error)
}

// storageLayer defines the interface for key-value storage operations.
type storageLayer interface {
	Set(context.Context, string, string) error
	Get(context.Context, string) (string, error)
	Del(context.Context, string) error
}
