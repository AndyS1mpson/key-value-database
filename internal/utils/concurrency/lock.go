package concurrency

import "sync"

// WithLock executes the action function while holding the mutex lock.
func WithLock(mutex sync.Locker, action func()) {
	if action == nil {
		return
	}

	mutex.Lock()
	defer mutex.Unlock()

	action()
}
