package strings

import (
	"fmt"
	"strconv"
)

// QuoteStrings is a function that quotes a slice of strings in-place.
//
// Parameters:
//   - values: The values to quote.
func QuoteStrings(values []string) {
	if len(values) == 0 {
		return
	}

	for i := 0; i < len(values); i++ {
		values[i] = strconv.Quote(values[i])
	}
}

// SliceOfRunes is a function that returns a slice of strings
// from a slice of runes.
//
// Parameters:
//   - values: The values to convert to a slice of strings.
//
// Returns:
//   - []string: The slice of strings.
func SliceOfRunes(values []rune) []string {
	if len(values) == 0 {
		return nil
	}

	elems := make([]string, 0, len(values))

	for _, value := range values {
		elems = append(elems, string(value))
	}

	return elems
}

// SliceOfStringer is a function that returns a slice of strings
// from a slice of stringers.
//
// Parameters:
//   - values: The values to convert to a slice of strings.
//
// Returns:
//   - []string: The slice of strings.
func SliceOfStringer[T fmt.Stringer](values []T) []string {
	if len(values) == 0 {
		return nil
	}

	elems := make([]string, 0, len(values))

	for _, value := range values {
		elems = append(elems, value.String())
	}

	return elems
}
