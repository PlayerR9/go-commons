package slices

// ApplyOnSlice applies a function to each element of a slice.
//
// Parameters:
//   - slice: The slice to apply the function to.
//   - fn: The function to apply.
func ApplyOnSlice[T any](slice []T, fn func(value T) T) {
	if len(slice) == 0 || fn == nil {
		return
	}

	for i := 0; i < len(slice); i++ {
		slice[i] = fn(slice[i])
	}
}
