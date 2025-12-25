package container

import (
	"fmt"

	"github.com/samber/lo"
	"go.uber.org/zap"

	"github.com/AndyS1mpson/key-value-database/internal/infrastructure/filesystem"
	"github.com/AndyS1mpson/key-value-database/internal/infrastructure/log"
	"github.com/AndyS1mpson/key-value-database/internal/infrastructure/network/tcp_server"
	"github.com/AndyS1mpson/key-value-database/internal/utils/container"
	"github.com/AndyS1mpson/key-value-database/internal/utils/size_parser"
	utils "github.com/AndyS1mpson/key-value-database/internal/utils/size_parser"
)

// GetLogger creates and returns a production logger instance.
func (c *Container) GetLogger() *zap.Logger {
	logger, err := log.ConfigureZapLogger(c.config.Logging)
	if err != nil {
		panic(fmt.Sprintf("configure logger: %s", err))
	}

	return logger
}

// GetTCPServer creates and configures a TCP server instance based on the configuration.
func (c *Container) GetTCPServer() *tcp_server.TCPServer {
	return container.MustOrGetNew(c.Container, func() *tcp_server.TCPServer {

		var options []tcp_server.Option
		options = append(options, tcp_server.WithServerIdleTimeout(lo.FromPtr(c.config.Network.IdleTimeout)))
		options = append(options, tcp_server.WithServerMaxConnectionsNumber(lo.FromPtr(c.config.Network.MaxConnectionsSize)))

		if c.config.Network.MaxMessageSize != nil {
			bufferSize, _ := utils.ParseSize(*c.config.Network.MaxMessageSize)

			options = append(options, tcp_server.WithServerBufferSize(uint(bufferSize)))
		}

		server, err := tcp_server.NewTCPServer(c.config.Network.Address, c.GetLogger(), options...)
		if err != nil {
			panic(fmt.Errorf("failed to init tcp server: %w", err))
		}

		return server
	})
}

// getSegmentsDirectory creates a SegmentDirectory instance for WAL segment access.
func (c *Container) getSegmentsDirectory() *filesystem.SegmentDirectory {
	return container.MustOrGetNew(c.Container, func() *filesystem.SegmentDirectory {
		return filesystem.NewSegmentsDirectory(c.config.WAL.DirectoryPath)
	})
}

// getSegment creates a Segment instance for WAL file writing.
func (c *Container) getSegment() *filesystem.Segment {
	maxSegmentSize, err := size_parser.ParseSize(c.config.WAL.MaxSegmentSize)
	if err != nil {
		c.GetLogger().Error("can not parse wal max_segment_size: %w", zap.Error(err))
	}

	return container.MustOrGetNew(c.Container, func() *filesystem.Segment {
		return filesystem.NewSegment(c.config.WAL.DirectoryPath, maxSegmentSize)
	})
}
