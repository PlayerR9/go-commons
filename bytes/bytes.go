package bytes

import (
	"index/suffixarray"
	"slices"
)

// ReverseSearch searches for the last occurrence of a byte in a byte slice.
//
// Parameters:
//   - data: the byte slice to search in.
//   - from: the index to start the search from. If greater than or equal to the length of the byte slice,
//     it is treated as the length of the byte slice minus 1.
//   - sep: the byte to search for.
//
// Returns:
//   - int: the index of the last occurrence of the byte in the byte slice, or -1 if not found.
func ReverseSearch(data []byte, from int, sep []byte) int {
	if len(data) == 0 || from < 0 || len(sep) == 0 {
		return -1
	}

	len_data := len(data)

	if from >= len_data {
		from = len_data
	}

	sub_data := data[:from+1]
	slices.Reverse(sub_data)

	rev_sep := make([]byte, len(sep))
	copy(rev_sep, sep)

	slices.Reverse(rev_sep)

	idx := suffixarray.New(sub_data)

	offsets := idx.Lookup(rev_sep, 1)

	if len(offsets) == 0 {
		return -1
	}

	return offsets[0]
}

// ForwardSearch searches for the first occurrence of a byte in a byte slice.
//
// Parameters:
//   - data: the byte slice to search in.
//   - from: the index to start the search from. If negative, it is treated as 0.
//   - sep: the byte to search for.
//
// Returns:
//   - int: the index of the first occurrence of the byte in the byte slice, or -1 if not found.
func ForwardSearch(data []byte, from int, sep []byte) int {
	if len(data) == 0 || len(sep) == 0 || from >= len(data) {
		return -1
	}

	if from < 0 {
		from = 0
	}

	idx := suffixarray.New(data[from:])

	offsets := idx.Lookup(sep, 1)
	if len(offsets) == 0 {
		return -1
	}

	return offsets[0]
}
