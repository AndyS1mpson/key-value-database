package tcp_server

import "time"

const defaultBufferSize = 4 << 10

// Option options for configurate tcp server
type Option func(*TCPServer)

// WithServerIdleTimeout set server idle timeout
func WithServerIdleTimeout(timeout time.Duration) Option {
	return func(server *TCPServer) {
		server.idleTimeout = timeout
	}
}

// WithServerBufferSize set server buffer size
func WithServerBufferSize(size uint) Option {
	return func(server *TCPServer) {
		server.bufferSize = int(size)
	}
}

// WithServerMaxConnectionsNumber set server limit of connections
func WithServerMaxConnectionsNumber(count uint) Option {
	return func(server *TCPServer) {
		server.maxConnections = int(count)
	}
}
