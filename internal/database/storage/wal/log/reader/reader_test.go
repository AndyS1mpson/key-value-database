package reader_test

import (
	"errors"
	"testing"

	"github.com/stretchr/testify/assert"
	"go.uber.org/mock/gomock"

	. "github.com/AndyS1mpson/key-value-database/internal/database/storage/wal/log/reader"
	mock "github.com/AndyS1mpson/key-value-database/internal/database/storage/wal/log/reader/mocks"
)

func TestReader_ReadWithError(t *testing.T) {
	t.Parallel()

	expectedErr := errors.New("read error")

	ctrl := gomock.NewController(t)
	directory := mock.NewMocksegmentsDirectory(ctrl)
	directory.EXPECT().
		ForEach(gomock.Any()).
		Return(expectedErr)

	reader := NewReader(directory)

	logs, err := reader.Read()
	assert.True(t, errors.Is(err, expectedErr))
	assert.Nil(t, logs)
}

func TestReader_Read(t *testing.T) {
	t.Parallel()

	ctrl := gomock.NewController(t)
	directory := mock.NewMocksegmentsDirectory(ctrl)
	directory.EXPECT().
		ForEach(gomock.Any()).
		Return(nil)

	reader := NewReader(directory)

	logs, err := reader.Read()
	assert.Nil(t, err)
	assert.Nil(t, logs)
}
