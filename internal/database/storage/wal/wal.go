package wal

import (
	"context"
	"sync"
	"time"

	"github.com/AndyS1mpson/key-value-database/internal/common"
	"github.com/AndyS1mpson/key-value-database/internal/database/compute/models"
	"github.com/AndyS1mpson/key-value-database/internal/database/storage/wal/log"
	"github.com/AndyS1mpson/key-value-database/internal/utils/concurrency"
)

const (
	defaultFlushingBatchSize    = 100
	defaultFlushingBatchTimeout = time.Millisecond * 10
)

// WAL implements Write-Ahead Logging for durability.
// It batches write requests and flushes them periodically or when batch size is reached.
type WAL struct {
	logsWriter logsWriter
	logsReader logsReader

	flushTimeout time.Duration
	maxBatchSize int

	batches chan []log.WriteRequest
	mutex   sync.Mutex
	batch   []log.WriteRequest
}

// NewWAL creates a new WAL instance with the given writer, reader, flush timeout, and batch size.
// Uses default values if timeout or batch size is 0.
func NewWAL(writer logsWriter, reader logsReader, flushTimeout time.Duration, maxBatchSize int) *WAL {
	if flushTimeout == 0 {
		flushTimeout = defaultFlushingBatchTimeout
	}

	if maxBatchSize == 0 {
		maxBatchSize = defaultFlushingBatchSize
	}

	return &WAL{
		logsWriter:   writer,
		logsReader:   reader,
		flushTimeout: flushTimeout,
		maxBatchSize: maxBatchSize,
		batches:      make(chan []log.WriteRequest, 1),
	}
}

// Start begins the WAL background goroutine that processes batches and flushes them periodically.
func (w *WAL) Start(ctx context.Context) {
	go func() {
		ticker := time.NewTicker(w.flushTimeout)
		defer ticker.Stop()

		for {
			select {
			case <-ctx.Done():
				w.flushBatch()
				return
			default:
			}

			select {
			case <-ctx.Done():
				w.flushBatch()
				return
			case batch := <-w.batches:
				w.logsWriter.Write(batch)
				ticker.Reset(w.flushTimeout)
			case <-ticker.C:
				w.flushBatch()
			}
		}
	}()
}

// Recover reads and returns all WAL logs for database state recovery.
func (w *WAL) Recover() ([]log.Log, error) {
	// TODO: need to compact WAL segments
	return w.logsReader.Read()
}

// Set logs a SET operation to WAL and returns a future that completes when the log is written.
func (w *WAL) Set(ctx context.Context, key, value string) concurrency.FutureError {
	return w.push(ctx, models.Query{Command: models.CommandSet, Args: []string{key, value}})
}

// Del logs a DEL operation to WAL and returns a future that completes when the log is written.
func (w *WAL) Del(ctx context.Context, key string) concurrency.FutureError {
	return w.push(ctx, models.Query{Command: models.CommandDel, Args: []string{key}})
}

// push adds a query to the current batch and triggers flush if batch size is reached.
func (w *WAL) push(ctx context.Context, query models.Query) concurrency.FutureError {
	txID := common.GetTxIDFromContext(ctx)
	record := log.NewWriteRequest(txID, query)

	concurrency.WithLock(&w.mutex, func() {
		w.batch = append(w.batch, record)
		if len(w.batch) == w.maxBatchSize {
			w.batches <- w.batch
			w.batch = nil
		}
	})

	return record.FutureResponse()
}

// flushBatch atomically extracts the current batch and writes it to disk.
func (w *WAL) flushBatch() {
	var batch []log.WriteRequest

	concurrency.WithLock(&w.mutex, func() {
		batch = w.batch
		w.batch = nil
	})

	if len(batch) != 0 {
		w.logsWriter.Write(batch)
	}
}
