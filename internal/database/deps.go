package database

import (
	"context"

	"github.com/AndyS1mpson/key-value-database/internal/database/compute/models"
)

//go:generate mockgen -source=deps.go -destination=./mocks/mock.go

type computeLayer interface {
	Parse(string) (*models.Query, error)
}

type storageLayer interface {
	Set(context.Context, string, string) error
	Get(context.Context, string) (string, error)
	Del(context.Context, string) error
}
