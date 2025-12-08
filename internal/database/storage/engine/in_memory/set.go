package in_memory

import (
	"context"
	"fmt"
)

// Set save data to database
func (e *Engine) Set(ctx context.Context, key, value string) {
	e.data[key] = value

	e.logger.Debug(fmt.Sprintf("successfull set by key: %s", key))
}
