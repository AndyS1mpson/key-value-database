package in_memory_test

import (
	"testing"

	"github.com/stretchr/testify/assert"
	"go.uber.org/mock/gomock"
	"go.uber.org/zap"

	. "github.com/AndyS1mpson/key-value-database/internal/database/storage/engine/in_memory"
)

func TestEngine_Get(t *testing.T) {
	testCases := []struct {
		name string
		key  string
	}{
		{
			name: "success get data",
			key:  "test_key",
		},
	}

	for _, tc := range testCases {
		ctrl := gomock.NewController(t)
		defer ctrl.Finish()

		engine := NewEngine(zap.NewNop())


		value, found := engine.Get(t.Context(), tc.key)
		assert.False(t, found)
		assert.Empty(t, value)
	}
}
