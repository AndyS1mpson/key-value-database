package reader

import (
	"bytes"
	"fmt"
	"sort"

	"github.com/AndyS1mpson/key-value-database/internal/database/storage/wal/log"
)

// Reader reads WAL log entries from segment files in a directory.
type Reader struct {
	segmentsDirectory segmentsDirectory
}

// NewReader creates a new Reader instance with the given segments directory.
func NewReader(segmentsDirectory segmentsDirectory) *Reader {
	return &Reader{
		segmentsDirectory: segmentsDirectory,
	}
}

// Read reads all logs from all segment files and returns them sorted by LSN.
func (r *Reader) Read() ([]log.Log, error) {
	var logs []log.Log

	err := r.segmentsDirectory.ForEach(func(data []byte) error {
		var err error

		logs, err = r.readSegment(logs, data)

		return err
	})
	if err != nil {
		return nil, fmt.Errorf("failed to read segments: %w", err)
	}

	sort.Slice(logs, func(i, j int) bool {
		return logs[i].LSN < logs[j].LSN
	})

	return logs, nil
}

// readSegment decodes all log entries from a segment file's data.
func (r *Reader) readSegment(logs []log.Log, data []byte) ([]log.Log, error) {
	buffer := bytes.NewBuffer(data)
	for buffer.Len() > 0 {
		var log log.Log
		if err := log.Decode(buffer); err != nil {
			return nil, fmt.Errorf("failed to parse logs data: %w", err)
		}

		logs = append(logs, log)
	}

	return logs, nil
}
