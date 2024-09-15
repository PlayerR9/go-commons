package strings

import (
	"strconv"
)

// SliceOfInts is a function that returns a slice of strings
// from a slice of integers.
//
// Parameters:
//   - values: The values to convert to a slice of strings.
//
// Returns:
//   - []string: The slice of strings.
func SliceOfInts(values []int) []string {
	if len(values) == 0 {
		return nil
	}

	elems := make([]string, 0, len(values))

	for _, value := range values {
		elems = append(elems, strconv.Itoa(value))
	}

	return elems
}

// SliceOfBytes is a function that returns a slice of strings
// from a slice of bytes.
//
// Parameters:
//   - values: The values to convert to a slice of strings.
//
// Returns:
//   - []string: The slice of strings.
func SliceOfBytes(values [][]byte) []string {
	if len(values) == 0 {
		return nil
	}

	elems := make([]string, 0, len(values))

	for _, value := range values {
		elems = append(elems, string(value))
	}

	return elems
}

// SliceOfErrors is a function that returns a slice of strings
// from a slice of errors.
//
// Parameters:
//   - values: The values to convert to a slice of strings.
//
// Returns:
//   - []string: The slice of strings.
func SliceOfErrors(values []error) []string {
	if len(values) == 0 {
		return nil
	}

	elems := make([]string, 0, len(values))

	for _, value := range values {
		elems = append(elems, value.Error())
	}

	return elems
}
