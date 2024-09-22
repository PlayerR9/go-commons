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

// StringOfSlice is a function that returns a slice of strings
// from a slice of values.
//
// Parameters:
//   - slice: The values to convert to a slice of strings.
//
// Returns:
//   - []string: The slice of strings.
func StringOfSlice[T any](slice []T) []string {
	if len(slice) == 0 {
		return nil
	}

	elems := make([]string, 0, len(slice))

	for i := 0; i < len(slice); i++ {
		elems = append(elems, fmt.Sprintf("%v", slice[i]))
	}

	return elems
}

// ExpectedValue is a function that returns an expected value
// message.
//
// Parameters:
//   - kind: The name of the thing that was expected.
//   - expected: The value that was expected.
//   - got: The value that was received.
//
// Returns:
//   - string: The expected value message.
//
// Behaviors:
//   - If 'expected' is an empty string, it will be replaced with 'nothing'. Same for 'got'.
func ExpectedValue(kind, expected, got string) string {
	if expected == "" {
		expected = "nothing"
	}

	if got == "" {
		got = "nothing"
	}

	if kind == "" {
		return fmt.Sprintf("expected %s, got %s instead", expected, got)
	} else {
		return fmt.Sprintf("expected %s to be %s, got %s instead", kind, expected, got)
	}
}
