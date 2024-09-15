package slices

// PredicateFilter is a type that defines a slice filter function.
//
// Parameters:
//   - T: The type of the elements in the slice.
//
// Returns:
//   - bool: True if the element satisfies the filter function, otherwise false.
type PredicateFilter[T any] func(T) bool

// SliceFilter is a function that iterates over the slice and applies the filter
// function to each element.
//
// Parameters:
//   - S: slice of elements.
//   - filter: function that takes an element and returns a bool.
//
// Returns:
//   - []T: slice of elements that satisfy the filter function.
//
// Behavior:
//   - If S is empty, the function returns a nil slice.
//   - If S has only one element and it satisfies the filter function, the function
//     returns a slice with that element. Otherwise, it returns a nil slice.
//   - An element is said to satisfy the filter function if the function returns true
//     when applied to the element.
//   - If the filter function is nil, the function returns the original slice.
func SliceFilter[T any](S []T, filter PredicateFilter[T]) []T {
	if len(S) == 0 {
		return nil
	} else if filter == nil {
		return S
	}

	var top int

	for i := 0; i < len(S); i++ {
		ok := filter(S[i])
		if ok {
			S[top] = S[i]
			top++
		}
	}

	return S[:top:top]
}

// FilterNilValues is a function that iterates over the slice and removes the
// nil elements.
//
// Parameters:
//   - S: slice of elements.
//
// Returns:
//   - []*T: slice of elements that satisfy the filter function.
//
// Behavior:
//   - If S is empty, the function returns a nil slice.
func FilterNilValues[T any](S []*T) []*T {
	if len(S) == 0 {
		return nil
	}

	var top int

	for i := 0; i < len(S); i++ {
		if S[i] != nil {
			S[top] = S[i]
			top++
		}
	}

	return S[:top:top]
}
