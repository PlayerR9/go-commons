package assert

import (
	"fmt"
	"strconv"
	"strings"
)

// BoolAssert is the struct that is used to assert boolean values.
type BoolAssert struct {
	// value is the value to assert.
	value bool

	// is_true is the condition to assert.
	is_true bool
}

// Panic will panic if the condition is not met.
//
// The error message is "expected <name> to <message>; got <value> instead" where
// <name> is the name of the assertion, <message> is the message of the condition
// and <value> is the value of the assertion. Finally, this error message is used
// within the *ErrAssertionFailed error.
func (a BoolAssert) Message(target Target, is_negative bool) string {
	var builder strings.Builder

	builder.WriteString("expected ")
	builder.WriteString(target.String())

	if is_negative {
		builder.WriteString(" to not ")
	} else {
		builder.WriteString(" to ")
	}

	if a.is_true {
		builder.WriteString("be true")
	} else {
		builder.WriteString("be false")
	}

	builder.WriteString("; got ")
	builder.WriteString(strconv.Quote(fmt.Sprintf("%v", a.value)))
	builder.WriteString(" instead")

	return builder.String()
}

// Verify returns true iff the condition is met.
//
// Returns:
//   - bool: true if the condition is met. False otherwise.
func (a BoolAssert) Verify() bool {
	return a.is_true != a.value
}

// NewBoolAssert returns a new BoolAssert struct. By default, the condition checks if the value is true.
//
// Parameters:
//   - value: the value to assert.
//
// Returns:
//   - *BoolAssert: the new BoolAssert struct. Never returns nil.
func NewBoolAssert(value bool) *BoolAssert {
	return &BoolAssert{
		value:   value,
		is_true: true,
	}
}

// IsTrue is the assertion for checking if the value is true.
//
// If any other condition is specified, the furthest condition overwrites any
// other condition.
//
// Returns:
//   - *BoolAssert: the assertion for chaining. Nil only if the receiver is nil.
func (a *BoolAssert) IsTrue() *BoolAssert {
	if a == nil {
		return nil
	}

	a.is_true = true

	return a
}

// IsFalse is the assertion for checking if the value is false.
//
// If any other condition is specified, the furthest condition overwrites any
// other condition.
//
// Returns:
//   - *BoolAssert: the assertion for chaining. Nil only if the receiver is nil.
func (a *BoolAssert) IsFalse() *BoolAssert {
	if a == nil {
		return nil
	}

	a.is_true = false

	return a
}
