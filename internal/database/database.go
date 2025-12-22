package database

import (
	"context"
	"fmt"

	"github.com/AndyS1mpson/key-value-database/internal/database/compute/models"
	"github.com/AndyS1mpson/key-value-database/internal/database/storage"
	"go.uber.org/zap"
)

// Database provides methods for working with key-value storage.
// It coordinates between the compute layer (query parsing) and storage layer (data operations).
type Database struct {
	computeLayer computeLayer
	storageLayer storageLayer
	logger       *zap.Logger
}

// New creates a new Database instance with the given compute layer, storage layer, and logger.
func New(computeLayer computeLayer, storageLayer storageLayer, logger *zap.Logger) *Database {
	return &Database{
		computeLayer: computeLayer,
		storageLayer: storageLayer,
		logger:       logger,
	}
}

// HandleQuery processes a raw query string, parses it, and executes the corresponding command.
// Returns a formatted response string.
func (d *Database) HandleQuery(ctx context.Context, rawQuery string) string {
	d.logger.Debug("handling query", zap.String("query", rawQuery))

	query, err := d.computeLayer.Parse(rawQuery)
	if err != nil {
		return fmt.Sprintf("[error] %s", err.Error())
	}

	switch query.Command {
	case models.CommandSet:
		return d.handleSetQuery(ctx, *query)
	case models.CommandGet:
		return d.handleGetQuery(ctx, *query)
	case models.CommandDel:
		return d.handleDelQuery(ctx, *query)
	}

	d.logger.Error("compute layer is incorrect", zap.String("command", string(query.Command)))

	return "[error] internal error"
}

// handleSetQuery executes a SET command to store a key-value pair.
func (d *Database) handleSetQuery(ctx context.Context, query models.Query) string {
	arguments := query.Args
	if err := d.storageLayer.Set(ctx, arguments[0], arguments[1]); err != nil {
		return fmt.Sprintf("[error] %s", err.Error())
	}

	return "[ok]"
}

// handleGetQuery executes a GET command to retrieve a value by key.
func (d *Database) handleGetQuery(ctx context.Context, query models.Query) string {
	arguments := query.Args
	value, err := d.storageLayer.Get(ctx, arguments[0])
	if err == storage.ErrorNotFound {
		return "[not found]"
	} else if err != nil {
		return fmt.Sprintf("[error] %s", err.Error())
	}

	return fmt.Sprintf("[ok] %s", value)
}

// handleDelQuery executes a DEL command to delete a key-value pair.
func (d *Database) handleDelQuery(ctx context.Context, query models.Query) string {
	arguments := query.Args
	if err := d.storageLayer.Del(ctx, arguments[0]); err != nil {
		return fmt.Sprintf("[error] %s", err.Error())
	}

	return "[ok]"
}
