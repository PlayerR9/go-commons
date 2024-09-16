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
	if len(slice) == 0 {
		return nil
	} else if filter == nil {
		return slice
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

// PureSliceFilter is the same as SliceFilter, but without any side-effects.
func PureSliceFilter[T any](slice []T, filter PredicateFilter[T]) []T {
	if len(slice) == 0 {
		return nil
	} else if filter == nil {
		return slice
	}

	slice_copy := make([]T, 0, len(slice))

	for i := 0; i < len(slice); i++ {
		ok := filter(slice[i])
		if ok {
			slice_copy = append(slice_copy, slice[i])
		}
	}

	return slice_copy[:len(slice_copy):len(slice_copy)]
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
func FilterNilValues[T any](slice []*T) []*T {
	if len(slice) == 0 {
		return nil
	}

	var top int

	for i := 0; i < len(slice); i++ {
		if slice[i] != nil {
			slice[top] = slice[i]
			top++
		}
	}

	return slice[:top:top]
}

// PureFilterNilValues is the same as FilterNilValues, but without any side-effects.
func PureFilterNilValues[T any](slice []*T) []*T {
	if len(slice) == 0 {
		return nil
	}

	slice_copy := make([]*T, 0, len(slice))

	for i := 0; i < len(slice); i++ {
		if slice[i] != nil {
			slice_copy = append(slice_copy, slice[i])
		}
	}

	return slice_copy[:len(slice_copy):len(slice_copy)]
}
