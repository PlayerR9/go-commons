package assert

import (
	"cmp"
	"fmt"
	"slices"
	"strconv"
	"strings"
)

// OrderedAssert is the struct that is used to assert values that are ordered.
type OrderedAssert[T cmp.Ordered] struct {
	// value is the value to assert.
	value T

	// cond is the condition to assert.
	cond Conditioner[T]
}

// Message implements the Asserter interface.
//
// The error message is "expected <name> to <message>; got <value> instead" where:
//   - <name> is the name of the assertion.
//   - <message> is the message of the condition.
//   - <value> is the value of the assertion.
func (a OrderedAssert[T]) Message(target Target, is_negative bool) string {
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
func (a OrderedAssert[T]) Verify() bool {
	if a.cond == nil {
		return true
	}

	return a.cond.Verify(a.value)
}

// NewOrderedAssert returns a new OrderedAssert struct.
//
// Parameters:
//   - value: the value to assert.
//
// Returns:
//   - *OrderedAssert: the new OrderedAssert struct. Never returns nil.
func NewOrderedAssert[T cmp.Ordered](value T) *OrderedAssert[T] {
	return &OrderedAssert[T]{
		value: value,
	}
}

// Equal is the assertion for checking if the value is equal to another value.
//
// If any other condition is specified, the furthest condition overwrites any
// other condition.
//
// Parameters:
//   - b: the other value to compare with.
//
// Returns:
//   - *OrderedAssert[T]: the assertion for chaining. Never returns nil.
func (a *OrderedAssert[T]) Equal(b T) *OrderedAssert[T] {
	a.cond = &EqualCond[T]{other: b}

	return a
}

// GreaterThan is the assertion for checking if the value is greater than another value.
//
// If any other condition is specified, the furthest condition overwrites any
// other condition.
//
// Parameters:
//   - b: the other value to compare with.
//
// Returns:
//   - *Assertion: the assertion for chaining. Never returns nil.
func (a *OrderedAssert[T]) GreaterThan(b T) *OrderedAssert[T] {
	a.cond = &GreaterThanCond[T]{other: b}
	return a
}

// LessThan is the assertion for checking if the value is less than another value.
//
// If any other condition is specified, the furthest condition overwrites any
// other condition.
//
// Parameters:
//   - b: the other value to compare with.
//
// Returns:
//   - *Assertion: the assertion for chaining. Never returns nil.
func (a *OrderedAssert[T]) LessThan(b T) *OrderedAssert[T] {
	a.cond = &LessThanCond[T]{other: b}

	return a
}

// GreaterOrEqualThan is the assertion for checking if the value is greater or equal than another value.
//
// If any other condition is specified, the furthest condition overwrites any
// other condition.
//
// Parameters:
//   - b: the other value to compare with.
//
// Returns:
//   - *Assertion: the assertion for chaining. Never returns nil.
func (a *OrderedAssert[T]) GreaterOrEqualThan(b T) *OrderedAssert[T] {
	a.cond = &GreaterOrEqualThanCond[T]{other: b}

	return a
}

// LessOrEqualThan is the assertion for checking if the value is less or equal than another value.
//
// If any other condition is specified, the furthest condition overwrites any
// other condition.
//
// Parameters:
//   - b: the other value to compare with.
//
// Returns:
//   - *Assertion: the assertion for chaining. Never returns nil.
func (a *OrderedAssert[T]) LessOrEqualThan(b T) *OrderedAssert[T] {
	a.cond = &LessOrEqualThanCond[T]{other: b}

	return a
}

// InRange is the assertion for checking if the value is in a range.
//
// If any other condition is specified, the furthest condition overwrites any
// other condition.
//
// Parameters:
//   - min: the minimum value of the range.
//   - max: the maximum value of the range.
//
// Returns:
//   - *Assertion: the assertion for chaining. Never returns nil.
//
// If min is greater than max, the min and max values will be swapped. Moreover, if
// min is equal to max, the assertion will be equal to the EqualCond[T] with the min value.
func (a *OrderedAssert[T]) InRange(min, max T) *OrderedAssert[T] {
	if min > max {
		min, max = max, min
	}

	if min == max {
		a.cond = &EqualCond[T]{other: min}
	} else {
		a.cond = &InRangeCond[T]{min: min, max: max}
	}

	return a
}

// Zero is the assertion for checking if the value is the zero value for its type.
//
// If any other condition is specified, the furthest condition overwrites any
// other condition.
//
// Returns:
//   - *Assertion: the assertion for chaining. Never returns nil.
func (a *OrderedAssert[T]) Zero() *OrderedAssert[T] {
	a.cond = &ZeroCond[T]{}

	return a
}

// In is the assertion for checking if the value is in a list of values.
//
// If any other condition is specified, the furthest condition overwrites any
// other condition.
//
// Parameters:
//   - values: the list of values to check against.
//
// Returns:
//   - *Assertion: the assertion for chaining. Never returns nil.
//
// The list is sorted in ascending order and duplicates are removed. As a special case,
// if only one value is provided, the assertion will be equal to the EqualCond[T] with
// that value.
func (a *OrderedAssert[T]) In(values ...T) *OrderedAssert[T] {
	if len(values) > 2 {
		sorted := make([]T, 0, len(values))

		for _, val := range values {
			pos, ok := slices.BinarySearch(sorted, val)
			if !ok {
				sorted = slices.Insert(sorted, pos, val)
			}
		}

		values = sorted[:len(sorted):len(sorted)]
	}

	switch len(values) {
	case 0:
		a.cond = &InCond[T]{values: []T{}}
	case 1:
		a.cond = &EqualCond[T]{other: values[0]}
	default:
		a.cond = &InCond[T]{values: values}
	}

	return a
}
