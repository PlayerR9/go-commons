package stack

// Stacker is an interface that defines the behaviors of a stack.
type Stacker[T any] interface {
	// Peek returns the top element of the stack.
	//
	// Returns:
	//   - T: The top element of the stack.
	//   - bool: True if the stack is not empty, false otherwise.
	Peek() (T, bool)

	// Push adds an element to the top of the stack.
	//
	// Parameters:
	//   - elem: The element to add to the stack.
	Push(elem T)

	// Pop removes and returns the top element of the stack.
	//
	// Returns:
	//   - T: The top element of the stack.
	//   - bool: True if the stack is not empty, false otherwise.
	Pop() (T, bool)

	// IsEmpty checks whether the stack is empty.
	//
	// Returns:
	//   - bool: True if the stack is empty, false otherwise.
	IsEmpty() bool

	// Size returns the number of elements in the stack.
	//
	// Returns:
	//   - int: The number of elements in the stack.
	Size() int

	// Reset clears the stack.
	Reset()
}
