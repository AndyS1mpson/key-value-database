package database_test

import (
	"errors"
	"testing"

	"github.com/stretchr/testify/assert"
	"go.uber.org/mock/gomock"
	"go.uber.org/zap"

	. "github.com/AndyS1mpson/key-value-database/internal/database"
	"github.com/AndyS1mpson/key-value-database/internal/database/compute/models"
	mock "github.com/AndyS1mpson/key-value-database/internal/database/mocks"
	"github.com/AndyS1mpson/key-value-database/internal/database/storage"
)

func TestDatabase_HandleQuery(t *testing.T) {
	t.Parallel()

	tests := []struct {
		name             string
		query            string
		expectedResponse string

		computeLayer func(c *mock.MockcomputeLayer)
		storageLayer func(s *mock.MockstorageLayer)
	}{
		{
			name:  "handle incorrect query",
			query: "TRUNCATE",
			computeLayer: func(c *mock.MockcomputeLayer) {
				c.EXPECT().
					Parse("TRUNCATE").
					Return(&models.Query{}, errors.New("compute error"))
			},
			storageLayer:     func(s *mock.MockstorageLayer) {},
			expectedResponse: "[error] compute error",
		},
		{
			name:  "handle set query with error from storage",
			query: "SET key value",
			computeLayer: func(c *mock.MockcomputeLayer) {
				c.EXPECT().
					Parse("SET key value").
					Return(models.NewQuery(
						models.CommandSet,
						[]string{"key", "value"},
					), nil)
			},
			storageLayer: func(s *mock.MockstorageLayer) {
				s.EXPECT().
					Set(gomock.Any(), "key", "value").
					Return(errors.New("storage error"))
			},
			expectedResponse: "[error] storage error",
		},
		{
			name:  "handle set query",
			query: "SET key value",
			computeLayer: func(c *mock.MockcomputeLayer) {
				c.EXPECT().
					Parse("SET key value").
					Return(models.NewQuery(
						models.CommandSet,
						[]string{"key", "value"},
					), nil)
			},
			storageLayer: func(s *mock.MockstorageLayer) {
				s.EXPECT().
					Set(gomock.Any(), "key", "value").
					Return(nil)
			},
			expectedResponse: "[ok]",
		},
		{
			name:  "handle del query with error from storage",
			query: "DEL key",
			computeLayer: func(c *mock.MockcomputeLayer) {
				c.EXPECT().
					Parse("DEL key").
					Return(models.NewQuery(
						models.CommandDel,
						[]string{"key"},
					), nil)
			},
			storageLayer: func(s *mock.MockstorageLayer) {
				s.EXPECT().
					Del(gomock.Any(), "key").
					Return(errors.New("storage error"))
			},
			expectedResponse: "[error] storage error",
		},
		{
			name:  "handle del query",
			query: "DEL key",
			computeLayer: func(c *mock.MockcomputeLayer) {
				c.EXPECT().
					Parse("DEL key").
					Return(models.NewQuery(
						models.CommandDel,
						[]string{"key"},
					), nil)
			},
			storageLayer: func(s *mock.MockstorageLayer) {
				s.EXPECT().
					Del(gomock.Any(), "key").
					Return(nil)
			},
			expectedResponse: "[ok]",
		},
		{
			name:  "handle get query with error from storage",
			query: "GET key",
			computeLayer: func(c *mock.MockcomputeLayer) {
				c.EXPECT().
					Parse("GET key").
					Return(models.NewQuery(
						models.CommandGet,
						[]string{"key"},
					), nil)
			},
			storageLayer: func(s *mock.MockstorageLayer) {
				s.EXPECT().
					Get(gomock.Any(), "key").
					Return("", errors.New("storage error"))
			},
			expectedResponse: "[error] storage error",
		},
		{
			name:  "handle get query with not found error from storage",
			query: "GET key",
			computeLayer: func(c *mock.MockcomputeLayer) {
				c.EXPECT().
					Parse("GET key").
					Return(models.NewQuery(
						models.CommandGet,
						[]string{"key"},
					), nil)
			},
			storageLayer: func(s *mock.MockstorageLayer) {
				s.EXPECT().
					Get(gomock.Any(), "key").
					Return("", storage.ErrorNotFound)
			},
			expectedResponse: "[not found]",
		},
		{
			name:  "handle get query",
			query: "GET key",
			computeLayer: func(c *mock.MockcomputeLayer) {
				c.EXPECT().
					Parse("GET key").
					Return(models.NewQuery(
						models.CommandGet,
						[]string{"key"},
					), nil)
			},
			storageLayer: func(s *mock.MockstorageLayer) {
				s.EXPECT().
					Get(gomock.Any(), "key").
					Return("value", nil)
			},
			expectedResponse: "[ok] value",
		},
	}

	for _, tc := range tests {
		t.Run(tc.name, func(t *testing.T) {
			t.Parallel()

			ctrl := gomock.NewController(t)

			computeLayer := mock.NewMockcomputeLayer(ctrl)
			tc.computeLayer(computeLayer)

			storageLayer := mock.NewMockstorageLayer(ctrl)
			tc.storageLayer(storageLayer)

			storage := New(computeLayer, storageLayer, zap.NewNop())

			response := storage.HandleQuery(t.Context(), tc.query)

			assert.Equal(t, tc.expectedResponse, response)
		})
	}
}
