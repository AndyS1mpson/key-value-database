package datastructs

// Stack provides a generic LIFO (Last In First Out) data structure.
type Stack[T any] struct {
	s []T
}

// NewStack creates a new empty stack instance.
func NewStack[T any]() *Stack[T] {
	return &Stack[T]{
		s: make([]T, 0),
	}
}

// Push adds a value to the top of the stack.
func (s *Stack[T]) Push(val T) {
	s.s = append(s.s, val)
}

// Pop removes and returns the top value from the stack.
func (s *Stack[T]) Pop() T {
	val := s.s[len(s.s)-1]

	s.s = s.s[0 : len(s.s)-1]

	return val
}

// IsEmpty checks if the stack is empty.
func (s *Stack[T]) IsEmpty() bool {
	return len(s.s) == 0
}
