package errors

import (
	"fmt"
	"strconv"
	"strings"
)

// Got returns the string representation of a value.
//
// Parameters:
//   - quote: A flag indicating whether the value should be quoted.
//   - got: The value to get the string representation of.
//
// Returns:
//   - string: The string "got <value> instead"
//
// If the value is nil, the function returns "got nothing instead" regardless of the flag.
func Got(quote bool, got any) string {
	var builder strings.Builder

	builder.WriteString("got ")

	if got == nil {
		builder.WriteString("nothing")
	} else {
		str := fmt.Sprintf("%v", got)

		if quote {
			str = strconv.Quote(str)
		}

		builder.WriteString(str)
	}

	builder.WriteString(" instead")

	return builder.String()
}

// StringOfSlice returns the string representation of a slice of values.
//
// Parameters:
//   - quoted: A flag indicating whether the values should be quoted.
//   - elems: The slice of values to get the string representation of.
//
// Returns:
//   - []string: The string representation of the slice of values.
func StringOfSlice[T any](quoted bool, elems []T) []string {
	if len(elems) == 0 {
		return nil
	}

	values := make([]string, 0, len(elems))

	for i := 0; i < len(elems); i++ {
		str := strings.TrimSpace(fmt.Sprintf("%v", elems[i]))
		if str != "" {
			values = append(values, str)
		}
	}

	values = values[:len(values):len(values)]

	if len(values) == 0 {
		return nil
	}

	if !quoted {
		return values
	}

	for i := 0; i < len(values); i++ {
		values[i] = strconv.Quote(values[i])
	}

	return values
}
