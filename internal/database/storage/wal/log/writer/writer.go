package writer

import (
	"bytes"

	"github.com/AndyS1mpson/key-value-database/internal/database/storage/wal/log"
	"go.uber.org/zap"
)

// Writer writes WAL log entries to segment files.
type Writer struct {
	segment segment
	logger  *zap.Logger
}

// NewWriter creates a new Writer instance with the given segment and logger.
func NewWriter(segment segment, logger *zap.Logger) *Writer {
	return &Writer{
		segment: segment,
		logger:  logger,
	}
}

// Write encodes and writes a batch of write requests to the segment file.
// Acknowledges all requests with the write result.
func (w *Writer) Write(requests []log.WriteRequest) {
	var buffer bytes.Buffer

	for idx := range requests {
		log := requests[idx].Log()
		if err := log.Encode(&buffer); err != nil {
			w.logger.Warn("failed to encode logs data", zap.Error(err))
			w.acknowledgeWrite(requests, err)

			return
		}
	}

	err := w.segment.Write(buffer.Bytes())
	if err != nil {
		w.logger.Warn("failed to write logs data", zap.Error(err))

	}

	w.acknowledgeWrite(requests, err)
}

// acknowledgeWrite completes all write request promises with the given error.
func (w *Writer) acknowledgeWrite(requests []log.WriteRequest, err error) {
	for idx := range requests {
		requests[idx].SetResponse(err)
	}
}
