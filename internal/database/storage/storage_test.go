package storage_test

import (
	"context"
	"testing"

	"github.com/stretchr/testify/assert"
	"go.uber.org/mock/gomock"

	"github.com/AndyS1mpson/key-value-database/internal/database/storage"
	mock "github.com/AndyS1mpson/key-value-database/internal/database/storage/mocks"
)

func TestStorage_GET(t *testing.T) {
	testCases := []struct {
		name          string
		key           string
		engineGet     func(*mock.Mockengine)
		expectedValue string
		expectedError error
	}{
		{
			name: "success get data",
			key:  "test_key",
			engineGet: func(m *mock.Mockengine) {
				m.EXPECT().Get(gomock.Any(), "test_key").Return("test_value", true)
			},
			expectedValue: "test_value",
			expectedError: nil,
		},
		{
			name: "not found",
			key:  "test_key",
			engineGet: func(m *mock.Mockengine) {
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
			tc.engineGet(mockEngine)

			storageInstance := storage.NewStorage(mockEngine)

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
		engineSet func(*mock.Mockengine)
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
			engineSet: func(m *mock.Mockengine) {
				m.EXPECT().Set(gomock.Any(), "test_key", "test_value")
			},
		},
	}

	for _, tc := range testCases {
		t.Run(tc.name, func(t *testing.T) {
			ctrl := gomock.NewController(t)
			defer ctrl.Finish()

			mockEngine := mock.NewMockengine(ctrl)
			tc.engineSet(mockEngine)

			storageInstance := storage.NewStorage(mockEngine)

			ctx := context.Background()

			err := storageInstance.Set(ctx, tc.input.key, tc.input.value)

			assert.NoError(t, err)
		})
	}
}

func TestStorage_DEL(t *testing.T) {
	testCases := []struct {
		name      string
		key       string
		engineDel func(*mock.Mockengine)
	}{
		{
			name: "success delete data",
			key:  "test_key",
			engineDel: func(m *mock.Mockengine) {
				m.EXPECT().Del(gomock.Any(), "test_key")
			},
		},
	}

	for _, tc := range testCases {
		t.Run(tc.name, func(t *testing.T) {
			ctrl := gomock.NewController(t)
			defer ctrl.Finish()

			mockEngine := mock.NewMockengine(ctrl)
			tc.engineDel(mockEngine)

			storageInstance := storage.NewStorage(mockEngine)

			ctx := context.Background()

			err := storageInstance.Del(ctx, tc.key)

			assert.NoError(t, err)
		})
	}
}
