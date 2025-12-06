package shutdown

import (
	"context"
	"errors"
	"os"
	"testing"
	"time"
)

func TestWithContext(t *testing.T) {
	proc, err := os.FindProcess(os.Getpid())
	if err != nil {
		t.Fatal(err)
	}

	ctx, cancel := WithCancel(context.Background())
	defer cancel()

	go func() {
		_ = proc.Signal(os.Interrupt)
	}()

	<-ctx.Done()

	if !errors.Is(ctx.Err(), context.Canceled) {
		t.Errorf("excepted canceled, actual: %v", ctx.Err())
	}
}

func TestWithContextBlocked(t *testing.T) {
	ctx, cancel := WithCancel(context.Background())
	defer cancel()

	select {
	case <-ctx.Done():
		t.Errorf("Done() was not excepted")
	case <-time.After(10 * time.Millisecond):
		// excepted
	}
}
