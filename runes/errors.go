package runes

import (
	"strconv"
)

// ErrInvalidUTF8Encoding is an error type for invalid UTF-8 encoding.
type ErrInvalidUTF8Encoding struct {
	// At is the index of the invalid UTF-8 encoding.
	At int
}

// Error implements the error interface.
//
// Message:
//
//	"invalid UTF-8 encoding at index {At}"
func (e ErrInvalidUTF8Encoding) Error() string {
	return "invalid UTF-8 encoding at index " + strconv.Itoa(e.At)
}

// NewErrInvalidUTF8Encoding creates a new ErrInvalidUTF8Encoding error.
//
// Parameters:
//   - at: The index of the invalid UTF-8 encoding.
//
// Returns:
//   - *ErrInvalidUTF8Encoding: A pointer to the newly created error.
func NewErrInvalidUTF8Encoding(at int) *ErrInvalidUTF8Encoding {
	return &ErrInvalidUTF8Encoding{
		At: at,
	}
}
