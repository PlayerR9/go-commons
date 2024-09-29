package history

// Subjecter is an interface for a subject.
type Subjecter[E any] interface {
	// IsNil checks if the subject is nil.
	//
	// Returns:
	//   - bool: True if the subject is nil, false otherwise.
	IsNil() bool

	// HasError checks if the subject has an error.
	//
	// Returns:
	//   - bool: True if the subject has an error, false otherwise.
	HasError() bool

	// ApplyEvent applies an event to the subject.
	//
	// Parameters:
	//   - event: The event to apply. Assumed not nil.
	//
	// Returns:
	//   - bool: True if the subject is done, false otherwise.
	ApplyEvent(event E) bool

	// NextEvents returns the next events in the subject.
	//
	// Returns:
	//   - []E: The next events in the subject.
	NextEvents() []E
}
