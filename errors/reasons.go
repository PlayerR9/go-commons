package errors

import (
	"cmp"
	"errors"
	"fmt"
	"reflect"
	"strconv"
	"strings"
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

// ValueGT creates a new ErrGT error with the specified value.
//
// Parameters:
//   - value: The minimum value that is not allowed.
//
// Returns:
//   - *ErrGT: A pointer to the newly created ErrGT. Never returns nil.
func ValueGT[T cmp.Ordered](value T) string {
	return fmt.Sprintf("value must ge greater than %v", value)
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
func (e ErrOutOfBounds) Error() string {
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
	return &ErrOutOfBounds{
		LowerBound:     lowerBound,
		UpperBound:     upperBound,
		LowerInclusive: true,
		UpperInclusive: false,
		Value:          value,
	}
}

// WithLowerBound sets the lower bound of the value.
//
// Parameters:
//   - is_inclusive: True if the lower bound is inclusive. False if the lower bound is exclusive.
//
// Returns:
//   - *ErrOutOfBounds: A pointer to the newly created ErrOutOfBounds.
//
// Only when the receiver is nil, this function returns nil.
func (e *ErrOutOfBounds) WithLowerBound(is_inclusive bool) *ErrOutOfBounds {
	if e == nil {
		return nil
	}

	e.LowerInclusive = is_inclusive

	return e
}

// WithUpperBound sets the upper bound of the value.
//
// Parameters:
//   - is_inclusive: True if the upper bound is inclusive. False if the upper bound is exclusive.
//
// Returns:
//   - *ErrOutOfBounds: A pointer to the newly created ErrOutOfBounds.
//
// Only when the receiver is nil, this function returns nil.
func (e *ErrOutOfBounds) WithUpperBound(is_inclusive bool) *ErrOutOfBounds {
	if e == nil {
		return nil
	}

	e.UpperInclusive = is_inclusive

	return e
}

// ErrValue represents an error when a value is not expected.
type ErrValue struct {
	// Kind is the name of the thing that was expected.
	Kind string

	// Expected is the value that was expected.
	Expected any

	// Got is the value that was received.
	Got any

	// ShouldQuote is true if the expected and got values should be quoted,
	// false otherwise.
	ShouldQuote bool
}

// Error implements the error interface.
//
// Message:
//
//	"expected <kind> to be <expected>, got <got> instead"
func (e ErrValue) Error() string {
	var builder strings.Builder

	builder.WriteString("expected ")

	if e.Kind != "" {
		builder.WriteString(e.Kind)
		builder.WriteString(" to be ")
	}

	if e.Expected == nil {
		builder.WriteString("nothing")
	} else if e.ShouldQuote {
		fmt.Fprintf(&builder, "%q", e.Expected)
	} else {
		fmt.Fprintf(&builder, "%s", e.Expected)
	}

	builder.WriteString(", got ")

	if e.Got == nil {
		builder.WriteString("nothing")
	} else if e.ShouldQuote {
		fmt.Fprintf(&builder, "%q", e.Got)
	} else {
		fmt.Fprintf(&builder, "%s", e.Got)
	}

	builder.WriteString(" instead")

	return builder.String()
}

// NewErrValue creates a new ErrValue error.
//
// Parameters:
//   - kind: The name of the thing that was expected.
//   - expected: The value that was expected.
//   - got: The value that was received.
//   - should_quote: True if the expected and got values should be quoted,
//     false otherwise.
//
// Returns:
//   - *ErrValue: A pointer to the newly created ErrValue. Never returns nil.
func NewErrValue(kind string, expected, got any, should_quote bool) *ErrValue {
	return &ErrValue{
		Kind:        kind,
		Expected:    expected,
		Got:         got,
		ShouldQuote: should_quote,
	}
}

// ErrValues represents an error when multiple value are not expected.
type ErrValues[T any] struct {
	// Kind is the name of the thing that was expected.
	Kind string

	// Expecteds is the values that were expected.
	Expecteds []T

	// Got is the value that was received.
	Got any

	// ShouldQuote is true if the expected and got values should be quoted,
	// false otherwise.
	ShouldQuote bool
}

// Error implements the error interface.
//
// Message:
//
//	"expected <kind> to be <expected>, got <got> instead"
func (e ErrValues[T]) Error() string {
	var builder strings.Builder

	builder.WriteString("expected ")

	if e.Kind != "" {
		builder.WriteString(e.Kind)
		builder.WriteString(" to be ")
	}

	switch len(e.Expecteds) {
	case 0:
		builder.WriteString("nothing")
	case 1:
		if e.ShouldQuote {
			builder.WriteString(strconv.Quote(fmt.Sprintf("%v", e.Expecteds[0])))
		} else {
			fmt.Fprintf(&builder, "%v", e.Expecteds[0])
		}
	default:
		elems := make([]string, 0, len(e.Expecteds))

		if e.ShouldQuote {
			for i := 0; i < len(e.Expecteds); i++ {
				elems = append(elems, strconv.Quote(fmt.Sprintf("%v", e.Expecteds[i])))
			}
		} else {
			for i := 0; i < len(e.Expecteds); i++ {
				elems = append(elems, fmt.Sprintf("%v", e.Expecteds[i]))
			}
		}

		builder.WriteString("either ")
		builder.WriteString(strings.Join(elems[:len(elems)-1], ", "))
		builder.WriteString(" or ")
		builder.WriteString(elems[len(elems)-1])
	}

	builder.WriteString(", got ")

	if e.Got == nil {
		builder.WriteString("nothing")
	} else if e.ShouldQuote {
		fmt.Fprintf(&builder, "%q", e.Got)
	} else {
		fmt.Fprintf(&builder, "%s", e.Got)
	}

	builder.WriteString(" instead")

	return builder.String()
}

// NewErrValues creates a new ErrValues error.
//
// Parameters:
//   - kind: The name of the thing that was expected.
//   - expected: The values that were expected.
//   - got: The value that was received.
//   - should_quote: True if the expected and got values should be quoted,
//     false otherwise.
//
// Returns:
//   - *ErrValue: A pointer to the newly created ErrValue. Never returns nil.
func NewErrValues[T any](kind string, expected []T, got any, should_quote bool) *ErrValue {
	return &ErrValue{
		Kind:        kind,
		Expected:    expected,
		Got:         got,
		ShouldQuote: should_quote,
	}
}
