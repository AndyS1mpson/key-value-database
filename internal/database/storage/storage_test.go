package storage_test

import (
	"context"
	"testing"

	"github.com/stretchr/testify/assert"
	"go.uber.org/mock/gomock"
	"go.uber.org/zap"

	"github.com/AndyS1mpson/key-value-database/internal/database/storage"
	mock "github.com/AndyS1mpson/key-value-database/internal/database/storage/mocks"
)

func TestStorage_GET(t *testing.T) {
	testCases := []struct {
		name          string
		key           string
		expectedValue string
		expectedError error

		engine func(e *mock.Mockengine)
		wal    func(w *mock.Mockwal)
	}{
		{
			name: "success get data",
			key:  "test_key",
			expectedValue: "test_value",
			expectedError: nil,

			engine: func(e *mock.Mockengine) {
				e.EXPECT().Get(gomock.Any(), "test_key").Return("test_value", true)
			},
			wal: func(w *mock.Mockwal) {
				w.EXPECT().Recover().Return()
			},
		},
		{
			name: "not found",
			key:  "test_key",
			engine: func(m *mock.Mockengine) {
				m.EXPECT().Get(gomock.Any(), "test_key").Return("", false)
			},
			expectedValue: "",
			expectedError: storage.ErrorNotFound,
		},
	}

	for _, tc := range testCases {
		t.Run(tc.name, func(t *testing.T) {
			ctrl := gomock.NewController(t)
			defer ctrl.Finish()

			mockEngine := mock.NewMockengine(ctrl)
			tc.engine(mockEngine)

			storageInstance := storage.NewStorage(mockEngine, zap.NewNop())

			ctx := context.Background()

			value, err := storageInstance.Get(ctx, tc.key)

			if tc.expectedError != nil {
				assert.Error(t, err)
				assert.Equal(t, tc.expectedError, err)
			} else {
				assert.NoError(t, err)
			}
			assert.Equal(t, tc.expectedValue, value)
		})
	}
}

func TestStorage_SET(t *testing.T) {
	testCases := []struct {
		name  string
		input struct {
			key   string
			value string
		}
		engine func(*mock.Mockengine)
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
			engine: func(m *mock.Mockengine) {
				m.EXPECT().Set(gomock.Any(), "test_key", "test_value")
			},
		},
	}

	for _, tc := range testCases {
		t.Run(tc.name, func(t *testing.T) {
			ctrl := gomock.NewController(t)
			defer ctrl.Finish()

			mockEngine := mock.NewMockengine(ctrl)
			tc.engine(mockEngine)

			storageInstance := storage.NewStorage(mockEngine, zap.NewNop())

			ctx := context.Background()

			err := storageInstance.Set(ctx, tc.input.key, tc.input.value)

			assert.NoError(t, err)
		})
	}
}

func TestStorage_DEL(t *testing.T) {
	testCases := []struct {
		name   string
		key    string
		engine func(*mock.Mockengine)
	}{
		{
			name: "success delete data",
			key:  "test_key",
			engine: func(m *mock.Mockengine) {
				m.EXPECT().Del(gomock.Any(), "test_key")
			},
		},
	}

	for _, tc := range testCases {
		t.Run(tc.name, func(t *testing.T) {
			ctrl := gomock.NewController(t)
			defer ctrl.Finish()

			mockEngine := mock.NewMockengine(ctrl)
			tc.engine(mockEngine)

			storageInstance := storage.NewStorage(mockEngine, zap.NewNop())

			ctx := context.Background()

			err := storageInstance.Del(ctx, tc.key)

			assert.NoError(t, err)
		})
	}
}
