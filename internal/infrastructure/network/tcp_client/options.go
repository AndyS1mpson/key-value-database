package tcp_client

import "time"

const defaultBufferSize = 4 << 10

// Option is a function type for configuring TCP client settings.
type Option func(*TCPClient)

// WithClientIdleTimeout sets the idle timeout for the client connection.
func WithClientIdleTimeout(timeout time.Duration) Option {
	return func(client *TCPClient) {
		client.idleTimeout = timeout
	}
}

// WithClientBufferSize sets the buffer size for reading server responses.
func WithClientBufferSize(size uint) Option {
	return func(client *TCPClient) {
		client.bufferSize = int(size)
	}
}
