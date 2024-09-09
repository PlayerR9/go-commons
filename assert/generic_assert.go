package assert

import (
	"fmt"
	"strconv"
	"strings"
)

// Assertion is the struct that is used to perform assertions.
type GenericAssert[T any] struct {
	// name is the name of the value.
	name string

	// value is the value to assert.
	value T

	// cond is the condition to perform.
	cond Conditioner[T]

	// negative is true if the assertion should be negated.
	negative bool
}

// Panic will panic if the condition is not met.
//
// The error message is "expected <name> to <message>; got <value> instead" where
// <name> is the name of the assertion, <message> is the message of the condition
// and <value> is the value of the assertion. Finally, this error message is used
// within the *ErrAssertionFailed error.
func (a GenericAssert[T]) Message(target Target, is_negative bool) string {
	var builder strings.Builder

	builder.WriteString("expected ")
	builder.WriteString(target.String())

	if is_negative {
		builder.WriteString(" to not ")
	} else {
		builder.WriteString(" to ")
	}

	builder.WriteString(a.cond.Message())
	builder.WriteString("; got ")
	builder.WriteString(strconv.Quote(fmt.Sprintf("%v", a.value)))
	builder.WriteString(" instead")

	return builder.String()
}

// Verify returns true iff the condition is met.
//
// Returns:
//   - bool: true if the condition is met. False otherwise.
func (a *GenericAssert[T]) Verify() bool {
	if a.cond == nil {
		return true
	}

	return a.cond.Verify(a.value)
}

// NewGenericAssert returns a new GenericAssert struct.
//
// Parameters:
//   - name: the name of the value.
//   - value: the value to assert.
//
// Returns:
//   - *GenericAssert: the new GenericAssert struct. Never returns nil.
func NewGenericAssert[T any](name string, value T) *GenericAssert[T] {
	return &GenericAssert[T]{
		name:  name,
		value: value,
	}
}

// Satisfies is the assertion for checking custom conditions.
//
// If any other condition is specified, the furthest condition overwrites any
// other condition. However, if cond is nil, this function will be a no-op.
//
// Parameters:
//   - cond: the condition to check.
//
// Returns:
//   - *Assertion: the assertion for chaining. Never returns nil.
func (a *GenericAssert[T]) Satisfies(cond Conditioner[T]) *GenericAssert[T] {
	if cond == nil {
		return a
	}

	a.cond = cond

	return a
}

// Applies is the same as Satisfies but without needing to provide a custom definition
// that implements Conditioner. Best used for checks that are only done once.
//
// If any other condition is specified, the furthest condition overwrites any
// other condition.
//
// Parameters:
//   - msg: the message of the condition.
//   - cond: the condition to check.
//
// Returns:
//   - *Assertion: the assertion for chaining. Never returns nil.
func (a *GenericAssert[T]) Applies(msg func() string, cond func(value T) bool) *GenericAssert[T] {
	a.cond = &GenericCond[T]{
		message: msg,
		verify:  cond,
	}

	return a
}
