package log

import (
	"bytes"
	"encoding/gob"

	"github.com/AndyS1mpson/key-value-database/internal/database/compute/models"
)

// Log represents a single WAL log entry with a Log Sequence Number (LSN) and query.
type Log struct {
	LSN   int64
	Query models.Query
}

// Encode serializes the log entry to the buffer using gob encoding.
func (l *Log) Encode(buffer *bytes.Buffer) error {
	encoder := gob.NewEncoder(buffer)

	return encoder.Encode(*l)
}

// Decode deserializes a log entry from the buffer using gob decoding.
func (l *Log) Decode(buffer *bytes.Buffer) error {
	decoder := gob.NewDecoder(buffer)
	return decoder.Decode(l)
}
