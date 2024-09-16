package slices

// Builder is a slice builder.
type Builder[T any] struct {
	// slice is the slice.
	slice []T
}

// Append appends a value to the slice. Does nothing if the receiver is nil.
//
// Parameters:
//   - v: the value to append.
func (sb *Builder[T]) Append(v T) {
	if sb == nil {
		return
	}

	sb.slice = append(sb.slice, v)
}

// Build builds the slice.
//
// Returns:
//   - []T: the slice.
func (sb Builder[T]) Build() []T {
	slice := make([]T, len(sb.slice))
	copy(slice, sb.slice)

	return slice
}

// Reset resets the slice.
func (sb *Builder[T]) Reset() {
	if sb == nil {
		return
	}

	if len(sb.slice) > 0 {
		for i := range sb.slice {
			sb.slice[i] = *new(T)
		}

		sb.slice = sb.slice[:0]
	}
}
