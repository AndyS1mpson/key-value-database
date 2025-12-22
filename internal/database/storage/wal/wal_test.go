package wal_test

import (
	"context"
	"testing"
	"time"

	"github.com/stretchr/testify/assert"
	"go.uber.org/mock/gomock"

	"github.com/AndyS1mpson/key-value-database/internal/common"
	. "github.com/AndyS1mpson/key-value-database/internal/database/storage/wal"
	"github.com/AndyS1mpson/key-value-database/internal/database/storage/wal/log"
	mock "github.com/AndyS1mpson/key-value-database/internal/database/storage/wal/mocks"
)

func TestWAL_FlushByTimeout(t *testing.T) {
	t.Parallel()

	ctrl := gomock.NewController(t)
	logsReader := mock.NewMocklogsReader(ctrl)
	logsWriter := mock.NewMocklogsWriter(ctrl)
	logsWriter.EXPECT().
		Write(gomock.Any()).
		Do(func(requests []log.WriteRequest) {
			for _, request := range requests {
				request.SetResponse(nil)
			}
		})

	const timeout = 50 * time.Millisecond
	wal := NewWAL(logsWriter, logsReader, timeout, 1000)

	ctx, cancel := context.WithCancel(t.Context())
	defer cancel()

	wal.Start(ctx)

	future1 := wal.Set(common.ContextWithTxID(t.Context(), 10), "key1", "value1")
	future2 := wal.Set(common.ContextWithTxID(t.Context(), 20), "key2", "value2")

	time.Sleep(timeout)
	assert.NoError(t, future1.Get())
	assert.NoError(t, future2.Get())
}

func TestWAL_FlushBySize(t *testing.T) {
	t.Parallel()

	ctrl := gomock.NewController(t)
	logsReader := mock.NewMocklogsReader(ctrl)
	logsWriter := mock.NewMocklogsWriter(ctrl)

	logsWriter.EXPECT().
		Write(gomock.Any()).
		Do(func(requests []log.WriteRequest) {
			for _, request := range requests {
				request.SetResponse(nil)
			}
		})

	wal := NewWAL(logsWriter, logsReader, time.Minute, 2)

	ctx, cancel := context.WithCancel(t.Context())
	defer cancel()

	wal.Start(ctx)

	future1 := wal.Set(common.ContextWithTxID(t.Context(), 10), "key1", "value1")
	future2 := wal.Set(common.ContextWithTxID(t.Context(), 20), "key2", "value2")

	time.Sleep(100 * time.Millisecond)
	assert.NoError(t, future1.Get())
	assert.NoError(t, future2.Get())
}
