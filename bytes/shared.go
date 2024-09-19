package bytes

import (
	"index/suffixarray"
	"io"
)

// filter_equals returns the indices of the other in the data.
//
// Parameters:
//   - indices: The indices.
//   - data: The data.
//   - other: The other value.
//   - offset: The offset to start the search from.
//
// Returns:
//   - []int: The indices.
func filter_equals(indices []int, data []byte, other byte, offset int) []int {
	var top int

	for i := 0; i < len(indices); i++ {
		idx := indices[i]

		if data[idx+offset] == other {
			indices[top] = idx
			top++
		}
	}

	indices = indices[:top]

	return indices
}

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
func IndicesOf(data []byte, sep []byte, exclude_sep bool) []int {
	if len(data) == 0 || len(sep) == 0 {
		return nil
	}

	idx := suffixarray.New(data)

	offsets := idx.Lookup(sep, -1)
	if len(offsets) == 0 {
		return nil
	}

	if !exclude_sep {
		return offsets
	}

	sep_len := len(sep)

	for i := 0; i < len(offsets); i++ {
		offsets[i] += sep_len
	}

	return offsets
}

// Write writes the data to the writer.
//
// Parameters:
//   - w: the writer.
//   - data: the data.
//
// Returns:
//   - error: if an error occurred.
//
// Errors:
//   - io.ErrShortWrite if the data is not fully written or the writer is nil.
//   - any other error returned by the writer.
func Write(w io.Writer, data []byte) error {
	if len(data) == 0 {
		return nil
	} else if w == nil {
		return io.ErrShortWrite
	}

	n, err := w.Write(data)
	if err != nil {
		return err
	} else if n != len(data) {
		return io.ErrShortWrite
	}

	return nil
}
