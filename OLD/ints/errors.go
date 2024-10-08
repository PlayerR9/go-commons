package ints

import (
	"fmt"
	"strings"

	gers "github.com/PlayerR9/go-errors"
	gerr "github.com/PlayerR9/go-errors/error"
	"github.com/dustin/go-humanize"
)

// ErrWhileAt represents an error that occurs while doing something at a specific index.
type ErrWhileAt struct {
	// Idx is the index of the error.
	Idx int

	// IdxType is the type of the index.
	IdxType string

	// Operation is the operation being performed.
	Operation string

	// Reason is the reason for the error.
	Reason error
}

// Error implements the error interface.
//
// Message:
// - "an error occurred while <operation> <ordinal> <idx_type>" if Reason is nil
// - "while <operation> <ordinal> <idx_type>: <reason>" if Reason is not nil
func (e ErrWhileAt) Error() string {
	var idx_type string

	if e.IdxType != "" {
		idx_type = e.IdxType
	} else {
		idx_type = "index"
	}

	var operation string

	if e.Operation != "" {
		operation = e.Operation
	} else {
		operation = "doing something"
	}

	var builder strings.Builder

	if e.Reason == nil {
		builder.WriteString("an error occurred while ")
		builder.WriteString(operation)
		builder.WriteRune(' ')
		builder.WriteString(humanize.Ordinal(e.Idx))
		builder.WriteRune(' ')
		builder.WriteString(idx_type)
	} else {
		builder.WriteString("while ")
		builder.WriteString(operation)
		builder.WriteRune(' ')
		builder.WriteString(humanize.Ordinal(e.Idx))
		builder.WriteRune(' ')
		builder.WriteString(idx_type)
		builder.WriteString(": ")
		builder.WriteString(e.Reason.Error())
	}

	return builder.String()
}

// Unwrap implements the errors.Unwrapper interface.
func (e ErrWhileAt) Unwrap() error {
	return e.Reason
}

// ChangeReason implements the errors.Unwrapper interface.
func (e *ErrWhileAt) ChangeReason(reason error) bool {
	if e == nil {
		return false
	}

	e.Reason = reason

	return true
}

// NewErrWhileAt creates a new ErrWhileAt error.
//
// Parameters:
//   - operation: The operation being performed.
//   - idx: The index of the error.
//   - idx_type: The type of the index.
//   - reason: The reason for the error.
//
// Returns:
//   - *ErrWhileAt: A pointer to the newly created ErrWhileAt. Never returns nil.
//
// Empty name will default to "index" and empty operation will default to "doing something".
func NewErrWhileAt(operation string, idx int, idx_type string, reason error) *ErrWhileAt {
	e := &ErrWhileAt{
		Idx:       idx,
		Operation: operation,
		IdxType:   idx_type,
		Reason:    reason,
	}
	return e
}

// NewErrInvalidDigit is just a convenience function for creating an *ErrAt
// with the *ErrOutOfBounds error.
//
// Parameters:
//   - idx: The index of the invalid digit.
//   - digit: The invalid digit.
//   - base: The base number.
//
// Returns:
//   - *ErrAt: A pointer to the newly created ErrAt. Never returns nil.
func NewErrInvalidDigit(idx, digit, base int) *gerr.Err {
	return gers.NewErrAt(humanize.Ordinal(idx+1)+" digit", fmt.Errorf("value of %d is not in the range [0, %d)", digit, base))
}

// NewErrInvalidBase is a convenience function for creating an *errors.ErrInvalidParameter
// with the *errors.ErrGT error.
//
// Parameters:
//   - param_name: The name of the invalid parameter. Defaults to "base" if empty.
//
// Returns:
//   - *errors.ErrInvalidParameter: A pointer to the newly created ErrInvalidParameter. Never returns nil.
func NewErrInvalidBase(param_name string) *gerr.Err {
	if param_name == "" {
		param_name = "base"
	}

	err := gerr.New(gers.BadParameter, fmt.Sprintf("parameter (%q) must be positive", param_name))

	return err
}

// ErrTokenNotFound is a struct that represents an error when a token is not
// found in the content.
type ErrTokenNotFound struct {
	// IsOpening is the type of the token (opening or closing).
	IsOpening bool
}

// Error implements the error interface.
//
// Message: "{Type} token is not in the content"
func (e ErrTokenNotFound) Error() string {
	var builder strings.Builder

	if e.IsOpening {
		builder.WriteString("opening")
	} else {
		builder.WriteString("closing")
	}

	builder.WriteString(" token is not in the content")

	return builder.String()
}

// NewErrTokenNotFound is a constructor of ErrTokenNotFound.
//
// Parameters:
//   - is_opening: The type of the token (opening or closing).
//
// Returns:
//   - *ErrTokenNotFound: A pointer to the newly created error.
func NewErrTokenNotFound(is_opening bool) *ErrTokenNotFound {
	return &ErrTokenNotFound{
		IsOpening: is_opening,
	}
}
