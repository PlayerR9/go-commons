package errors

import (
	"reflect"
)

// ErrNilValue represents an error when a pointer is nil.
type ErrNilValue struct{}

// Error implements the error interface.
//
// Message: "pointer must not be nil"
func (e *ErrNilValue) Error() string {
	return "pointer must not be nil"
}

// NewErrNilValue creates a new ErrNilValue error.
//
// Returns:
//   - *ErrNilValue: A pointer to the newly created ErrNilValue. Never returns nil.
func NewErrNilValue() *ErrNilValue {
	return &ErrNilValue{}
}

// ErrEmpty represents an error when a value is empty.
type ErrEmpty struct {
	// Type is the type of the empty value.
	Type any
}

// Error implements the error interface.
//
// Message: "{{ .Type }} must not be empty"
func (e *ErrEmpty) Error() string {
	var t_string string

	if e.Type == nil {
		t_string = "nil"
	} else {
		to := reflect.TypeOf(e.Type)
		t_string = to.String()
	}

	return t_string + " must not be empty"
}

// NewErrEmpty creates a new ErrEmpty error.
//
// Parameters:
//   - var_type: The type of the empty value.
//
// Returns:
//   - *ErrEmpty: A pointer to the newly created ErrEmpty. Never returns nil.
func NewErrEmpty(var_type any) *ErrEmpty {
	return &ErrEmpty{
		Type: var_type,
	}
}
