package errors

import (
	"strconv"

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

	// OperationFail occurs when an operation cannot be completed
	// due to an internal error.
	OperationFail
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
	msg := "parameter (" + strconv.Quote(parameter) + ") must not be nil"

	err := gcers.NewErr(gcers.FATAL, BadParameter, msg)

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

// NewErrAt creates a new error.Err[ErrorCode] error.
//
// Parameters:
//   - at: The operation at which the error occurred.
//   - reason: The reason for the error.
//
// Returns:
//   - *error.Err[ErrorCode]: The new error. Never returns nil.
func NewErrAt(at string, reason error) *gcers.Err[ErrorCode] {
	var msg string

	if at == "" {
		msg = "an error occurred somewhere"
	} else {
		msg = "an error occurred at " + at
	}

	err := gcers.NewErr(gcers.FATAL, OperationFail, msg)
	err.SetInner(reason)

	return err
}

// NewErrAfter creates a new error.Err[ErrorCode] error.
//
// Parameters:
//   - before: The operation after which the error occurred.
//   - reason: The reason for the error.
//
// Returns:
//   - *error.Err[ErrorCode]: The new error. Never returns nil.
func NewErrAfter(before string, reason error) *gcers.Err[ErrorCode] {
	var msg string

	if before == "" {
		msg = "an error occurred after something"
	} else {
		msg = "an error occurred after " + before
	}

	err := gcers.NewErr(gcers.FATAL, OperationFail, msg)
	err.SetInner(reason)

	return err
}

// NewErrBefore creates a new error.Err[ErrorCode] error.
//
// Parameters:
//   - after: The operation before which the error occurred.
//   - reason: The reason for the error.
//
// Returns:
//   - *error.Err[ErrorCode]: The new error. Never returns nil.
func NewErrBefore(after string, reason error) *gcers.Err[ErrorCode] {
	var msg string

	if after == "" {
		msg = "an error occurred before something"
	} else {
		msg = "an error occurred before " + after
	}

	err := gcers.NewErr(gcers.FATAL, OperationFail, msg)
	err.SetInner(reason)

	return err
}
