package replication

import (
	"bytes"
	"context"
	"encoding/gob"
	"fmt"
	"time"

	wal "github.com/AndyS1mpson/key-value-database/internal/database/storage/wal/log"
	segments "github.com/AndyS1mpson/key-value-database/internal/infrastructure/filesystem"
	"github.com/AndyS1mpson/key-value-database/internal/utils/filesystem"
	"go.uber.org/zap"
)

// Slave represents a slave replica node that periodically synchronizes with a master node.
// It receives WAL segments from the master and applies them to maintain data consistency.
type Slave struct {
	client tcpClient
	stream chan []wal.Log

	syncInterval    time.Duration
	walDirectory    string
	lastSegmentName string

	logger *zap.Logger
}

// NewSlave creates a new Slave instance with the given TCP client, WAL directory, sync interval, and logger.
// It automatically finds the last WAL segment in the directory to track replication progress.
func NewSlave(client tcpClient, walDirectory string, syncInterval time.Duration, logger *zap.Logger) *Slave {
	segmentName, err := segments.SegmentLast(walDirectory)
	if err != nil {
		logger.Error("failed to finc last WAL segment", zap.Error(err))
	}

	return &Slave{
		client: client,
		stream: make(chan []wal.Log),

		walDirectory:    walDirectory,
		syncInterval:    syncInterval,
		lastSegmentName: segmentName,

		logger: logger,
	}
}

// Start begins periodic synchronization with the master node.
// It runs synchronization at the configured interval until the context is cancelled.
func (s *Slave) Start(ctx context.Context) {
	ticker := time.NewTicker(s.syncInterval)
	defer func() {
		ticker.Stop()
		s.client.Close()
	}()

	go func() {
		for {
			select {
			case <-ctx.Done():
				return
			default:
			}

			select {
			case <-ctx.Done():
				return
			case <-ticker.C:
				s.synchronize()
			}
		}
	}()
}

// IsMaster returns false indicating this is a slave replica node.
func (s *Slave) IsMaster() bool {
	return false
}

// ReplicationStream returns a read-only channel that receives replicated WAL logs from the master.
func (s *Slave) ReplicationStream() <-chan []wal.Log {
	return s.stream
}

// synchronize sends a replication request to the master node and processes the response.
// It requests WAL segments newer than the last known segment.
func (s *Slave) synchronize() {
	requestData, err := Encode(&Request{LastSegmentName: s.lastSegmentName})
	if err != nil {
		s.logger.Error("failed to encode replication request", zap.Error(err))
	}

	responseData, err := s.client.Send(requestData)
	if err != nil {
		s.logger.Error("failed to send replication request", zap.Error(err))
	}

	var response Response
	if err = Decode(&response, responseData); err != nil {
		s.logger.Error("failed to decode replication response", zap.Error(err))
	}

	if response.Succeed {
		s.handleResponse(response)
	} else {
		s.logger.Error("failed to apply replication data: master error")
	}
}

// handleResponse processes a successful replication response from the master.
// It saves the WAL segment to disk and writes the logs to the replication stream.
func (s *Slave) handleResponse(response Response) {
	if response.SegmentName == "" {
		s.logger.Debug("no changes from replication")
		return
	}

	if err := s.saveWALSegment(response.SegmentName, response.SegmentData); err != nil {
		s.logger.Error("failed to apply replication data", zap.Error(err))
	}

	if err := s.writeDataToStream(response.SegmentData); err != nil {
		s.logger.Error("failed to write data to stream", zap.Error(err))
	}

	s.lastSegmentName = response.SegmentName
}

// saveWALSegment saves a WAL segment file to disk in the configured WAL directory.
func (s *Slave) saveWALSegment(segmentName string, segmentData []byte) error {
	filename := fmt.Sprintf("%s/%s", s.walDirectory, segmentName)

	file, err := filesystem.CreateFile(filename)
	if err != nil {
		return fmt.Errorf("failed to create wal segment: %w", err)
	}

	_, err = filesystem.WriteFile(file, segmentData)
	return err
}

// writeDataToStream decodes WAL logs from segment data and writes them to the replication stream.
func (s *Slave) writeDataToStream(segmentData []byte) error {
	var logs []wal.Log
	buffer := bytes.NewBuffer(segmentData)
	decoder := gob.NewDecoder(buffer)
	if err := decoder.Decode(&logs); err != nil {
		return fmt.Errorf("failed to decode data: %w", err)
	}

	s.stream <- logs
	return nil
}
