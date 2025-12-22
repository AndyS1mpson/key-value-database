package concurrency

// FutureError is a type alias for Future[error].
type FutureError = Future[error]

// Future represents an asynchronous computation that will produce a value of type T.
type Future[T any] struct {
	result <-chan T
}

// NewFuture creates a new Future from a channel that will receive the result.
func NewFuture[T any](result <-chan T) Future[T] {
	return Future[T]{
		result: result,
	}
}

// Get blocks until the future's result is available and returns it.
func (f *Future[T]) Get() T {
	return <-f.result
}
