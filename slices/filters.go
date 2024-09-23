package slices

// PredicateFilter is a function that checks whether an element should be filtered.
//
// Parameters:
//   - elem: The element of a slice.
//
// Returns:
//   - bool: True if the element should be included, false otherwise.
type PredicateFilter[T any] func(elem T) bool

// SliceFilter is a function that iterates over the slice and applies the filter
// function to each element.
//
// Parameters:
//   - slice: slice of elements.
//   - filter: function that takes an element and returns a bool.
//
// Returns:
//   - []T: slice of elements that satisfy the filter function.
//
// Behavior:
//   - If 'slice' is empty, the function returns a nil slice.
//   - If 'slice' has only one element and it satisfies the filter function, the function
//     returns a slice with that element. Otherwise, it returns a nil slice.
//   - An element is said to satisfy the filter function if the function returns true
//     when applied to the element.
//   - If the filter function is nil, the function returns the original slice.
//   - This function has side-effects on 'slice'.
func SliceFilter[T any](slice []T, filter PredicateFilter[T]) []T {
	if len(slice) == 0 || filter == nil {
		return nil
	}

	var top int

	for i := 0; i < len(slice); i++ {
		ok := filter(slice[i])
		if ok {
			slice[top] = slice[i]
			top++
		}
	}

	return slice[:top:top]
}

// FilterNilValues is a function that iterates over the slice and removes the
// nil elements.
//
// Parameters:
//   - slice: slice of elements.
//
// Returns:
//   - []*T: slice of elements that satisfy the filter function.
//
// Behavior:
//   - If 'slice' is empty, the function returns a nil slice.
//   - This function has side-effects on 'slice'.
func FilterNilValues[T comparable](slice []T) []T {
	if len(slice) == 0 {
		return nil
	}

	zero := *new(T)
	var top int

	for i := 0; i < len(slice); i++ {
		if slice[i] != zero {
			slice[top] = slice[i]
			top++
		}
	}

	return slice[:top:top]
}

// SFSeparate is a function that iterates over the slice and applies the filter
// function to each element. The returned slices contain the elements that
// satisfy and do not satisfy the filter function.
//
// Parameters:
//   - S: slice of elements.
//   - filter: function that takes an element and returns a bool.
//
// Returns:
//   - []T: slice of elements that satisfy the filter function.
//   - []T: slice of elements that do not satisfy the filter function.
//
// Behavior:
//   - If S is empty, the function returns two empty slices.
func SFSeparate[T any](slice []T, filter PredicateFilter[T]) ([]T, []T) {
	if len(slice) == 0 || filter == nil {
		return nil, slice
	}

	var failed []T
	var top int

	for i := 0; i < len(slice); i++ {
		ok := filter(slice[i])
		if ok {
			slice[top] = slice[i]
			top++
		} else {
			failed = append(failed, slice[i])
		}
	}

	return slice[:top:top], failed[:len(failed):len(failed)]
}

// SFSeparateEarly is a variant of SFSeparate that returns all successful elements.
// If there are none, it returns the original slice and false.
//
// Parameters:
//   - S: slice of elements.
//   - filter: function that takes an element and returns a bool.
//
// Returns:
//   - []T: slice of elements that satisfy the filter function or the original slice.
//   - bool: true if there are successful elements, otherwise false.
//
// Behavior:
//   - If S is empty, the function returns an empty slice and true.
func SFSeparateEarly[T any](slice []T, filter PredicateFilter[T]) ([]T, bool) {
	if len(slice) == 0 {
		return nil, true
	} else if filter == nil {
		return slice, false
	}

	var top int

	for i := 0; i < len(slice); i++ {
		ok := filter(slice[i])
		if ok {
			slice[top] = slice[i]
			top++
		}
	}

	if top == 0 {
		return slice, false
	} else {
		return slice[:top:top], true
	}
}
