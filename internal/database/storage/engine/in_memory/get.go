package in_memory

import (
	"context"
	"fmt"
)

// Get get data from database by key
func (e *Engine) Get(ctx context.Context, key string) (string, bool) {
	value, found := e.data.Get(key)

	e.logger.Debug(fmt.Sprintf("successfull get by key: %s", key))
	return value, found
}
