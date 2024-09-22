package errors

import (
	"fmt"
	"strconv"
	"strings"
)

// ExpectedValue creates a new ErrValue error.
//
// Parameters:
//   - kind: The name of the thing that was expected.
//   - expected: The value that was expected.
//   - got: The value that was received.
//   - should_quote: True if the expected and got values should be quoted,
//     false otherwise.
//
// Returns:
//   - string: The string representation of the error.
func ExpectedValue(kind string, expected, got any, should_quote bool) string {
	var builder strings.Builder

	builder.WriteString("expected ")

	if kind != "" {
		builder.WriteString(kind)
		builder.WriteString(" to be ")
	}

	if expected == nil {
		builder.WriteString("nothing")
	} else if should_quote {
		fmt.Fprintf(&builder, "%q", fmt.Sprintf("%v", expected))
	} else {
		fmt.Fprintf(&builder, "%v", expected)
	}

	builder.WriteString(", got ")

	if got == nil {
		builder.WriteString("nothing")
	} else if should_quote {
		fmt.Fprintf(&builder, "%q", fmt.Sprintf("%v", got))
	} else {
		fmt.Fprintf(&builder, "%v", got)
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
func ExpectedValues[T any](kind string, expecteds []T, got any, should_quote bool) string {
	var builder strings.Builder

	builder.WriteString("expected ")

	if kind != "" {
		builder.WriteString(kind)
		builder.WriteString(" to be ")
	}

	switch len(expecteds) {
	case 0:
		builder.WriteString("nothing")
	case 1:
		if should_quote {
			builder.WriteString(strconv.Quote(fmt.Sprintf("%v", expecteds[0])))
		} else {
			fmt.Fprintf(&builder, "%v", expecteds[0])
		}
	default:
		elems := make([]string, 0, len(expecteds))

		if should_quote {
			for i := 0; i < len(expecteds); i++ {
				elems = append(elems, strconv.Quote(fmt.Sprintf("%v", expecteds[i])))
			}
		} else {
			for i := 0; i < len(expecteds); i++ {
				elems = append(elems, fmt.Sprintf("%v", expecteds[i]))
			}
		}

		builder.WriteString("either ")
		builder.WriteString(strings.Join(elems[:len(elems)-1], ", "))
		builder.WriteString(" or ")
		builder.WriteString(elems[len(elems)-1])
	}

	builder.WriteString(", got ")

	if got == nil {
		builder.WriteString("nothing")
	} else if should_quote {
		fmt.Fprintf(&builder, "%s", fmt.Sprintf("%v", got))
	} else {
		fmt.Fprintf(&builder, "%v", got)
	}

	builder.WriteString(" instead")

	return builder.String()
}
