package concurrency

// Semaphore provides a counting semaphore for limiting concurrent access to resources.
type Semaphore struct {
	tickets chan struct{}
}

// NewSemaphore creates a new semaphore with the given number of available tickets.
func NewSemaphore(ticketsNumber int) Semaphore {
	return Semaphore{
		tickets: make(chan struct{}, ticketsNumber),
	}
}

// Acquire acquires a ticket from the semaphore, blocking if none are available.
func (s *Semaphore) Acquire() {
	if s == nil || s.tickets == nil {
		return
	}

	s.tickets <- struct{}{}
}

// Release returns a ticket to the semaphore.
func (s *Semaphore) Release() {
	if s == nil || s.tickets == nil {
		return
	}

	<-s.tickets
}

// WithAcquire executes the action function while holding a semaphore ticket.
func (s *Semaphore) WithAcquire(action func()) {
	if s == nil || action == nil {
		return
	}

	s.Acquire()
	action()
	s.Release()
}
