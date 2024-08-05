package ints

import (
	"strconv"
	"strings"
)

// ErrOutOfBounds represents an error when a value is out of bounds.
type ErrOutOfBounds struct {
	// LowerBound is the lower bound of the value.
	LowerBound int

	// UpperBound is the upper bound of the value.
	UpperBound int

	// LowerInclusive is true if the lower bound is inclusive.
	LowerInclusive bool

	// UpperInclusive is true if the upper bound is inclusive.
	UpperInclusive bool

	// Value is the value that is out of bounds.
	Value int
}

// Error implements the error interface.
//
// Message: "value ( <value> ) not in range <lower_bound> , <upper_bound>"
func (e *ErrOutOfBounds) Error() string {
	left_bound := strconv.Itoa(e.LowerBound)
	right_bound := strconv.Itoa(e.UpperBound)

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
	builder.WriteString(strconv.Itoa(e.Value))
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
//   - *ErrOutOfBounds: A pointer to the newly created ErrOutOfBounds. Never returns nil.
//
// By default, the lower bound is inclusive and the upper bound is exclusive.
func NewErrOutOfBounds(value, lowerBound, upperBound int) *ErrOutOfBounds {
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
//   - *ErrOutOfBounds: A pointer to the newly created ErrOutOfBounds. Never returns nil.
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
//   - *ErrOutOfBounds: A pointer to the newly created ErrOutOfBounds. Never returns nil.
func (e *ErrOutOfBounds) WithUpperBound(isInclusive bool) *ErrOutOfBounds {
	e.UpperInclusive = isInclusive

	return e
}

// ErrInvalidBase represents an error when the base is invalid.
type ErrInvalidBase struct{}

// Error implements the error interface.
//
// Message: "base must be positive"
func (e *ErrInvalidBase) Error() string {
	return "base must be positive"
}

// NewErrInvalidBase creates a new ErrInvalidBase error.
//
// Returns:
//   - *ErrInvalidBase: A pointer to the newly created ErrInvalidBase. Never returns nil.
func NewErrInvalidBase() *ErrInvalidBase {
	return &ErrInvalidBase{}
}
