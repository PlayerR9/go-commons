package internal

import (
	"slices"
	"strings"
)

// StackTrace represents a stack trace.
type StackTrace struct {
	// Trace is the stack trace.
	trace []string
}

// String implements the fmt.Stringer interface.
func (st StackTrace) String() string {
	elem := make([]string, len(st.trace))
	copy(elem, st.trace)

	slices.Reverse(elem)

	return strings.Join(elem, " -> ")
}

// NewStackTrace creates a new StackTrace.
//
// Parameters:
//   - initial_frame: The initial frame of the stack trace.
//
// Returns:
//   - *StackTrace: A pointer to the new stack trace. Never returns nil
func NewStackTrace(initial_frame string) *StackTrace {
	return &StackTrace{
		trace: []string{initial_frame},
	}
}

// AddFrame prepends the frame into the stack trace. Does nothing
// if the receiver is nil.
//
// Parameters:
//   - frame: The frame to add.
func (st *StackTrace) AddFrame(frame string) {
	if st == nil {
		return
	}

	st.trace = append(st.trace, frame)
}
