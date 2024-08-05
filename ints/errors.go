package ints

import (
	"strings"
)

// ErrAt represents an error that occurs at a specific index.
type ErrAt struct {
	// Idx is the index of the error.
	Idx int

	// IdxType is the type of the index.
	IdxType string

	// Reason is the reason for the error.
	Reason error
}

// Error implements the error interface.
//
// Message:
//   - "something went wrong at the <ordinal> <idx_type>" if Reason is nil
//   - "<ordinal> <idx_type> is invalid: <reason>" if Reason is not nil
func (e *ErrAt) Error() string {
	var idx_type string

	if e.IdxType != "" {
		idx_type = e.IdxType
	} else {
		idx_type = "index"
	}

	var builder strings.Builder

	if e.Reason == nil {
		builder.WriteString("something went wrong at the ")
		builder.WriteString(GetOrdinalSuffix(e.Idx))
		builder.WriteRune(' ')
		builder.WriteString(idx_type)
	} else {
		builder.WriteString(GetOrdinalSuffix(e.Idx))
		builder.WriteRune(' ')
		builder.WriteString(idx_type)
		builder.WriteString(" is invalid: ")
		builder.WriteString(e.Reason.Error())
	}

	return builder.String()
}

// NewErrAt creates a new ErrAt error.
//
// Parameters:
//   - idx: The index of the error.
//   - idx_type: The type of the index.
//   - reason: The reason for the error.
//
// Returns:
//   - *ErrAt: A pointer to the newly created ErrAt. Never returns nil.
//
// Empty name will default to "index".
func NewErrAt(idx int, idx_type string, reason error) *ErrAt {
	return &ErrAt{
		Idx:     idx,
		IdxType: idx_type,
		Reason:  reason,
	}
}

// Unwrap is a method that returns the error wrapped by the ErrAt.
//
// Returns:
//   - error: The error wrapped by the ErrAt.
func (e *ErrAt) Unwrap() error {
	return e.Reason
}

// ChangeReason changes the reason for the error.
//
// Parameters:
//   - reason: The new reason for the error.
func (e *ErrAt) ChangeReason(reason error) {
	e.Reason = reason
}

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
func (e *ErrWhileAt) Error() string {
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
		builder.WriteString(GetOrdinalSuffix(e.Idx))
		builder.WriteRune(' ')
		builder.WriteString(idx_type)
	} else {
		builder.WriteString("while ")
		builder.WriteString(operation)
		builder.WriteRune(' ')
		builder.WriteString(GetOrdinalSuffix(e.Idx))
		builder.WriteRune(' ')
		builder.WriteString(idx_type)
		builder.WriteString(": ")
		builder.WriteString(e.Reason.Error())
	}

	return builder.String()
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

// Unwrap is a method that returns the error wrapped by the ErrWhileAt.
//
// Returns:
//   - error: The error wrapped by the ErrWhileAt.
func (e *ErrWhileAt) Unwrap() error {
	return e.Reason
}

// ChangeReason changes the reason for the error.
//
// Parameters:
//   - reason: The new reason for the error.
func (e *ErrWhileAt) ChangeReason(reason error) {
	e.Reason = reason
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
func NewErrInvalidDigit(idx, digit, base int) *ErrAt {
	return NewErrAt(idx+1, "digit", NewErrOutOfBounds(digit, 0, base))
}
