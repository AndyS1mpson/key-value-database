package in_memory

import (
	"context"
	"fmt"
)

// Del remove data by ker from database
func (e *Engine) Del(ctx context.Context, key string) {
	delete(e.data, key)

	e.logger.Debug(fmt.Sprintf("successfull del by key: %s", key))
}
