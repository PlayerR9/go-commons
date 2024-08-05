package math

import "fmt"

// ErrInvalidBase represents an error when the base is invalid.
type ErrInvalidBase struct{}

// Error implements the error interface.
//
// Message: "base must be positive"
func (e *ErrInvalidBase) Error() string {
	return "base must be positive"
}

// NewErrInvalidBase creates a new ErrInvalidBase error.
//
// Returns:
//   - *ErrInvalidBase: A pointer to the newly created ErrInvalidBase. Never returns nil.
func NewErrInvalidBase() *ErrInvalidBase {
	return &ErrInvalidBase{}
}

// ErrInvalidDigit represents an error when a digit is invalid.
type ErrInvalidDigit struct {
	// Idx is the index of the invalid digit.
	Idx int

	// Digit is the invalid digit.
	Digit int

	// Base is the base number.
	Base int
}

// Error implements the error interface.
//
// Message: "digit at index %d (%d) is not in the range [0, %d]"
func (e *ErrInvalidDigit) Error() string {
	return fmt.Sprintf("digit at index %d (%d) is not in the range [0, %d]", e.Idx, e.Digit, e.Base-1)
}

// NewErrInvalidDigit creates a new ErrInvalidDigit error.
//
// Parameters:
//   - idx: The index of the invalid digit.
//   - digit: The invalid digit.
//   - base: The base number.
//
// Returns:
//   - *ErrInvalidDigit: A pointer to the newly created ErrInvalidDigit. Never returns nil.
func NewErrInvalidDigit(idx int, digit int, base int) *ErrInvalidDigit {
	return &ErrInvalidDigit{
		Idx:   idx,
		Digit: digit,
		Base:  base,
	}
}
