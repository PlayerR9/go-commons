package slices

import "slices"

// Pointer is an interface for a pointer.
type Pointer interface {
	// IsNil checks if the pointer is nil.
	//
	// Returns:
	//   - bool: True if the pointer is nil, false otherwise.
	IsNil() bool
}

// FilterNonNil removes all nil elements from the given slice of elements.
//
// This function assumes that the receiver is not nil and that the receiver is
// not empty.
//
// Returns:
//   - []T: The slice of elements with all nils removed.
func FilterNonNil[T Pointer](elems []T) []T {
	if len(elems) == 0 {
		return nil
	}

	var top int

	for i := 0; i < len(elems); i++ {
		if !elems[i].IsNil() {
			elems[top] = elems[i]
			top++
		}
	}

	return elems[:top:top]
}

// UniqueEquals removes all duplicate elements from the given slice of elements using the Equals method.
//
// Parameters:
//   - elems: The slice of elements.
//
// Returns:
//   - []T: The slice of elements with all duplicates removed.
func UniqueEquals[T interface {
	Equals(other T) bool
}](elems []T) []T {
	if len(elems) == 0 {
		return nil
	} else if len(elems) == 1 {
		return elems
	}

	for i := 0; i < len(elems)-1; i++ {
		elem := elems[i]

		top := i + 1

		for j := i + 1; j < len(elems); j++ {
			other := elems[j]

			if !elem.Equals(other) {
				elems[top] = other
				top++
			}
		}

		elems = elems[:top:top]
	}

	return elems
}

// MergeEquals removes all duplicate elements from the given slices of elements using the Equals method.
//
// Parameters:
//   - elems1: The first slice of elements.
//   - elems2: The second slice of elements.
//
// Returns:
//   - []T: The slice of elements with all duplicates removed.
func MergeEquals[T interface {
	Equals(other T) bool
}](elems1, elems2 []T) []T {
	elems1 = UniqueEquals(elems1)
	elems2 = UniqueEquals(elems2)

	elems := make([]T, len(elems1), len(elems1)+len(elems2))
	copy(elems, elems1)

	for _, elem := range elems2 {
		ok := slices.ContainsFunc(elems1, elem.Equals)
		if !ok {
			elems = append(elems, elem)
		}
	}

	return elems[:len(elems):len(elems)]
}
