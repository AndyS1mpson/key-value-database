package in_memory_test

import (
	"testing"

	"go.uber.org/mock/gomock"
	"go.uber.org/zap"

	. "github.com/AndyS1mpson/key-value-database/internal/database/storage/engine/in_memory"
)

func TestEngine_Set(t *testing.T) {
	testCases := []struct {
		name  string
		input struct {
			key   string
			value string
		}
	}{
		{
			name: "success set data",
			input: struct {
				key   string
				value string
			}{
				key:   "test_key",
				value: "test_value",
			},
		},
	}

	for _, tc := range testCases {
		ctrl := gomock.NewController(t)
		defer ctrl.Finish()

		engine := NewEngine(zap.NewNop())

		engine.Set(t.Context(), tc.input.key, tc.input.value)
	}
}
