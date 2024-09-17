package runes

// Indices returns the indices of the separator in the data.
//
// Parameters:
//   - data: The data.
//   - sep: The separator.
//   - exclude_sep: Whether the separator is inclusive. If set to true, the indices will point to the character right after the
//     separator. Otherwise, the indices will point to the separator itself.
//
// Returns:
//   - []int: The indices.
func IndicesOf(data []rune, sep rune, exclude_sep bool) []int {
	if len(data) == 0 {
		return nil
	}

	var indices []int

	for i := 0; i < len(data); i++ {
		if data[i] == sep {
			indices = append(indices, i)
		}
	}

	if len(indices) == 0 {
		return nil
	}

	if exclude_sep {
		for i := 0; i < len(indices); i++ {
			indices[i] += 1
		}
	}

	return indices
}
