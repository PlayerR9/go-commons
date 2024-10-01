package runes

import (
	"strconv"
	"strings"
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

// ErrUnexpectedChar is an error that occurs when an unexpected character is encountered.
type ErrUnexpectedChar struct {
	// Expected is the expected character.
	Expecteds []rune

	// Previous is the previous character.
	Previous rune

	// Got is the current character.
	Got *rune
}

// Error implements the error interface.
//
// Message:
//
//	"expected {expected} after {previous}, got {got} instead".
func (e ErrUnexpectedChar) Error() string {
	var got string

	if e.Got == nil {
		got = "nothing"
	} else {
		got = strconv.QuoteRune(*e.Got)
	}

	var builder strings.Builder

	builder.WriteString("expected ")

	if len(e.Expecteds) == 0 {
		builder.WriteString("nothing")
	} else {
		builder.WriteString(EitherOrString(e.Expecteds, true))
	}

	builder.WriteString(" after ")
	builder.WriteString(strconv.QuoteRune(e.Previous))
	builder.WriteString(", got ")
	builder.WriteString(got)
	builder.WriteString(" instead")

	return builder.String()
}

// NewErrUnexpectedChar creates a new ErrUnexpectedChar error.
//
// Parameters:
//   - previous: the previous character.
//   - expecteds: the expected characters.
//   - got: the current character.
//
// Returns:
//   - *ErrUnexpectedChar: the error. Never returns nil.
func NewErrUnexpectedChar(previous rune, expecteds []rune, got *rune) *ErrUnexpectedChar {
	return &ErrUnexpectedChar{
		Expecteds: expecteds,
		Previous:  previous,
		Got:       got,
	}
}
