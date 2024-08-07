package errors

import (
	"strconv"
	"strings"
)

// ErrInvalidParameter represents an error when a parameter is invalid.
type ErrInvalidParameter struct {
	// Parameter is the invalid parameter.
	Parameter string

	// Reason is the reason for the error.
	Reason error
}

// Error implements the error interface.
//
// Message:
// - "parameter (<parameter>) is invalid" if Reason is nil
// - "parameter (<parameter>) is invalid: <reason>" if Reason is not nil
func (e *ErrInvalidParameter) Error() string {
	var parameter string

	if e.Parameter != "" {
		parameter = "(" + strconv.Quote(e.Parameter) + ")"
	}

	var builder strings.Builder

	builder.WriteString("parameter ")
	builder.WriteString(parameter)
	builder.WriteString(" is invalid")

	if e.Reason != nil {
		builder.WriteString(": ")
		builder.WriteString(e.Reason.Error())
	}

	return builder.String()
}

// NewErrInvalidParameter creates a new ErrInvalidParameter error.
//
// Parameters:
//   - parameter: The invalid parameter.
//   - reason: The reason for the error.
//
// Returns:
//   - *ErrInvalidParameter: A pointer to the newly created ErrInvalidParameter. Never returns nil.
func NewErrInvalidParameter(parameter string, reason error) *ErrInvalidParameter {
	return &ErrInvalidParameter{
		Parameter: parameter,
		Reason:    reason,
	}
}

// Unwrap is a method that returns the wrapped error.
//
// Returns:
//   - error: The wrapped error.
func (e *ErrInvalidParameter) Unwrap() error {
	return e.Reason
}

// ChangeReason is a method that changes the reason for the error.
//
// Parameters:
//   - reason: The new reason for the error.
//
// Returns:
//   - error: The new reason for the error.
func (e *ErrInvalidParameter) ChangeReason(reason error) {
	e.Reason = reason
}

// NewErrNilParameter is a convenience method that creates a new *ErrInvalidParameter error
// with a NilValue as the reason.
//
// Parameters:
//   - parameter: The invalid parameter.
//
// Returns:
//   - *ErrInvalidParameter: A pointer to the newly created ErrInvalidParameter. Never returns nil.
func NewErrNilParameter(parameter string) *ErrInvalidParameter {
	return &ErrInvalidParameter{
		Parameter: parameter,
		Reason:    NilValue,
	}
}
