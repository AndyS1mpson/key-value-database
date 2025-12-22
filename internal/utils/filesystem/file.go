package filesystem

import "os"

// CreateFile creates a new file in the filesystem with write-only permissions.
func CreateFile(filename string) (*os.File, error) {
	flags := os.O_CREATE | os.O_WRONLY
	file, err := os.OpenFile(filename, flags, 0644)
	if err != nil {
		return nil, err
	}

	return file, err
}

// WriteFile writes data to a file and syncs it to disk.
// Returns the number of bytes written.
func WriteFile(file *os.File, data []byte) (int, error) {
	writtenBytes, err := file.Write(data)
	if err != nil {
		return 0, err
	}

	if err = file.Sync(); err != nil {
		return 0, err
	}

	return writtenBytes, nil
}
