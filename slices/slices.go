package slices

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
