package container

import (
	"fmt"

	"github.com/samber/lo"
	"go.uber.org/zap"

	"github.com/AndyS1mpson/key-value-database/internal/infrastructure/network/tcp_server"
	"github.com/AndyS1mpson/key-value-database/internal/utils/container"
	utils "github.com/AndyS1mpson/key-value-database/internal/utils/size_parser"
)

// GetLogger configure logger
func (c *Container) GetLogger() *zap.Logger {
	logger, err := zap.NewProduction()
	if err != nil {
		panic(fmt.Sprintf("can not init logger: %s", err))
	}

	return logger
}

// GetTCPServer configure tcp server
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
