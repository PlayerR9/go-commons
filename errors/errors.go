package errors

import (
	"fmt"
	"strconv"
	"strings"

	gcers "github.com/PlayerR9/go-commons/errors/error"
)

//go:generate stringer -type=ErrorCode

type ErrorCode int

const (
	// BadParameter occurs when a parameter is invalid or is not
	// valid for some reason. For example, a nil pointer when nil
	// pointers are not allowed.
	BadParameter ErrorCode = iota

	// InvalidUsage occurs when users call a function without
	// proper setups or preconditions.
	InvalidUsage

	// FailFix occurs when a struct cannot be fixed or resolved
	// due to an invalid internal state.
	FailFix
)

// NewErrInvalidParameter creates a new error.Err[ErrorCode] error.
//
// Parameters:
//   - message: The message of the error.
//
// Returns:
//   - *error.Err[ErrorCode]: The new error. Never returns nil.
func NewErrInvalidParameter(message string) *gcers.Err[ErrorCode] {
	err := gcers.NewErr(gcers.FATAL, BadParameter, message)

	return err
}

// NewErrNilParameter creates a new error.Err[ErrorCode] error.
//
// Parameters:
//   - parameter: the name of the invalid parameter.
//
// Returns:
//   - *error.Err[ErrorCode]: The new error. Never returns nil.
func NewErrNilParameter(parameter string) *gcers.Err[ErrorCode] {
	err := gcers.NewErr(gcers.FATAL, BadParameter, fmt.Sprintf("parameter (%q) cannot be nil", parameter))

	return err
}

// NewErrInvalidUsage creates a new error.Err[ErrorCode] error.
//
// Parameters:
//   - message: The message of the error.
//   - usage: The usage/suggestion to solve the problem.
//
// Returns:
//   - *error.Err[ErrorCode]: The new error. Never returns nil.
func NewErrInvalidUsage(message string, usage string) *gcers.Err[ErrorCode] {
	err := gcers.NewErr(gcers.FATAL, InvalidUsage, message)

	err.AddSuggestion(usage)

	return err
}

// NewErrFix creates a new error.Err[ErrorCode] error.
//
// Parameters:
//   - message: The message of the error.
//
// Returns:
//   - *error.Err[ErrorCode]: The new error. Never returns nil.
func NewErrFix(message string) *gcers.Err[ErrorCode] {
	err := gcers.NewErr(gcers.FATAL, FailFix, message)

	return err
}

////////////////////////////////////////////////////////////////////////////////

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
func (e ErrAt) Error() string {
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

// Unwrap implements the errors.Unwrapper interface.
func (e ErrAt) Unwrap() error {
	return e.Reason
}

// ChangeReason implements the errors.Unwrapper interface.
func (e *ErrAt) ChangeReason(new_reason error) {
	if e == nil {
		return
	}

	e.Reason = new_reason
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

// ErrAfter represents an error that occurs after something.
type ErrAfter struct {
	// After is the element that was processed before the error occurred.
	After string

	// Reason is the reason for the error.
	Reason error

	// ShoulQuote is whether the error should be quoted.
	ShouldQuote bool
}

// Error implements the error interface.
//
// Message: "after {after}: {reason}".
//
// However, if the reason is nil, the message is "something went wrong after {after}"
// instead.
func (e ErrAfter) Error() string {
	var after string

	if e.After == "" {
		after = "something"
	} else if e.ShouldQuote {
		after = strconv.Quote(e.After)
	} else {
		after = e.After
	}

	var builder strings.Builder

	if e.Reason == nil {
		builder.WriteString("something went wrong after ")
		builder.WriteString(after)
	} else {
		builder.WriteString("after ")
		builder.WriteString(after)
		builder.WriteString(": ")
		builder.WriteString(e.Reason.Error())
	}

	return builder.String()
}

// Unwrap implements the errors.Unwrapper interface.
func (e ErrAfter) Unwrap() error {
	return e.Reason
}

// ChangeReason implements the errors.Unwrapper interface.
func (e *ErrAfter) ChangeReason(reason error) {
	if e == nil {
		return
	}

	e.Reason = reason
}

// NewErrAfter creates a new ErrAfter error.
//
// Parameters:
//   - after: The element that was processed before the error occurred.
//   - reason: The reason for the error.
//   - should_quote: Whether the error should be quoted.
//
// Returns:
//   - *ErrAfter: A pointer to the new ErrAfter error. Never returns nil.
func NewErrAfter(after string, reason error, should_quote bool) *ErrAfter {
	return &ErrAfter{
		After:       after,
		Reason:      reason,
		ShouldQuote: should_quote,
	}
}
