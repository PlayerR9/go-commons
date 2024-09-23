package slices

// PureSliceFilter is the same as SliceFilter, but without any side-effects.
//
// Parameters:
//   - slice: slice of elements.
//   - filter: function that takes an element and returns a bool.
//
// Returns:
//   - []T: slice of elements that satisfy the filter function.
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

// PureFilterNilValues is the same as FilterNilValues, but without any side-effects.
func PureFilterNilValues[T comparable](slice []T) []T {
	if len(slice) == 0 {
		return nil
	}

	zero := *new(T)

	slice_copy := make([]T, 0, len(slice))

	for i := 0; i < len(slice); i++ {
		if slice[i] != zero {
			slice_copy = append(slice_copy, slice[i])
		}
	}

	return slice_copy[:len(slice_copy):len(slice_copy)]
}
