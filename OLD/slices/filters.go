package slices

import (
	gcslc "github.com/PlayerR9/go-commons/slices"
)

// Intersect returns a PredicateFilter function that checks if an element
// satisfies all the PredicateFilter functions in funcs.
//
// Parameters:
//   - funcs: A slice of PredicateFilter functions.
//
// Returns:
//   - PredicateFilter: A PredicateFilter function that checks if a element satisfies
//     all the PredicateFilter functions in funcs.
//
// Behavior:
//   - If no filter functions are provided, then all elements are considered to satisfy
//     the filter function.
//   - It returns false as soon as it finds a function in funcs that the element
//     does not satisfy.
func Intersect[T any](funcs ...gcslc.PredicateFilter[T]) gcslc.PredicateFilter[T] {
	if len(funcs) == 0 {
		return func(elem T) bool { return true }
	}

	return func(elem T) bool {
		for _, f := range funcs {
			ok := f(elem)
			if !ok {
				return false
			}
		}

		return true
	}
}

// ParallelIntersect returns a PredicateFilter function that checks if an element
// satisfies all the PredicateFilter functions in funcs concurrently.
//
// Parameters:
//   - funcs: A slice of PredicateFilter functions.
//
// Returns:
//   - PredicateFilter: A PredicateFilter function that checks if a element satisfies
//     all the PredicateFilter functions in funcs.
//
// Behavior:
//   - If no filter functions are provided, then all elements are considered to satisfy
//     the filter function.
//   - It returns false as soon as it finds a function in funcs that the element
//     does not satisfy.
func ParallelIntersect[T any](funcs ...gcslc.PredicateFilter[T]) gcslc.PredicateFilter[T] {
	if len(funcs) == 0 {
		return func(elem T) bool { return true }
	}

	return func(elem T) bool {
		resultChan := make(chan bool, len(funcs))

		for _, f := range funcs {
			go func(f gcslc.PredicateFilter[T]) {
				resultChan <- f(elem)
			}(f)
		}

		for range funcs {
			if !<-resultChan {
				return false
			}
		}

		return true
	}
}

// Union returns a PredicateFilter function that checks if an element
// satisfies at least one of the PredicateFilter functions in funcs.
//
// Parameters:
//   - funcs: A slice of PredicateFilter functions.
//
// Returns:
//   - PredicateFilter: A PredicateFilter function that checks if a element satisfies
//     at least one of the PredicateFilter functions in funcs.
//
// Behavior:
//   - If no filter functions are provided, then no elements are considered to satisfy
//     the filter function.
//   - It returns true as soon as it finds a function in funcs that the element
//     satisfies.
func Union[T any](funcs ...gcslc.PredicateFilter[T]) gcslc.PredicateFilter[T] {
	if len(funcs) == 0 {
		return func(elem T) bool { return false }
	}

	return func(elem T) bool {
		for _, f := range funcs {
			ok := f(elem)
			if ok {
				return true
			}
		}

		return false
	}
}

// ParallelUnion returns a PredicateFilter function that checks if an element
// satisfies at least one of the PredicateFilter functions in funcs concurrently.
//
// Parameters:
//   - funcs: A slice of PredicateFilter functions.
//
// Returns:
//   - PredicateFilter: A PredicateFilter function that checks if a element satisfies
//     at least one of the PredicateFilter functions in funcs.
//
// Behavior:
//   - If no filter functions are provided, then no elements are considered to satisfy
//     the filter function.
//   - It returns true as soon as it finds a function in funcs that the element
//     satisfies.
func ParallelUnion[T any](funcs ...gcslc.PredicateFilter[T]) gcslc.PredicateFilter[T] {
	if len(funcs) == 0 {
		return func(elem T) bool { return false }
	}

	return func(elem T) bool {
		resultChan := make(chan bool, len(funcs))

		for _, f := range funcs {
			go func(f gcslc.PredicateFilter[T]) {
				resultChan <- f(elem)
			}(f)
		}

		for range funcs {
			if <-resultChan {
				return true
			}
		}

		return false
	}
}

// FilterNilPredicates is a function that iterates over the slice and removes the
// nil predicate functions.
//
// Parameters:
//   - S: slice of predicate functions.
//
// Returns:
//   - []PredicateFilter: slice of predicate functions that are not nil.
//
// Behavior:
//   - If S is empty, the function returns a nil slice.
func FilterNilPredicates[T any](S []gcslc.PredicateFilter[T]) []gcslc.PredicateFilter[T] {
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

// RemoveEmpty is a function that removes the empty elements from a slice.
//
// Parameters:
//   - elems: The slice of elements.
//
// Returns:
//   - []T: The slice of elements without the empty elements.
func RemoveEmpty[T comparable](elems []T) []T {
	var top int

	for i := 0; i < len(elems); i++ {
		empty := *new(T)
		if elems[i] != empty {
			elems[top] = elems[i]
			top++
		}
	}

	return elems[:top:top]
}
