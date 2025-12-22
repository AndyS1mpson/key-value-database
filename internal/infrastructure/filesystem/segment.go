package filesystem

import (
	"fmt"
	"os"
	"time"

	fsutils "github.com/AndyS1mpson/key-value-database/internal/utils/filesystem"
)

const defaultMaxSegmentSize = 10 << 20

var (
	now             = time.Now
	segmentTemplate = "%s/wal_%d.log"
)

// Segment represents a WAL segment file that automatically rotates when size limit is reached.
type Segment struct {
	file      *os.File
	directory string

	segmentSize    int
	maxSegmentSize int
}

// NewSegment creates a new Segment instance with the given directory and maximum size.
// Uses default values if directory is empty or maxSegmentSize is 0.
func NewSegment(directory string, maxSegmentSize int) *Segment {
	if directory == "" {
		directory = defaultWALDataDirectory
	}

	if maxSegmentSize == 0 {
		maxSegmentSize = defaultMaxSegmentSize
	}

	return &Segment{
		directory:      directory,
		maxSegmentSize: maxSegmentSize,
	}
}

// Write appends data to the current segment file, rotating to a new file if size limit is reached.
func (s *Segment) Write(data []byte) error {
	if s.file == nil || s.segmentSize >= s.maxSegmentSize {
		if err := s.rotateSegment(); err != nil {
			return fmt.Errorf("failed to rotate segment file: %w", err)
		}
	}

	writtenBytes, err := fsutils.WriteFile(s.file, data)
	if err != nil {
		return fmt.Errorf("failed to write data to segment file: %w", err)
	}

	s.segmentSize += writtenBytes

	return nil
}

// rotateSegment creates a new segment file with a timestamp-based name and closes the previous one.
func (s *Segment) rotateSegment() error {
	segmentName := fmt.Sprintf(segmentTemplate, s.directory, now().UnixMilli())
	file, err := fsutils.CreateFile(segmentName)
	if err != nil {
		return err
	}

	s.file = file
	s.segmentSize = 0
	return nil
}
