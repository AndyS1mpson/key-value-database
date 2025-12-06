package parser_test

import (
	"errors"
	"testing"

	"github.com/stretchr/testify/assert"
	"go.uber.org/mock/gomock"
	"go.uber.org/zap"

	"github.com/AndyS1mpson/key-value-database/internal/database/compute/models"
	. "github.com/AndyS1mpson/key-value-database/internal/database/compute/parser"
)

func TestQueryParser_Parse(t *testing.T) {
	testCases := []struct {
		name     string
		rawQuery string
		want     *models.Query
		wantErr  error
	}{
		{
			name:     "success parse query",
			rawQuery: "GET test_key",
			want: &models.Query{
				Command: models.CommandGet,
				Args:    []string{"test_key"},
			},
		},
		{
			name:     "empty query string error",
			rawQuery: "",
			wantErr:  errors.New("empty query string"),
		},
		{
			name:     "unknown command error",
			rawQuery: "MSET test_key",
			wantErr:  errors.New("unknown command"),
		},
		{
			name:     "incorrect number of command arguments error",
			rawQuery: "GET test_key test_value",
			wantErr:  errors.New("incorrect number of command arguments"),
		},
	}

	for _, tc := range testCases {
		ctrl := gomock.NewController(t)
		defer ctrl.Finish()

		logger := zap.NewNop()

		parser := New(logger)

		query, err := parser.Parse(tc.rawQuery)
		assert.Equal(t, tc.want, query)
		assert.Equal(t, tc.wantErr, err)
	}
}
