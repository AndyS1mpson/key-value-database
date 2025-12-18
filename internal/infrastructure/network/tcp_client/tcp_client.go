package tcp_client

import (
	"errors"
	"fmt"
	"io"
	"net"
	"time"
)

var ErrSmallBufferSize = errors.New("small buffer size")

// TCPClient provides a TCP client for communicating with the database server.
type TCPClient struct {
	connection  net.Conn
	idleTimeout time.Duration
	bufferSize  int
}

// NewTCPClient creates a new TCP client connected to the given address.
// Configurable via options for timeout and buffer size.
func NewTCPClient(address string, options ...Option) (*TCPClient, error) {
	connection, err := net.Dial("tcp", address)
	if err != nil {
		return nil, fmt.Errorf("failed to dial: %w", err)
	}

	client := &TCPClient{
		connection: connection,
		bufferSize: defaultBufferSize,
	}

	for _, option := range options {
		option(client)
	}

	if client.idleTimeout != 0 {
		if err := connection.SetDeadline(time.Now().Add(client.idleTimeout)); err != nil {
			return nil, fmt.Errorf("failed to set deadline for connection: %w", err)
		}
	}

	return client, nil
}

// Send sends a request to the server and waits for a response.
// Returns an error if the buffer size is insufficient or network error occurs.
func (c *TCPClient) Send(request []byte) ([]byte, error) {
	if _, err := c.connection.Write(request); err != nil {
		return nil, err
	}

	response := make([]byte, c.bufferSize)

	count, err := c.connection.Read(response)
	if err != nil && err != io.EOF {
		return nil, err
	} else if count == c.bufferSize {
		return nil, ErrSmallBufferSize
	}

	return response[:count], nil
}

// Close closes the active TCP connection.
func (c *TCPClient) Close() {
	if c.connection != nil {
		_ = c.connection.Close()
	}
}
