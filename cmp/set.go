package cmp

import (
	"cmp"
	"fmt"
	"iter"
	"slices"
	"strings"
)

// SetElem represents an element in a set.
type SetElem interface {
	cmp.Ordered
	fmt.Stringer
}

// Set represents a set of elements.
type Set[T SetElem] struct {
	// elems is the set of elements
	elems []T
}

// String implements the fmt.Stringer interface
//
// Format:
//
//	{elem1, elem2, ...}
func (s Set[T]) String() string {
	elems := make([]string, 0, len(s.elems))
	for _, elem := range s.elems {
		elems = append(elems, elem.String())
	}

	return "{" + strings.Join(elems, ", ") + "}"
}

// NewSet creates a new Set.
//
// Returns:
//   - *Set[T]: The created Set. Never returns nil.
func NewSet[T SetElem]() *Set[T] {
	return &Set[T]{
		elems: make([]T, 0),
	}
}

// Add adds an element to the set. If the element is already in the set, this method does nothing.
//
// Parameters:
//   - elem: The element to add.
//
// Returns:
//   - bool: True if the element was added to the set, false otherwise.
func (s *Set[T]) Add(elem T) bool {
	pos, ok := slices.BinarySearch(s.elems, elem)
	if ok {
		return false
	}

	s.elems = slices.Insert(s.elems, pos, elem)

	return true
}

// Contains checks whether the set contains the given element.
//
// Parameters:
//   - elem: The element to check for.
//
// Returns:
//   - bool: True if the set contains the element, false otherwise.
func (s Set[T]) Contains(elem T) bool {
	_, ok := slices.BinarySearch(s.elems, elem)
	return ok
}

// Equals checks whether the set is equal to another set. Two Sets are said to be equal
// if the intersection of their elements is non-empty.
//
// Parameters:
//   - other: The other set to compare with.
//
// Returns:
//   - bool: True if the sets are equal, false otherwise.
//
// If 'others' is nil, this method returns false.
func (s Set[T]) Equals(other *Set[T]) bool {
	if other == nil {
		return false
	}

	for _, elem := range s.elems {
		if !other.Contains(elem) {
			return false
		}
	}

	return true
}

// IsEmpty checks whether the set is empty.
//
// Returns:
//   - bool: True if the set is empty, false otherwise.
func (s Set[T]) IsEmpty() bool {
	return len(s.elems) == 0
}

// Union adds all elements from another set to the set.
//
// Parameters:
//   - other: The other set to add.
//
// Returns:
//   - int: The number of elements added.
func (s *Set[T]) Union(other *Set[T]) int {
	if other == nil {
		return 0
	}

	var count int

	for _, elem := range other.elems {
		pos, ok := slices.BinarySearch(s.elems, elem)
		if ok {
			continue
		}

		s.elems = slices.Insert(s.elems, pos, elem)
		count++
	}

	return count
}

// All returns an iterator over all elements in the set.
//
// Returns:
//   - iter.Seq[T]: An iterator over all elements in the set.
func (s Set[T]) All() iter.Seq[T] {
	fn := func(yield func(T) bool) {
		for _, elem := range s.elems {
			if !yield(elem) {
				return
			}
		}
	}

	return fn
}

// Slice returns the slice of elements in the set.
//
// Returns:
//   - []T: The slice of elements in the set.
func (s Set[T]) Slice() []T {
	return s.elems
}

// Len returns the number of elements in the set.
//
// Returns:
//   - int: The number of elements in the set.
func (s Set[T]) Len() int {
	return len(s.elems)
}
