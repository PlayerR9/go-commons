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
func (e *ErrAt) ChangeReason(reason error) {
	if e == nil {
		return
	}

	e.Reason = reason
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
