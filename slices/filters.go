package slices

// PredicateFilter checks whether an element should be filtered.
//
// Parameters:
//   - elem: The element of a slice.
//
// Returns:
//   - bool: True if the element should be included, false otherwise.
type PredicateFilter[T any] func(elem T) bool

// FilterSlice removes every element that do not satisfy the filter function.
// More specifically, if the 'filter' returns true for a given element, it is kept in the slice.
//
// Parameters:
//   - slice: The slice to filter.
//   - filter: The filter function to use.
//
// Returns:
//   - []T: The filtered slice.
//
// If 'filter' is nil, a nil slice is returned.
//
// NOTES: This function has side-effects, meaning that it changes the original slice.
// To avoid unintended side-effects, you may either want to use the optimized PureSliceFilter
// or just copy the slice before applying the filter.
func FilterSlice[T any](slice []T, filter PredicateFilter[T]) []T {
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

// FilterZeroValues removes every element that is equal to zero.
// More specifically:
//   - pointers, slice, map, etc.: Checks if the pointer is nil.
//   - string: Checks if the string is empty.
//   - int, int8, ..., uint, rune, ...: Checks if the value is 0.
//   - struct{}: Checks if the struct is the zero struct.
//   - and so on.
//
// Parameters:
//   - slice: The slice to filter.
//
// Returns:
//   - []T: The filtered slice.
//
// NOTES: This function has side-effects, meaning that it changes the original slice.
// To avoid unintended side-effects, you may either want to use the optimized PureFilterZeroValues
// or just copy the slice before applying the filter.
func FilterZeroValues[T comparable](slice []T) []T {
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

// GroupByFilter splits the slice into two slices according to the filter function.
//
// Parameters:
//   - slice: The slice to split.
//   - filter: The filter function to use.
//
// Returns:
//   - []T: The elements that satisfy the filter function.
//   - []T: The elements that do not satisfy the filter function.
//
// If 'filter' is nil, then the original slice is returned as the second slice while
// the first is nil.
//
// NOTES: This function has side-effects, meaning that it changes the original slice.
// To avoid unintended side-effects, you may either want to use the optimized PureGroupByFilter
// or just copy the slice before applying the filter.
func GroupByFilter[T any](slice []T, filter PredicateFilter[T]) ([]T, []T) {
	if len(slice) == 0 || filter == nil {
		return nil, slice
	}

	failed := make([]T, 0, len(slice)/2)
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

// SuccessOrSlice returns the successful elements of a slice and true. However,
// if no elements are successful, it returns the original slice and false.
//
// Parameters:
//   - slice: The slice to filter.
//   - filter: The filter function to use.
//
// Returns:
//   - []T: The filtered slice.
//   - bool: true if there are successful elements, otherwise false.
//
// As a special case:
//   - If the slice is empty, the function returns an empty slice and true.
//   - If the filter is nil, the function returns the original slice and false.
//
// NOTES: This function has side-effects, meaning that it changes the original slice.
// To avoid unintended side-effects, you may either want to use the optimized PureSuccessOrSlice
// or just copy the slice before applying the filter.
func SuccessOrSlice[T any](slice []T, filter PredicateFilter[T]) ([]T, bool) {
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

// Unique removes duplicate elements from a slice. Order is guaranteed to be preserved.
//
// Parameters:
//   - slice: The slice to remove duplicates from.
//
// Returns:
//   - []T: The filtered slice.
//
// NOTES: This function has side-effects, meaning that it changes the original slice.
// To avoid unintended side-effects, you may either want to use the optimized PureUnique
// or just copy the slice before applying the filter.
//
// If order is not important, consider using PureUnique instead. Otherwise, copy the slice
// before applying the filter.
func Unique[T comparable](slice []T) []T {
	if len(slice) == 0 {
		return nil
	}

	for i := 0; i < len(slice)-1; i++ {
		elem := slice[i]

		top := i + 1

		for j := i + 1; j < len(slice); j++ {
			if slice[j] != elem {
				slice[top] = slice[j]
				top++
			}
		}

		slice = slice[:top:top]
	}

	return slice
}
