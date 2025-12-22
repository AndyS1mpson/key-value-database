package log

import (
	"github.com/AndyS1mpson/key-value-database/internal/database/compute/models"
	"github.com/AndyS1mpson/key-value-database/internal/utils/concurrency"
)

// WriteRequest represents a WAL write operation with an associated promise for async completion.
type WriteRequest struct {
	log     Log
	promise concurrency.PromiseError
}

// NewWriteRequest creates a new write request with the given LSN and query.
func NewWriteRequest(lsn int64, query models.Query) WriteRequest {
	return WriteRequest{
		log: Log{
			LSN:   lsn,
			Query: query,
		},
		promise: concurrency.NewPromise[error](),
	}
}

// Log returns the log entry associated with this write request.
func (l *WriteRequest) Log() Log {
	return l.log
}

// SetResponse completes the promise with the write operation result.
func (l *WriteRequest) SetResponse(err error) {
	l.promise.Set(err)
}

// FutureResponse returns a future that will be completed when the write operation finishes.
func (l *WriteRequest) FutureResponse() concurrency.FutureError {
	return l.promise.GetFuture()
}
