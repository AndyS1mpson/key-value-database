package concurrency

import (
	"sync"
)

// HashTable provides a thread-safe key-value storage using a read-write mutex.
type HashTable struct {
	mutex sync.RWMutex
	data  map[string]string
}

// NewHashTable creates a new thread-safe hash table instance.
func NewHashTable() *HashTable {
	return &HashTable{
		data: make(map[string]string),
	}
}

// Set stores a key-value pair in the hash table.
func (s *HashTable) Set(key, value string) {
	s.mutex.Lock()
	defer s.mutex.Unlock()

	s.data[key] = value
}

// Get retrieves a value by key from the hash table.
// Returns the value and a boolean indicating if the key exists.
func (s *HashTable) Get(key string) (string, bool) {
	s.mutex.RLock()
	defer s.mutex.RUnlock()

	value, found := s.data[key]
	return value, found
}

// Del removes a key-value pair from the hash table.
func (s *HashTable) Del(key string) {
	s.mutex.Lock()
	defer s.mutex.Unlock()

	delete(s.data, key)
}
