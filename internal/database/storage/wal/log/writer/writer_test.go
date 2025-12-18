package writer_test

import (
	"bytes"
	"errors"
	"testing"

	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/require"
	"go.uber.org/mock/gomock"
	"go.uber.org/zap"

	"github.com/AndyS1mpson/key-value-database/internal/database/compute/models"
	"github.com/AndyS1mpson/key-value-database/internal/database/storage/wal/log"
	. "github.com/AndyS1mpson/key-value-database/internal/database/storage/wal/log/writer"
	mock "github.com/AndyS1mpson/key-value-database/internal/database/storage/wal/log/writer/mocks"
)

func TestWriter_WriteWithErrors(t *testing.T) {
	t.Parallel()

	expectedErr := errors.New("write error")
	requests := []log.WriteRequest{
		log.NewWriteRequest(100, models.Query{Command: models.CommandSet, Args: []string{"key", "value"}}),
		log.NewWriteRequest(200, models.Query{Command: models.CommandGet, Args: []string{"key"}}),
		log.NewWriteRequest(300, models.Query{Command: models.CommandDel, Args: []string{"key"}}),
	}

	var buffer bytes.Buffer
	for idx := range requests {
		log := requests[idx].Log()
		err := log.Encode(&buffer)
		require.NoError(t, err)
	}

	ctrl := gomock.NewController(t)
	segment := mock.NewMocksegment(ctrl)
	segment.EXPECT().
		Write(buffer.Bytes()).
		Return(expectedErr)

	writer := NewWriter(segment, zap.NewNop())

	writer.Write(requests)

	for _, request := range requests {
		futureResponse := request.FutureResponse()
		assert.Equal(t, expectedErr, futureResponse.Get())
	}
}

func TestWriter_Write(t *testing.T) {
	t.Parallel()

	requests := []log.WriteRequest{
		log.NewWriteRequest(100, models.Query{Command: models.CommandSet, Args: []string{"key", "value"}}),
		log.NewWriteRequest(200, models.Query{Command: models.CommandGet, Args: []string{"key"}}),
		log.NewWriteRequest(300, models.Query{Command: models.CommandDel, Args: []string{"key"}}),
	}

	var buffer bytes.Buffer
	for idx := range requests {
		log := requests[idx].Log()
		err := log.Encode(&buffer)
		require.NoError(t, err)
	}

	ctrl := gomock.NewController(t)
	segment := mock.NewMocksegment(ctrl)
	segment.EXPECT().
		Write(buffer.Bytes()).
		Return(nil)

	writer := NewWriter(segment, zap.NewNop())
	writer.Write(requests)

	for _, request := range requests {
		futureResponse := request.FutureResponse()
		assert.Nil(t, futureResponse.Get())
	}
}
