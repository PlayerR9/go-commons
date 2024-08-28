package stack

import "slices"

// Stack is a stack of elements.
type Stack[T any] struct {
	// elems is the stack of elements.
	elems []T
}

// Peek implements the Stacker interface.
func (s *Stack[T]) Peek() (T, bool) {
	if len(s.elems) == 0 {
		return *new(T), false
	}

	return s.elems[len(s.elems)-1], true
}

// Push implements the Stacker interface.
func (s *Stack[T]) Push(elem T) {
	s.elems = append(s.elems, elem)
}

// Pop implements the Stacker interface.
func (s *Stack[T]) Pop() (T, bool) {
	if len(s.elems) == 0 {
		return *new(T), false
	}

	elem := s.elems[len(s.elems)-1]
	s.elems = s.elems[:len(s.elems)-1]

	return elem, true
}

// IsEmpty implements the Stacker interface.
func (s *Stack[T]) IsEmpty() bool {
	return len(s.elems) == 0
}

// Size implements the Stacker interface.
func (s *Stack[T]) Size() int {
	return len(s.elems)
}

// Reset implements the Stacker interface.
func (s *Stack[T]) Reset() {
	for i := 0; i < len(s.elems); i++ {
		s.elems[i] = *new(T)
	}

	s.elems = s.elems[:0]
}

// NewStack creates a new stack. The first element of the slice is the top of the stack.
//
// Parameters:
//   - elems: The initial elements of the stack.
//
// Returns:
//   - *Stack[T]: The created stack. Never returns nil.
func NewStack[T any](elems ...T) *Stack[T] {
	slices.Reverse(elems)

	return &Stack[T]{
		elems: elems,
	}
}
