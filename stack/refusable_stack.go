package stack

import "slices"

// RefusableStack is a stack that can be refused.
type RefusableStack[T any] struct {
	// elems is the stack of elements.
	elems []T

	// popped is the elements that have been popped from the stack.
	popped []T
}

// Peek implements the Stacker interface.
func (s *RefusableStack[T]) Peek() (T, bool) {
	if len(s.elems) == 0 {
		return *new(T), false
	}

	return s.elems[len(s.elems)-1], true
}

// Push implements the Stacker interface.
func (s *RefusableStack[T]) Push(elem T) {
	s.elems = append(s.elems, elem)
}

// Pop implements the Stacker interface.
func (s *RefusableStack[T]) Pop() (T, bool) {
	if len(s.elems) == 0 {
		return *new(T), false
	}

	top := s.elems[len(s.elems)-1]
	s.elems = s.elems[:len(s.elems)-1]

	s.popped = append(s.popped, top)

	return top, true
}

// IsEmpty implements the Stacker interface.
func (s *RefusableStack[T]) IsEmpty() bool {
	return len(s.elems) == 0
}

// Size implements the Stacker interface.
func (s *RefusableStack[T]) Size() int {
	return len(s.elems)
}

// Reset implements the Stacker interface.
func (s *RefusableStack[T]) Reset() {
	for i := 0; i < len(s.elems); i++ {
		s.elems[i] = *new(T)
	}

	s.elems = s.elems[:0]
}

// NewRefusableStack creates a new refusable stack.
//
// Returns:
//   - *RefusableStack: The new refusable stack. Never returns nil.
func NewRefusableStack[T any]() *RefusableStack[T] {
	return &RefusableStack[T]{}
}

// Popped returns the elements that have been popped from the stack.
//
// Returns:
//   - []T: The popped elements.
//
// The first element of the returned slice is the top of the stack.
func (s *RefusableStack[T]) Popped() []T {
	popped := make([]T, len(s.popped))
	copy(popped, s.popped)

	slices.Reverse(popped)

	return popped
}

// Accept accepts the current state. Use it when you know you don't need any previous state.
func (s *RefusableStack[T]) Accept() {
	s.popped = nil
}

// Refuse allows to returned to the state before the last accept call.
func (s *RefusableStack[T]) Refuse() {
	for len(s.popped) > 0 {
		top := s.popped[len(s.popped)-1]
		s.popped = s.popped[:len(s.popped)-1]

		s.elems = append(s.elems, top)
	}
}
