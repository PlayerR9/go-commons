package errors

import (
	"fmt"
	"strconv"
)

// ErrUnexpectedType represents an error when a value has an invalid type.
type ErrUnexpectedType[T any] struct {
	// Elem is the element that caused the error.
	Elem T

	// Kind is the category of the type that was expected.
	Kind string
}

// Error implements the error interface.
//
// Message: "type <type> is not a valid <kind> type"
func (e ErrUnexpectedType[T]) Error() string {
	return fmt.Sprintf("type (%T) is not a valid %s type", e.Elem, e.Kind)
}

// NewErrUnexpectedType creates a new ErrUnexpectedType error.
//
// Parameters:
//   - kind: The name of the type that was expected.
//   - elem: The element that caused the error.
//
// Returns:
//   - *ErrUnexpectedType: A pointer to the newly created ErrUnexpectedType. Never returns nil.
func NewErrUnexpectedType[T any](kind string, elem T) *ErrUnexpectedType[T] {
	e := &ErrUnexpectedType[T]{
		Elem: elem,
		Kind: kind,
	}
	return e
}

// ErrUnexpectedValue represents an error when an unexpected value is encountered.
// This is mostly used in the 'default' case of switch statements.
type ErrUnexpectedValue[T fmt.Stringer] struct {
	// Elem is the element that caused the error.
	Elem T

	// Kind is the category of the type that was expected.
	Kind string
}

// Error implements the error interface.
//
// Message: "<type> (<elem>) is not a supported <kind> type"
func (e ErrUnexpectedValue[T]) Error() string {
	return fmt.Sprintf("%T (%s) is not a supported %s", e.Elem, strconv.Quote(e.Elem.String()), e.Kind)
}

// NewErrUnexpectedValue creates a new ErrUnexpectedValue error.
//
// Parameters:
//   - elem: The element that caused the error.
//   - kind: The name of the type that was expected.
//
// Returns:
//   - *ErrUnexpectedValue: A pointer to the newly created ErrUnexpectedValue. Never returns nil.
func NewErrUnexpectedValue[T fmt.Stringer](elem T, kind string) *ErrUnexpectedValue[T] {
	return &ErrUnexpectedValue[T]{
		Elem: elem,
		Kind: kind,
	}
}

// ErrPanic represents an error when a panic occurs.
type ErrPanic struct {
	// Value is the value that caused the panic.
	Value any
}

// Error implements the error interface.
//
// Message: "panic: {value}"
func (e ErrPanic) Error() string {
	return fmt.Sprintf("panic: %v", e.Value)
}

// NewErrPanic creates a new ErrPanic error.
//
// Parameters:
//   - value: The value that caused the panic.
//
// Returns:
//   - *ErrPanic: A pointer to the newly created ErrPanic. Never returns nil.
func NewErrPanic(value any) *ErrPanic {
	return &ErrPanic{
		Value: value,
	}
}

// AssertionFailed is the message that is shown when an assertion fails.
const AssertionFailed string = "[ASSERTION FAILED]: "

// ErrAssertFailed is the error that is shown when an assertion fails.
type ErrAssertFailed struct {
	// Msg is the message that is shown when the assertion fails.
	Msg string
}

// Error implements the error interface.
//
// Message: "[ASSERTION FAILED]: <msg>"
func (e ErrAssertFailed) Error() string {
	var msg string

	if e.Msg == "" {
		msg = "something went wrong"
	} else {
		msg = e.Msg
	}

	return AssertionFailed + msg
}

// NewErrAssertFailed is a constructor for ErrAssertFailed.
//
// Parameters:
//   - msg: the message that is shown when the assertion fails.
//
// Returns:
//   - *ErrAssertFailed: the error. Never returns nil.
func NewErrAssertFailed(msg string) *ErrAssertFailed {
	return &ErrAssertFailed{
		Msg: msg,
	}
}
