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
