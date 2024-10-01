package runes

import (
	"strconv"
	"strings"
)

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
//	EitherOrString([]rune{'a', 'b', 'c'}, false) // "either a, b or c"
func EitherOrString(elems []rune, quote bool) string {
	if len(elems) == 0 {
		return ""
	} else if len(elems) == 1 {
		if quote {
			return strconv.QuoteRune(elems[0])
		} else {
			return string(elems[0])
		}
	}

	values := make([]string, 0, len(elems))

	if quote {
		for _, elem := range elems {
			values = append(values, strconv.QuoteRune(elem))
		}
	} else {
		for _, elem := range elems {
			values = append(values, string(elem))
		}
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

// Indices returns the indices of the separator in the data.
//
// Parameters:
//   - data: The data.
//   - sep: The separator.
//   - exclude_sep: Whether the separator is inclusive. If set to true, the indices will point to the character right after the
//     separator. Otherwise, the indices will point to the separator itself.
//
// Returns:
//   - []int: The indices.
func IndicesOf(data []rune, sep rune, exclude_sep bool) []int {
	if len(data) == 0 {
		return nil
	}

	var indices []int

	for i := 0; i < len(data); i++ {
		if data[i] == sep {
			indices = append(indices, i)
		}
	}

	if len(indices) == 0 {
		return nil
	}

	if exclude_sep {
		for i := 0; i < len(indices); i++ {
			indices[i] += 1
		}
	}

	return indices
}
