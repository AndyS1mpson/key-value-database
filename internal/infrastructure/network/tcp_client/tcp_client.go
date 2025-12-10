package tcp_client

import (
	"errors"
	"fmt"
	"io"
	"net"
	"time"
)

var ErrSmallBufferSize = errors.New("small buffer size")

// TCPClient implementation of tcp client
type TCPClient struct {
	connection  net.Conn
	idleTimeout time.Duration
	bufferSize  int
}

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

// Send sending bytes over the network
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

// Close close active connection
func (c *TCPClient) Close() {
	if c.connection != nil {
		_ = c.connection.Close()
	}
}
