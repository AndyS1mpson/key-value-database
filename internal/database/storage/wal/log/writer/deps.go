package writer

//go:generate mockgen -source=deps.go -destination=./mocks/mock.go

// segment defines the interface for writing data to WAL segment files.
type segment interface {
	Write([]byte) error
}
