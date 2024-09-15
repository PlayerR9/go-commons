package ints

import (
	"strconv"
	"strings"
)

// ErrGT represents an error when a value is less than or equal to a specified value.
type ErrGT struct {
	// Value is the value that caused the error.
	Value int
}

// Error implements the error interface.
//
// Message: "value must be greater than <value>"
//
// If the value is 0, the message is "value must be positive".
func (e ErrGT) Error() string {
	if e.Value == 0 {
		return "value must be positive"
	}

	value := strconv.Itoa(e.Value)

	var builder strings.Builder

	builder.WriteString("value must be greater than ")
	builder.WriteString(value)

	str := builder.String()

	return str
}

// NewErrGT creates a new ErrGT error with the specified value.
//
// Parameters:
//   - value: The minimum value that is not allowed.
//
// Returns:
//   - *ErrGT: A pointer to the newly created ErrGT.
func NewErrGT(value int) *ErrGT {
	e := &ErrGT{
		Value: value,
	}
	return e
}

// ErrLT represents an error when a value is greater than or equal to a specified value.
type ErrLT struct {
	// Value is the value that caused the error.
	Value int
}

// Error implements the error interface.
//
// Message: "value must be less than <value>"
//
// If the value is 0, the message is "value must be negative".
func (e ErrLT) Error() string {
	if e.Value == 0 {
		return "value must be negative"
	}

	value := strconv.Itoa(e.Value)

	var builder strings.Builder

	builder.WriteString("value must be less than ")
	builder.WriteString(value)

	str := builder.String()
	return str
}

// NewErrLT creates a new ErrLT error with the specified value.
//
// Parameters:
//   - value: The maximum value that is not allowed.
//
// Returns:
//   - *ErrLT: A pointer to the newly created ErrLT.
func NewErrLT(value int) *ErrLT {
	e := &ErrLT{
		Value: value,
	}
	return e
}

// ErrGTE represents an error when a value is less than a specified value.
type ErrGTE struct {
	// Value is the value that caused the error.
	Value int
}

// Error implements the error interface.
//
// Message: "value must be greater than or equal to <value>"
//
// If the value is 0, the message is "value must be non-negative".
func (e ErrGTE) Error() string {
	if e.Value == 0 {
		return "value must be non-negative"
	}

	value := strconv.Itoa(e.Value)

	var builder strings.Builder

	builder.WriteString("value must be greater than or equal to ")
	builder.WriteString(value)

	str := builder.String()
	return str
}

// NewErrGTE creates a new ErrGTE error with the specified value.
//
// Parameters:
//   - value: The minimum value that is allowed.
//
// Returns:
//   - *ErrGTE: A pointer to the newly created ErrGTE.
func NewErrGTE(value int) *ErrGTE {
	e := &ErrGTE{
		Value: value,
	}
	return e
}

// ErrLTE represents an error when a value is greater than a specified value.
type ErrLTE struct {
	// Value is the value that caused the error.
	Value int
}

// Error implements the error interface.
//
// Message: "value must be less than or equal to <value>"
//
// If the value is 0, the message is "value must be non-positive".
func (e ErrLTE) Error() string {
	if e.Value == 0 {
		return "value must be non-positive"
	}

	value := strconv.Itoa(e.Value)

	var builder strings.Builder

	builder.WriteString("value must be less than or equal to ")
	builder.WriteString(value)

	str := builder.String()
	return str
}

// NewErrLTE creates a new ErrLTE error with the specified value.
//
// Parameters:
//   - value: The maximum value that is allowed.
//
// Returns:
//   - *ErrLTE: A pointer to the newly created ErrLTE.
func NewErrLTE(value int) *ErrLTE {
	e := &ErrLTE{
		Value: value,
	}
	return e
}
