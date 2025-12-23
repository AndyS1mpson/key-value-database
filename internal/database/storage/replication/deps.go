package replication

import "context"

//go:generate mockgen -source=deps.go -destination=./mocks/mock.go

type tcpServer interface {
	HandleQueries(context.Context, func(context.Context, []byte) []byte)
}

type tcpClient interface {
	Send([]byte) ([]byte, error)
	Close()
}
