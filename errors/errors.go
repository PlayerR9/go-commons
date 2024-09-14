package errors

import (
	"errors"
	"strconv"
	"strings"
)

var (
	// NilValue is an error that is returned when a value is nil.
	NilValue error

	// NilReceiver is an error that is returned when a receiver is nil.
	NilReceiver error
)

func init() {
	NilValue = errors.New("value must not be nil")

	NilReceiver = errors.New("receiver must not be nil")
}

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
func (e *ErrInvalidParameter) ChangeReason(new_reason error) {
	if e == nil {
		return
	}

	e.Reason = new_reason
}
