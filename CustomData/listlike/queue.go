package listlike

import (
	"slices"
)

// Stack is a stack of events that can be replayed.
type Stack[T any] struct {
	// elems is the stack of events.
	elems []T
}

// IsNil checks if the stack is nil.
//
// Returns:
//   - bool: True if the stack is nil, false otherwise.
func (s *Stack[T]) IsNil() bool {
	return s == nil
}

// NewStack creates a new stack.
//
// Returns:
//   - *Stack: The new stack. Never returns nil.
func NewStack[T any]() *Stack[T] {
	return &Stack[T]{
		elems: make([]T, 0),
	}
}

// Push adds an element to the stack. Doesn't add nil elements or when the
// receiver is nil.
//
// Parameters:
//   - elem: The element to add to the stack.
//
// Returns:
//   - bool: True if the element was added to the stack, false otherwise.
func (s *Stack[T]) Push(elem T) bool {
	if s == nil {
		return false
	}

	s.elems = append(s.elems, elem)

	return true
}

// PushMany adds multiple elements to the stack. Doesn't add nil elements or
// when the receiver is nil.
//
// Parameters:
//   - elems: The elements to add to the stack.
//
// Returns:
//   - bool: True if the element was added to the stack, false otherwise.
func (s *Stack[T]) PushMany(elems []T) bool {
	if s == nil {
		return false
	}

	slices.Reverse(elems)
	s.elems = append(s.elems, elems...)

	return true
}

// Pop removes an element from the stack. Doesn't remove nil elements or when
// the receiver is nil.
//
// Returns:
//   - T: The removed element.
//   - bool: True if the element was removed from the stack, false otherwise.
func (s *Stack[T]) Pop() (T, bool) {
	if s == nil || len(s.elems) == 0 {
		return *new(T), false
	}

	elem := s.elems[len(s.elems)-1]
	s.elems = s.elems[:len(s.elems)-1]

	return elem, true
}
