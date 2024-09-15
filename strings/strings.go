package strings

import (
	"fmt"
	"strconv"
	"strings"
)

// GoStringOf returns a string representation of the element.
//
// Parameters:
//   - elem: The element to get the string representation of.
//
// Returns:
//   - string: The string representation of the element.
//
// Behaviors:
//   - If the element is nil, the function returns "nil".
//   - If the element implements the fmt.GoStringer interface, the function
//     returns the result of the GoString method.
//   - If the element implements the fmt.Stringer interface, the function
//     returns the result of the String method.
//   - If the element is a string, the function returns the string enclosed in
//     double quotes.
//   - If the element is an error, the function returns the error message
//     enclosed in double quotes.
//   - Otherwise, the function returns the result of the %#v format specifier.
func GoStringOf(elem any) string {
	if elem == nil {
		return "nil"
	}

	switch elem := elem.(type) {
	case fmt.GoStringer:
		return elem.GoString()
	case fmt.Stringer:
		return strconv.Quote(elem.String())
	case string:
		return strconv.Quote(elem)
	case error:
		return strconv.Quote(elem.Error())
	default:
		return fmt.Sprintf("%#v", elem)
	}
}

// TrimEmpty removes empty strings from a slice of strings.
// Empty spaces at the beginning and end of the strings are also removed from
// the strings.
//
// Parameters:
//   - values: The slice of strings to trim.
//
// Returns:
//   - []string: The slice of strings with empty strings removed.
func TrimEmpty(values []string) []string {
	if len(values) == 0 {
		return values
	}

	res := make([]string, 0, len(values))

	for _, value := range values {
		str := strings.TrimSpace(value)
		if str != "" {
			res = append(res, str)
		}
	}

	return res[:len(res):len(res)]
}

// EitherOrString is a function that returns a string representation of a slice
// of strings. Empty strings are ignored.
//
// Parameters:
//   - values: The values to convert to a string.
//
// Returns:
//   - string: The string representation.
//
// Example:
//
//	EitherOrString([]string{"a", "b", "c"}, false) // "a, b or c"
func EitherOrString(values []string) string {
	values = TrimEmpty(values)

	if len(values) == 0 {
		return ""
	}

	if len(values) == 1 {
		return values[0]
	}

	var builder strings.Builder

	builder.WriteString("either ")

	if len(values) > 2 {
		builder.WriteString(strings.Join(values[:len(values)-1], ", "))
		builder.WriteRune(',')
	} else {
		builder.WriteString(values[0])
	}

	builder.WriteString(" or ")
	builder.WriteString(values[len(values)-1])

	return builder.String()
}

// OrString is a function that returns a string representation of a slice of
// strings. Empty strings are ignored.
//
// Parameters:
//   - values: The values to convert to a string.
//   - is_negative: True if the string should use "nor" instead of "or", false
//     otherwise.
//
// Returns:
//   - string: The string representation.
//
// Example:
//
//	OrString([]string{"a", "b", "c"}, true) // "a, b, nor c"
func OrString(values []string, is_negative bool) string {
	values = TrimEmpty(values)
	if len(values) == 0 {
		return ""
	}

	if len(values) == 1 {
		return values[0]
	}

	var sep string

	if is_negative {
		sep = " nor "
	} else {
		sep = " or "
	}

	var builder strings.Builder

	if len(values) > 2 {
		builder.WriteString(strings.Join(values[:len(values)-1], ", "))
		builder.WriteRune(',')
	} else {
		builder.WriteString(values[0])
	}

	builder.WriteString(sep)
	builder.WriteString(values[len(values)-1])

	return builder.String()
}
