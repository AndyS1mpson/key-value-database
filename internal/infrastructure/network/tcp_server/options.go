package tcp_server

import "time"

const defaultBufferSize = 4 << 10

// Option is a function type for configuring TCP server settings.
type Option func(*TCPServer)

// WithServerIdleTimeout sets the idle timeout for client connections.
func WithServerIdleTimeout(timeout time.Duration) Option {
	return func(server *TCPServer) {
		server.idleTimeout = timeout
	}
}

// WithServerBufferSize sets the buffer size for reading client requests.
func WithServerBufferSize(size uint) Option {
	return func(server *TCPServer) {
		server.bufferSize = int(size)
	}
}

// WithServerMaxConnectionsNumber sets the maximum number of concurrent connections.
func WithServerMaxConnectionsNumber(count uint) Option {
	return func(server *TCPServer) {
		server.maxConnections = int(count)
	}
}
