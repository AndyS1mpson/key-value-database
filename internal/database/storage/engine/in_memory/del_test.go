package in_memory_test

import (
	"testing"

	"go.uber.org/mock/gomock"
	"go.uber.org/zap"

	"github.com/AndyS1mpson/key-value-database/internal/common"
	. "github.com/AndyS1mpson/key-value-database/internal/database/storage/engine/in_memory"
)

func TestEngine_Delete(t *testing.T) {
	txID := int64(1)

	testCases := []struct {
		name string
		key  string
	}{
		{
			name: "success delete data",
			key:  "test_key",
		},
	}

	for _, tc := range testCases {
		ctrl := gomock.NewController(t)
		defer ctrl.Finish()

		engine := NewEngine(zap.NewNop(), Config{PartitionsCount: 2})

		ctx := common.ContextWithTxID(t.Context(), txID)

		engine.Del(ctx, tc.key)
	}
}
