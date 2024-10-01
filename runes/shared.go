package runes

import "io"

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
