package errors

import (
	"strconv"
	"strings"
)

// ErrInvalidParameter is an error that is returned when a parameter is invalid.
type ErrInvalidParameter struct {
	// Parameter is the name of the invalid parameter.
	Parameter string

	// Reason is the reason for the error.
	Reason error
}

// Error implements the error interface.
//
// Message:
//
//	"parameter <parameter> is invalid: <reason>"
func (e ErrInvalidParameter) Error() string {
	var builder strings.Builder

	builder.WriteString("parameter ")
	builder.WriteString(strconv.Quote(e.Parameter))
	builder.WriteString(" is invalid")

	if e.Reason != nil {
		builder.WriteString(": ")
		builder.WriteString(e.Reason.Error())
	}

	return builder.String()
}

// Unwrap implements the errors.Unwrap interface.
func (e ErrInvalidParameter) Unwrap() error {
	return e.Reason
}

// NewErrInvalidParameter creates a new ErrInvalidParameter error.
//
// Parameters:
//   - parameter: the name of the invalid parameter.
//   - reason: the reason for the error.
//
// Returns:
//   - *ErrInvalidParameter: the new error. Never returns nil.
func NewErrInvalidParameter(parameter string, reason error) *ErrInvalidParameter {
	return &ErrInvalidParameter{
		Parameter: parameter,
		Reason:    reason,
	}
}

// NewErrNilParameter creates a new ErrInvalidParameter error.
//
// Parameters:
//   - parameter: the name of the invalid parameter.
//
// Returns:
//   - *ErrInvalidParameter: the new error. Never returns nil.
func NewErrNilParameter(parameter string) *ErrInvalidParameter {
	return &ErrInvalidParameter{
		Parameter: parameter,
		Reason:    NilValue,
	}
}

// ChangeReason changes the reason for the error.
//
// Parameters:
//   - new_reason: the new reason for the error.
func (e *ErrInvalidParameter) ChangeReason(new_reason error) {
	if e == nil {
		return
	}

	e.Reason = new_reason
}

// ErrFix is an error that is returned when an object cannot be fixed.
type ErrFix struct {
	// Name is the name of the object.
	Name string

	// Reason is the reason for the error.
	Reason error
}

// Error implements the error interface.
//
// Message:
//
//	"failed to fix <name>: <reason>"
func (e ErrFix) Error() string {
	var name string

	if e.Name == "" {
		name = "object"
	} else {
		name = strconv.Quote(e.Name)
	}

	var builder strings.Builder

	builder.WriteString("failed to fix ")
	builder.WriteString(name)

	if e.Reason != nil {
		builder.WriteString(": ")
		builder.WriteString(e.Reason.Error())
	}

	return builder.String()
}

// Unwrap implements the errors.Unwrap interface.
func (e ErrFix) Unwrap() error {
	return e.Reason
}

// NewErrFix creates a new ErrFix error.
//
// Parameters:
//   - name: the name of the object.
//   - reason: the reason for the error.
//
// Returns:
//   - *ErrFix: the new error. Never returns nil.
func NewErrFix(name string, reason error) *ErrFix {
	return &ErrFix{
		Name:   name,
		Reason: reason,
	}
}

// ChangeReason changes the reason for the error.
//
// Parameters:
//   - new_reason: the new reason for the error.
func (e *ErrFix) ChangeReason(new_reason error) {
	if e == nil {
		return
	}

	e.Reason = new_reason
}
