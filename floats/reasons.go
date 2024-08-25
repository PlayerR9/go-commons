package floats

import (
	"strconv"
	"strings"
)

// ErrOutOfBounds represents an error when a value is out of bounds.
type ErrOutOfBounds struct {
	// LowerBound is the lower bound of the value.
	LowerBound float64

	// UpperBound is the upper bound of the value.
	UpperBound float64

	// LowerInclusive is true if the lower bound is inclusive.
	LowerInclusive bool

	// UpperInclusive is true if the upper bound is inclusive.
	UpperInclusive bool

	// Value is the value that is out of bounds.
	Value float64
}

// Error implements the error float64erface.
//
// Message: "value ( <value> ) not in range <lower_bound> , <upper_bound>"
func (e *ErrOutOfBounds) Error() string {
	left_bound := strconv.FormatFloat(e.LowerBound, 'f', -1, 64)
	right_bound := strconv.FormatFloat(e.UpperBound, 'f', -1, 64)

	var open, close string

	if e.LowerInclusive {
		open = "[ "
	} else {
		open = "( "
	}

	if e.UpperInclusive {
		close = " ]"
	} else {
		close = " )"
	}

	var builder strings.Builder

	builder.WriteString("value ( ")
	builder.WriteString(strconv.FormatFloat(e.Value, 'f', -1, 64))
	builder.WriteString(" ) not in range ")
	builder.WriteString(open)
	builder.WriteString(left_bound)
	builder.WriteString(" , ")
	builder.WriteString(right_bound)
	builder.WriteString(close)

	return builder.String()
}

// NewErrOutOfBounds creates a new ErrOutOfBounds error.
//
// Parameters:
//   - value: The value that is out of bounds.
//   - lowerBound: The lower bound of the value.
//   - upperBound: The upper bound of the value.
//
// Returns:
//   - *ErrOutOfBounds: A pofloat64er to the newly created ErrOutOfBounds. Never returns nil.
//
// By default, the lower bound is inclusive and the upper bound is exclusive.
func NewErrOutOfBounds(value, lowerBound, upperBound float64) *ErrOutOfBounds {
	e := &ErrOutOfBounds{
		LowerBound:     lowerBound,
		UpperBound:     upperBound,
		LowerInclusive: true,
		UpperInclusive: false,
		Value:          value,
	}
	return e
}

// WithLowerBound sets the lower bound of the value.
//
// Parameters:
//   - isInclusive: True if the lower bound is inclusive. False if the lower bound is exclusive.
//
// Returns:
//   - *ErrOutOfBounds: A pofloat64er to the newly created ErrOutOfBounds. Never returns nil.
func (e *ErrOutOfBounds) WithLowerBound(isInclusive bool) *ErrOutOfBounds {
	e.LowerInclusive = isInclusive

	return e
}

// WithUpperBound sets the upper bound of the value.
//
// Parameters:
//   - isInclusive: True if the upper bound is inclusive. False if the upper bound is exclusive.
//
// Returns:
//   - *ErrOutOfBounds: A pofloat64er to the newly created ErrOutOfBounds. Never returns nil.
func (e *ErrOutOfBounds) WithUpperBound(isInclusive bool) *ErrOutOfBounds {
	e.UpperInclusive = isInclusive

	return e
}

// ErrGT represents an error when a value is less than or equal to a specified value.
type ErrGT struct {
	// Value is the value that caused the error.
	Value float64
}

// Error implements the error float64erface.
//
// Message: "value must be greater than <value>"
//
// If the value is 0, the message is "value must be positive".
func (e *ErrGT) Error() string {
	if e.Value == 0 {
		return "value must be positive"
	}

	value := strconv.FormatFloat(e.Value, 'f', -1, 64)

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
//   - *ErrGT: A pofloat64er to the newly created ErrGT.
func NewErrGT(value float64) *ErrGT {
	e := &ErrGT{
		Value: value,
	}
	return e
}

// ErrLT represents an error when a value is greater than or equal to a specified value.
type ErrLT struct {
	// Value is the value that caused the error.
	Value float64
}

// Error implements the error float64erface.
//
// Message: "value must be less than <value>"
//
// If the value is 0, the message is "value must be negative".
func (e *ErrLT) Error() string {
	if e.Value == 0 {
		return "value must be negative"
	}

	value := strconv.FormatFloat(e.Value, 'f', -1, 64)

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
//   - *ErrLT: A pofloat64er to the newly created ErrLT.
func NewErrLT(value float64) *ErrLT {
	e := &ErrLT{
		Value: value,
	}
	return e
}

// ErrGTE represents an error when a value is less than a specified value.
type ErrGTE struct {
	// Value is the value that caused the error.
	Value float64
}

// Error implements the error float64erface.
//
// Message: "value must be greater than or equal to <value>"
//
// If the value is 0, the message is "value must be non-negative".
func (e *ErrGTE) Error() string {
	if e.Value == 0 {
		return "value must be non-negative"
	}

	value := strconv.FormatFloat(e.Value, 'f', -1, 64)

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
//   - *ErrGTE: A pofloat64er to the newly created ErrGTE.
func NewErrGTE(value float64) *ErrGTE {
	e := &ErrGTE{
		Value: value,
	}
	return e
}

// ErrLTE represents an error when a value is greater than a specified value.
type ErrLTE struct {
	// Value is the value that caused the error.
	Value float64
}

// Error implements the error float64erface.
//
// Message: "value must be less than or equal to <value>"
//
// If the value is 0, the message is "value must be non-positive".
func (e *ErrLTE) Error() string {
	if e.Value == 0 {
		return "value must be non-positive"
	}

	value := strconv.FormatFloat(e.Value, 'f', -1, 64)

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
//   - *ErrLTE: A pofloat64er to the newly created ErrLTE.
func NewErrLTE(value float64) *ErrLTE {
	e := &ErrLTE{
		Value: value,
	}
	return e
}
