package filesystem

import (
	"fmt"
	"os"
)

// ReadDirectoryOrCreate reads directory entries, creating the directory if it doesn't exist.
func ReadDirectoryOrCreate(directory string) ([]os.DirEntry, error) {
	files, err := os.ReadDir(directory)
	if err != nil {
		if os.IsNotExist(err) {
			if err := ensureDirectory(directory); err != nil {
				return nil, fmt.Errorf("failed to create directory: %w", err)
			}
			// Retry reading the directory after creation
			files, err = os.ReadDir(directory)
			if err != nil {
				return nil, fmt.Errorf("failed to scan directory with segments: %w", err)
			}
		} else {
			return nil, fmt.Errorf("failed to scan directory with segments: %w", err)
		}
	}

	return files, nil
}

// ensureDirectory creates a directory at the given path, including all parent directories if needed.
// The path should end with the directory name that needs to be created.
func ensureDirectory(path string) error {
	return os.MkdirAll(path, 0755)
}
