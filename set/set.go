package set

import (
	"iter"
	"slices"
)

// Set represents a set of elements.
type Set[T interface {
	// Equals checks whether the given element is equal to the current element.
	// Nil elements are never equal.
	//
	// Parameters:
	//   - other: The other element.
	//
	// Returns:
	//   - bool: True if the given element is equal to the current element, false otherwise.
	Equals(other T) bool
}] struct {
	// elems is the set of elements
	elems []T
}

// NewSet creates a new empty set.
//
// Returns:
//   - *Set[T]: The created set. Never returns nil.
func NewSet[T interface {
	// Equals checks whether the given element is equal to the current element.
	// Nil elements are never equal.
	//
	// Parameters:
	//   - other: The other element.
	//
	// Returns:
	//   - bool: True if the given element is equal to the current element, false otherwise.
	Equals(other T) bool
}]() *Set[T] {
	return &Set[T]{
		elems: make([]T, 0),
	}
}

// NewSetWithItems creates a new set with the given items.
//
// Returns:
//   - *Set[T]: The created set. Never returns nil.
func NewSetWithItems[T interface {
	// Equals checks whether the given element is equal to the current element.
	// Nil elements are never equal.
	//
	// Parameters:
	//   - other: The other element.
	//
	// Returns:
	//   - bool: True if the given element is equal to the current element, false otherwise.
	Equals(other T) bool
}](items []T) *Set[T] {
	unique := make([]T, 0, len(items))

	for _, item := range items {
		if !slices.ContainsFunc(unique, item.Equals) {
			unique = append(unique, item)
		}
	}

	return &Set[T]{
		elems: unique,
	}
}

// IsEmpty checks whether the set is empty.
//
// Returns:
//   - bool: True if the set is empty, false otherwise.
func (s *Set[T]) IsEmpty() bool {
	return len(s.elems) == 0
}

// Size returns the number of elements in the set.
//
// Returns:
//   - int: The number of elements in the set.
func (s *Set[T]) Size() int {
	return len(s.elems)
}

// Add adds an element to the set. If the element is already in the set, this method does nothing.
//
// Parameters:
//   - elem: The element to add.
//
// Returns:
//   - bool: True if the element was added, false otherwise.
func (s *Set[T]) Add(elem T) bool {
	has_element := slices.ContainsFunc(s.elems, elem.Equals)

	if !has_element {
		s.elems = append(s.elems, elem)
	}

	return !has_element
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
		if !slices.ContainsFunc(s.elems, elem.Equals) {
			s.elems = append(s.elems, elem)
			count++
		}
	}

	return count
}

// Clear removes all elements from the set.
func (s *Set[T]) Clear() {
	if s == nil {
		return
	}

	for i := 0; i < len(s.elems); i++ {
		s.elems[i] = *new(T)
	}
	s.elems = s.elems[:0]
}

// All returns an iterator that iterates over all elements in the set.
//
// Returns:
//   - iter.Seq[T]: The iterator. Never returns nil.
func (s *Set[T]) All() iter.Seq[T] {
	var fn func(yield func(T) bool)

	if s == nil {
		fn = func(yield func(T) bool) {}
	} else {
		fn = func(yield func(T) bool) {
			for _, elem := range s.elems {
				if !yield(elem) {
					return
				}
			}
		}
	}

	return fn
}
