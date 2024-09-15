package slices

import (
	"cmp"
	"slices"
)

// ApplyOnSlice applies a function to each element of a slice.
//
// Parameters:
//   - slice: The slice to apply the function to.
//   - fn: The function to apply.
func ApplyOnSlice[T any](slice []T, fn func(value T) T) {
	if len(slice) == 0 || fn == nil {
		return
	}

	for i := 0; i < len(slice); i++ {
		slice[i] = fn(slice[i])
	}
}

// TryInsert is a helper function that inserts an element into a slice only
// if the element is not already in the slice.
//
// Parameters:
//   - slc: The slice to insert into.
//   - e: The element to insert.
//
// Returns:
//   - []T: The slice with the inserted element.
//
// This function only works if the slice is sorted.
func TryInsert[T cmp.Ordered](slc []T, e T) []T {
	pos, ok := slices.BinarySearch(slc, e)
	if ok {
		return slc
	}

	slc = slices.Insert(slc, pos, e)

	return slc
}
