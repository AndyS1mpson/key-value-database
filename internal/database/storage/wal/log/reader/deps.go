package reader

//go:generate mockgen -source=deps.go -destination=./mocks/mock.go

// segmentsDirectory defines the interface for iterating over WAL segment files.
type segmentsDirectory interface {
	ForEach(func([]byte) error) error
}
