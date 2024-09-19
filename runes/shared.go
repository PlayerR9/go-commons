package runes

import "io"

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

// Write writes the rune to the writer.
//
// Parameters:
//   - w: the writer.
//   - char: the rune.
//
// Returns:
//   - error: if an error occurred.
//
// Errors:
//   - io.ErrShortWrite if the rune is not fully written or the writer is nil.
//   - any other error returned by the writer.
func Write(w io.Writer, char rune) error {
	if w == nil {
		return io.ErrShortWrite
	}

	data := []byte(string(char))

	n, err := w.Write(data)
	if err != nil {
		return err
	} else if n != len(data) {
		return io.ErrShortWrite
	}

	return nil
}
