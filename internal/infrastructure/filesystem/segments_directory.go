package filesystem

import (
	"fmt"
	"os"

	fsutils "github.com/AndyS1mpson/key-value-database/internal/utils/filesystem"
)

var defaultWALDataDirectory = "./data/kv-db/wal_logs"

// SegmentDirectory provides access to WAL segment files in a directory.
type SegmentDirectory struct {
	directory string
}

// NewSegmentsDirectory creates a new SegmentDirectory instance.
// Uses default directory if the provided directory is empty.
func NewSegmentsDirectory(directory string) *SegmentDirectory {
	if directory == "" {
		directory = defaultWALDataDirectory
	}

	return &SegmentDirectory{
		directory: directory,
	}
}

// ForEach iterates over all segment files in the directory and calls the action function with each file's content.
func (d *SegmentDirectory) ForEach(action func([]byte) error) error {
	files, err := fsutils.ReadDirectoryOrCreate(d.directory)
	if err != nil {
		return err
	}

	for _, file := range files {
		if file.IsDir() {
			continue
		}

		filename := fmt.Sprintf("%s/%s", d.directory, file.Name())
		data, err := os.ReadFile(filename)
		if err != nil {
			return err
		}

		if err := action(data); err != nil {
			return err
		}
	}

	return nil
}
