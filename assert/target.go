package assert

import (
	"fmt"
	"strconv"
	"strings"
)

// Asserter is the interface that is used to assert values.
type Asserter interface {
	// Verify checks if the value satisfies the condition.
	//
	// Parameters:
	//   - cond: the condition to check.
	//
	// Returns:
	//   - bool: true if the condition is met. false otherwise.
	Verify() bool

	// Message returns the message that is shown when the condition is not met.
	//
	// Parameters:
	//   - target: The target being asserted.
	//   - is_negative: True if the condition is negated. False otherwise.
	//
	// Returns:
	//   - string: the message.
	Message(target Target, is_negative bool) string
}

//go:generate stringer -type=AssertTargetType -linecomment

// AssertTargetType is the type of the target being asserted.
type AssertTargetType int

const (
	// ReceiverFunction is the type of the receiver function being asserted.
	ReceiverFunction AssertTargetType = iota // receiver function

	// Function is the type of the function being asserted. This is used in any other
	// non-receiver function.
	Function // function

	// Struct is the type of the struct being asserted. Mostly used in receiver functions.
	Struct // struct

	// Variable is the type of the variable being asserted. Used within functions or receiver functions
	// to check anything that is not a parameter.
	Variable // variable

	// Parameter is the type of the parameter being asserted. Used within functions or receiver functions
	// to check parameters.
	Parameter // parameter

	// Condition is the type of the condition being asserted.
	Condition // condition

	// Other is the type of the other element being asserted that does not fit any other type.
	Other // element
)

// Target is the target being asserted.
type Target struct {
	// type_ is the type of the target being asserted.
	type_ AssertTargetType

	// sign is the signature of the target being asserted.
	sign string
}

func (t Target) String() string {
	if t.type_ == Other {
		return t.sign
	} else {
		return t.type_.String() + " " + t.sign
	}
}

// NewReceiverFunction returns a new Target with the type of ReceiverFunction.
//
// Parameters:
//   - receiver: the receiver of the function. If empty, no receiver is used.
//   - fn: the function name. If empty, "func" is used.
//   - format: the format string to generate the function signature.
//   - a: the arguments to pass to the format string.
//
// Returns:
//   - Target: the new Target.
func NewReceiverFunction(receiver, fn, format string, a ...any) Target {
	var builder strings.Builder

	if receiver != "" {
		builder.WriteString(receiver)
		builder.WriteRune('.')
	}

	if fn == "" {
		builder.WriteString("func(")
	} else {
		builder.WriteString(fn)
		builder.WriteRune('(')
	}

	fmt.Fprintf(&builder, format, a...)
	builder.WriteRune(')')

	return Target{
		type_: ReceiverFunction,
		sign:  builder.String(),
	}
}

// NewFunction returns a new Target with the type of Function.
//
// Parameters:
//   - fn: the function name. If empty, "func" is used.
//   - format: the format string to generate the function signature.
//   - a: the arguments to pass to the format string.
//
// Returns:
//   - Target: the new Target.
func NewFunction(fn, format string, a ...any) Target {
	var builder strings.Builder

	if fn == "" {
		builder.WriteString("func(")
	} else {
		builder.WriteString(fn)
		builder.WriteRune('(')
	}

	fmt.Fprintf(&builder, format, a...)
	builder.WriteRune(')')

	return Target{
		type_: Function,
		sign:  builder.String(),
	}
}

// NewStruct returns a new Target with the type of Struct.
//
// Parameters:
//   - name: the name of the struct. If empty, "struct{}" is used.
//
// Returns:
//   - Target: the new Target.
func NewStruct(name string) Target {
	if name == "" {
		name = "struct{}"
	}

	return Target{
		type_: Struct,
		sign:  name,
	}
}

// NewVariable returns a new Target with the type of Variable.
//
// Parameters:
//   - name: the name of the variable. If empty, "_" is used.
//
// Returns:
//   - Target: the new Target.
func NewVariable(name string) Target {
	if name == "" {
		name = "_"
	} else {
		name = strconv.Quote(name)
	}

	return Target{
		type_: Variable,
		sign:  "(" + name + ")",
	}
}

// NewParameter returns a new Target with the type of Parameter.
//
// Parameters:
//   - name: the name of the parameter. If empty, "_" is used.
//
// Returns:
//   - Target: the new Target.
func NewParameter(name string) Target {
	if name == "" {
		name = "_"
	} else {
		name = strconv.Quote(name)
	}

	return Target{
		type_: Parameter,
		sign:  "(" + name + ")",
	}
}

// NewCondition returns a new Target with the type of Condition.
//
// Parameters:
//   - cond: the condition. If empty, "true" is used.
//
// Returns:
//   - Target: the new Target.
func NewCondition(cond string) Target {
	if cond == "" {
		cond = "true"
	}

	return Target{
		type_: Condition,
		sign:  cond,
	}
}

// NewOther returns a new Target with the type of Other.
//
// Parameters:
//   - sign: the name of the target. If empty, the string representation of Other is used.
//
// Returns:
//   - Target: the new Target.
func NewOther(sign string) Target {
	if sign == "" {
		sign = Other.String()
	}

	return Target{
		type_: Other,
		sign:  sign,
	}
}
