package shutdown

import (
	"context"
	"os/signal"
	"syscall"
)

// WithCancel return context and canceling function
func WithCancel(parent context.Context) (context.Context, context.CancelFunc) {
	return signal.NotifyContext(parent, syscall.SIGINT, syscall.SIGTERM)
}
