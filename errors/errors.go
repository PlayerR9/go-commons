package errors

import (
	"fmt"
	"strconv"
	"strings"
)

// Enumer is an enum type.
type Enumer interface {
	~int

	// String returns the literal name of the enum.
	//
	// Returns:
	//   - string: The name of the enum.
	String() string
}

// Err is the base error type.
type Err[T Enumer] struct {
	// Code is the error code.
	Code T

	// Message is the error message.
	Message error

	// Suggestions are the error suggestions.
	Suggestions []string
}

// Error implements the error interface.
//
// Message: "{code}: {message}"
func (e Err[T]) Error() string {
	return fmt.Sprintf("%s: %s", e.Code.String(), Error(e.Message))
}

// Unwrap implements errors.Unwrap interface.
func (e Err[T]) Unwrap() error {
	return e.Message
}

// NewErr creates a new error.
//
// Parameters:
//   - code: The error code.
//   - message: The error message.
//
// Returns:
//   - *Err: The new error. Never returns nil.
func NewErr[T Enumer](code T, message error) *Err[T] {
	return &Err[T]{
		Code:        code,
		Message:     message,
		Suggestions: nil,
	}
}

// AddSuggestion adds a suggestion of the error.
//
// Parameters:
//   - suggestions: The suggestion of the error.
//
// Returns:
//   - *Err: The error. Never returns nil.
//
// Each element in the suggestion is separated by a space but each call to this function
// adds each suggestion on a new line.
func (e *Err[T]) AddSuggestion(suggestions ...string) *Err[T] {
	e.Suggestions = append(e.Suggestions, strings.Join(suggestions, " "))

	return e
}

// ChangeReason changes the reason of the error.
//
// Parameters:
//   - reason: The new reason of the error.
func (e *Err[T]) ChangeReason(reason error) {
	e.Message = reason
}

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

// Unwrap implements the errors.Unwrap interface.
func (e *ErrInvalidParameter) Unwrap() error {
	return e.Reason
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

// ErrInvalidUsage represents an error that occurs when a function is used incorrectly.
type ErrInvalidUsage struct {
	// Reason is the reason for the invalid usage.
	Reason error

	// Usage is the usage of the function.
	Usage string
}

// Error is a method of the Unwrapper interface.
//
// Message: "{reason}. {usage}".
//
// However, if the reason is nil, the message is "invalid usage. {usage}" instead.
//
// If the usage is empty, no usage is added to the message.
func (e *ErrInvalidUsage) Error() string {
	var builder strings.Builder

	if e.Reason == nil {
		builder.WriteString("invalid usage")
	} else {
		builder.WriteString(e.Reason.Error())
	}

	if e.Usage != "" {
		builder.WriteString(". ")
		builder.WriteString(e.Usage)
	}

	return builder.String()
}

// Unwrap implements the errors.Unwrap interface.
func (e *ErrInvalidUsage) Unwrap() error {
	return e.Reason
}

// NewErrInvalidUsage creates a new ErrInvalidUsage error.
//
// Parameters:
//   - reason: The reason for the invalid usage.
//   - usage: The usage of the function.
//
// Returns:
//   - *ErrInvalidUsage: A pointer to the new ErrInvalidUsage error.
func NewErrInvalidUsage(reason error, usage string) *ErrInvalidUsage {
	return &ErrInvalidUsage{
		Reason: reason,
		Usage:  usage,
	}
}

// ChangeReason is a method that changes the reason for the error.
//
// Parameters:
//   - reason: The new reason for the error.
func (e *ErrInvalidUsage) ChangeReason(reason error) {
	e.Reason = reason
}
