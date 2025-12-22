package shutdown

import (
	"context"
	"os/signal"
	"syscall"
)

// WithCancel creates a context that is cancelled when SIGINT or SIGTERM signals are received.
func WithCancel(parent context.Context) (context.Context, context.CancelFunc) {
	return signal.NotifyContext(parent, syscall.SIGINT, syscall.SIGTERM)
}
