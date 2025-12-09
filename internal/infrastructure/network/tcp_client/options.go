package tcp_client

import "time"

const defaultBufferSize = 4 << 10

// Option options for configurate tcp client
type Option func(*TCPClient)

// WithClientIdleTimeout set client idle timeout
func WithClientIdleTimeout(timeout time.Duration) Option {
	return func(client *TCPClient) {
		client.idleTimeout = timeout
	}
}

// WithClientBufferSize set client buffer size
func WithClientBufferSize(size uint) Option {
	return func(client *TCPClient) {
		client.bufferSize = int(size)
	}
}
