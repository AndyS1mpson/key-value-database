package datastructs

// Stack структура данных - стэк
type Stack[T any] struct {
	s []T
}

// NewStack конструктор для Stack
func NewStack[T any]() *Stack[T] {
	return &Stack[T]{
		s: make([]T, 0),
	}
}

// Push добавление значения в конец стэка
func (s *Stack[T]) Push(val T) {
	s.s = append(s.s, val)
}

// Pop извлечение значения из конца стэка
func (s *Stack[T]) Pop() T {
	val := s.s[len(s.s)-1]

	s.s = s.s[0 : len(s.s)-1]

	return val
}

// IsEmpty проверка стэка на пустоту
func (s *Stack[T]) IsEmpty() bool {
	return len(s.s) == 0
}
