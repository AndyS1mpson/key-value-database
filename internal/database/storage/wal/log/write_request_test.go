package log_test

import (
	"errors"
	"testing"

	"github.com/stretchr/testify/assert"

	"github.com/AndyS1mpson/key-value-database/internal/database/compute/models"
	. "github.com/AndyS1mpson/key-value-database/internal/database/storage/wal/log"
)

func TestWriteRequest_WithError(t *testing.T) {
	t.Parallel()

	query := models.Query{
		Command: models.CommandGet,
		Args:    []string{"key"},
	}

	request := NewWriteRequest(100, query)
	future := request.FutureResponse()

	go func() {
		request.SetResponse(errors.New("error"))
	}()

	err := future.Get()
	assert.Error(t, err, "error")
}

func TestWriteRequest(t *testing.T) {
	t.Parallel()

	query := models.Query{
		Command: models.CommandGet,
		Args:    []string{"key"},
	}

	request := NewWriteRequest(100, query)
	future := request.FutureResponse()

	go func() {
		request.SetResponse(nil)
	}()

	err := future.Get()
	assert.NoError(t, err)
}
