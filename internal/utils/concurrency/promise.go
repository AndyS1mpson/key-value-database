package concurrency

// PromiseError is a type alias for Promise[error].
type PromiseError = Promise[error]

// Promise represents a value that will be set asynchronously and can be retrieved via a Future.
type Promise[T any] struct {
	result   chan T
	promised bool
}

// NewPromise creates a new Promise instance.
func NewPromise[T any]() Promise[T] {
	return Promise[T]{
		result: make(chan T, 1),
	}
}

// Set completes the promise with the given value. Should not be called concurrently.
func (p *Promise[T]) Set(value T) {
	if p.promised {
		return
	}

	p.promised = true
	p.result <- value
	close(p.result)
}

// GetFuture returns a Future that will receive the promise's value when Set is called.
func (p *Promise[T]) GetFuture() Future[T] {
	return NewFuture(p.result)
}
