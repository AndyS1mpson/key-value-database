package container

import (
	"github.com/AndyS1mpson/key-value-database/internal/database/storage/wal"
	"github.com/AndyS1mpson/key-value-database/internal/database/storage/wal/log/reader"
	"github.com/AndyS1mpson/key-value-database/internal/database/storage/wal/log/writer"
	"github.com/AndyS1mpson/key-value-database/internal/utils/container"
)

// getLogReader creates a Reader instance for reading WAL logs.
func (c *Container) getLogReader() *reader.Reader {
	return container.MustOrGetNew(c.Container, func() *reader.Reader {
		return reader.NewReader(c.getSegmentsDirectory())
	})
}

// getLogWriter creates a Writer instance for writing WAL logs.
func (c *Container) getLogWriter() *writer.Writer {
	return container.MustOrGetNew(c.Container, func() *writer.Writer {
		return writer.NewWriter(c.getSegment(), c.GetLogger())
	})
}

// getWAL creates a WAL instance and starts its background processing goroutine.
func (c *Container) getWAL() *wal.WAL {
	return container.MustOrGetNew(c.Container, func() *wal.WAL {
		wal := wal.NewWAL(c.getLogWriter(), c.getLogReader(), c.config.WAL.FlushingBatchTimeout, c.config.WAL.FlushingBatchSize)

		wal.Start(c.Ctx())

		return wal
	})
}
