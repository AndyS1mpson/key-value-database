package concurrency

// Semaphore implementation of synhronized structure semaphore
type Semaphore struct {
	tickets chan struct{}
}

func NewSemaphore(ticketsNumber int) Semaphore {
	return Semaphore{
		tickets: make(chan struct{}, ticketsNumber),
	}
}

// Acquire attempts to acquire a permit, blocking if none are available (like Wait)
func (s *Semaphore) Acquire() {
	if s == nil || s.tickets == nil {
		return
	}

	s.tickets <- struct{}{}
}

// Release releases a permit, increasing the available count (like Signal)
func (s *Semaphore) Release() {
	if s == nil || s.tickets == nil {
		return
	}

	<-s.tickets
}

// WithAcquire  attempts to acquire a permit non-blockingly; returns success/failure immediately
func (s *Semaphore) WithAcquire(action func()) {
	if s == nil || action == nil {
		return
	}

	s.Acquire()
	action()
	s.Release()
}
