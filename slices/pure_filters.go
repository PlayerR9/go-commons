package slices

// PureFilterSlice is the same as SliceFilter, but without any side-effects.
//
// Parameters:
//   - slice: The slice to filter.
//   - filter: The filter function to use.
//
// Returns:
//   - []T: The filtered slice.
func PureFilterSlice[T any](slice []T, filter PredicateFilter[T]) []T {
	if len(slice) == 0 || filter == nil {
		return nil
	}

	result := make([]T, 0, len(slice))

	for i := 0; i < len(slice); i++ {
		ok := filter(slice[i])
		if !ok {
			result = append(result, slice[i])
		}
	}

	return result[:len(result):len(result)]
}

// PureFilterZeroValues is the same as FilterZeroValues, but without any side-effects.
//
// Parameters:
//   - slice: The slice to filter.
//
// Returns:
//   - []T: The filtered slice.
func PureFilterZeroValues[T comparable](slice []T) []T {
	if len(slice) == 0 {
		return nil
	}

	zero := *new(T)
	result := make([]T, 0, len(slice))

	for i := 0; i < len(slice); i++ {
		if slice[i] != zero {
			result = append(result, slice[i])
		}
	}

	return result[:len(result):len(result)]
}

// PureGroupByFilter is the same as GroupByFilter, but without any side-effects.
//
// Parameters:
//   - slice: The slice to split.
//   - filter: The filter function to use.
//
// Returns:
//   - []T: The elements that satisfy the filter function.
//   - []T: The elements that do not satisfy the filter function.
func PureGroupByFilter[T any](slice []T, filter PredicateFilter[T]) ([]T, []T) {
	if len(slice) == 0 || filter == nil {
		return nil, slice
	}

	result := make([]T, 0, len(slice)/2)
	failed := make([]T, 0, len(slice)/2)

	for i := 0; i < len(slice); i++ {
		ok := filter(slice[i])
		if ok {
			result = append(result, slice[i])
		} else {
			failed = append(failed, slice[i])
		}
	}

	return result[:len(result):len(result)], failed[:len(failed):len(failed)]
}

// PureSuccessOrSlice is the same as SuccessOrSlice, but without any side-effects.
//
// Parameters:
//   - slice: The slice to filter.
//   - filter: The filter function to use.
//
// Returns:
//   - []T: The filtered slice.
//   - bool: true if there are successful elements, otherwise false.
func PureSuccessOrSlice[T any](slice []T, filter PredicateFilter[T]) ([]T, bool) {
	if len(slice) == 0 {
		return nil, true
	} else if filter == nil {
		return slice, false
	}

	result := make([]T, 0, len(slice))

	for i := 0; i < len(slice); i++ {
		ok := filter(slice[i])
		if ok {
			result = append(result, slice[i])
		}
	}

	if len(result) == 0 {
		return slice, false
	} else {
		return result[:len(result):len(result)], true
	}
}

// PureUnique is like Unique, but without any side-effects. Yet, results are not guaranteed to be
// in the same order as in the original slice.
//
// Parameters:
//   - slice: The slice to remove duplicates from.
//
// Returns:
//   - []T: The filtered slice.
func PureUnique[T comparable](slice []T) []T {
	if len(slice) == 0 {
		return nil
	}

	seen := make(map[T]interface{})

	for i := 0; i < len(slice); i++ {
		_, ok := seen[slice[i]]
		if !ok {
			seen[slice[i]] = struct{}{}
		}
	}

	result := make([]T, 0, len(seen))

	for k := range seen {
		result = append(result, k)
	}

	return result
}
