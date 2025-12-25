package replication

import (
	"context"
	"fmt"
	"os"

	"github.com/AndyS1mpson/key-value-database/internal/infrastructure/filesystem"
	"go.uber.org/zap"
)

// Master represents a master replica node that handles replication requests from slave nodes.
// It serves WAL segments to slaves for synchronization.
type Master struct {
	server       tcpServer
	walDirectory string

	logger *zap.Logger
}

// NewMaster creates a new Master instance with the given TCP server, WAL directory, and logger.
func NewMaster(server tcpServer, walDirectory string, logger *zap.Logger) *Master {
	return &Master{
		server:       server,
		walDirectory: walDirectory,
		logger:       logger,
	}
}

// Start begins handling replication requests from slave nodes.
// It processes incoming requests and responds with WAL segment data.
func (m *Master) Start(ctx context.Context) {
	m.server.HandleQueries(ctx, func(ctx context.Context, requestData []byte) []byte {
		if ctx.Err() != nil {
			return nil
		}

		var request Request
		if err := Decode(&request, requestData); err != nil {
			m.logger.Error("failed to decode replication request", zap.Error(err))

			return nil
		}

		response := m.synchronize(request)
		responseData, err := Encode(&response)
		if err != nil {
			m.logger.Error("failed to encode replication response", zap.Error(err))
		}

		return responseData
	})
}

// IsMaster returns true indicating this is a master replica node.
func (m *Master) IsMaster() bool {
	return true
}

// synchronize processes a replication request from a slave node.
// It finds the next WAL segment after the last segment name provided in the request
// and returns it along with its data if available.
func (m *Master) synchronize(request Request) Response {
	var response Response
	segmentName, err := filesystem.SegmentNext(m.walDirectory, request.LastSegmentName)
	if err != nil {
		m.logger.Error("failed to find WAL segment", zap.Error(err))

		return response
	}

	if segmentName == "" {
		response.Succeed = true
		return response
	}

	filename := fmt.Sprintf("%s/%s", m.walDirectory, segmentName)

	data, err := os.ReadFile(filename)
	if err != nil {
		m.logger.Error("failed to read WAL segment", zap.Error(err))
		return response
	}

	response.Succeed = true
	response.SegmentData = data
	response.SegmentName = filename

	return response
}
