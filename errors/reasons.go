package errors

import (
	"cmp"
	"errors"
	"fmt"
	"reflect"
)

var (
	// NilValue is an error that is returned when a value is nil.
	NilValue error

	// NilReceiver is an error that is returned when a receiver is nil.
	NilReceiver error
)

func init() {
	NilValue = errors.New("value must not be nil")

	NilReceiver = errors.New("receiver must not be nil")
}

// ErrEmpty represents an error when a value is empty.
type ErrEmpty struct {
	// Type is the type of the empty value.
	Type any
}

// Error implements the error interface.
//
// Message: "{{ .Type }} must not be empty"
func (e ErrEmpty) Error() string {
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

// ErrGT represents an error when a value is less than or equal to a specified value.
type ErrGT[T cmp.Ordered] struct {
	// Value is the value that caused the error.
	Value T
}

// Error implements the error interface.
//
// Message: "value must be greater than <value>"
func (e ErrGT[T]) Error() string {
	return fmt.Sprintf("value must ge greater than %v", e.Value)
}

// NewErrGT creates a new ErrGT error with the specified value.
//
// Parameters:
//   - value: The minimum value that is not allowed.
//
// Returns:
//   - *ErrGT: A pointer to the newly created ErrGT. Never returns nil.
func NewErrGT[T cmp.Ordered](value T) *ErrGT[T] {
	return &ErrGT[T]{
		Value: value,
	}
}

// ErrGTE represents an error when a value is less than a specified value.
type ErrGTE[T cmp.Ordered] struct {
	// Value is the value that caused the error.
	Value T
}

// Error implements the error interface.
//
// Message: "value must be greater than or equal to <value>"
func (e ErrGTE[T]) Error() string {
	return fmt.Sprintf("value must be greater than or equal to %v", e.Value)
}

// NewErrGTE creates a new ErrGTE error with the specified value.
//
// Parameters:
//   - value: The minimum value that is allowed.
//
// Returns:
//   - *ErrGTE: A pointer to the newly created ErrGTE. Never returns nil.
func NewErrGTE[T cmp.Ordered](value T) *ErrGTE[T] {
	return &ErrGTE[T]{
		Value: value,
	}
}
