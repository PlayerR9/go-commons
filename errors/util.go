package errors

import (
	"errors"
	"strconv"
	"strings"
)

// Is is function that checks if an error is of type T.
//
// Parameters:
//   - err: The error to check.
//
// Returns:
//   - bool: true if the error is of type T, false otherwise (including if the error is nil).
func Is[T error](err error) bool {
	if err == nil {
		return false
	}

	var target T

	ok := errors.As(err, &target)
	return ok
}

// Error returns the error message of an error.
//
// Parameters:
//   - err: The error to get the message of.
//
// Returns:
//   - string: The error message of the error.
//
// If the error is nil, the function returns "something went wrong" as the error message.
func Error(err error) string {
	if err == nil {
		return "something went wrong"
	}

	return err.Error()
}

// Unwrapper is an interface that defines a method to unwrap an error.
type Unwrapper interface {
	// Unwrap returns the error that this error wraps.
	//
	// Returns:
	//   - error: The error that this error wraps.
	Unwrap() error

	// ChangeReason changes the reason of the error.
	//
	// Parameters:
	//   - reason: The new reason of the error.
	ChangeReason(reason error)
}

// LimitErrorMsg limits the error message to a certain number of unwraps.
// It returns the top level error for allowing to print the error message
// with the limit of unwraps applied.
//
// If the error is nil or the limit is less than 0, the function does nothing.
//
// Parameters:
//   - err: The error to limit.
//   - limit: The limit of unwraps.
//
// Returns:
//   - error: The top level error with the limit of unwraps applied.
func LimitErrorMsg(err error, limit int) error {
	if err == nil || limit < 0 {
		return err
	}

	currErr := err

	for i := 0; i < limit; i++ {
		target, ok := currErr.(Unwrapper)
		if !ok {
			return err
		}

		reason := target.Unwrap()
		if reason == nil {
			return err
		}

		currErr = reason
	}

	// Limit reached
	target, ok := currErr.(Unwrapper)
	if !ok {
		return err
	}

	target.ChangeReason(nil)

	return err
}

// GetOrdinalSuffix returns the ordinal suffix for a given integer.
//
// Parameters:
//   - number: The integer for which to get the ordinal suffix. Negative
//     numbers are treated as positive.
//
// Returns:
//   - string: The ordinal suffix for the number.
//
// Example:
//   - GetOrdinalSuffix(1) returns "1st"
//   - GetOrdinalSuffix(2) returns "2nd"
func GetOrdinalSuffix(number int) string {
	var builder strings.Builder

	builder.WriteString(strconv.Itoa(number))

	if number < 0 {
		number = -number
	}

	lastTwoDigits := number % 100
	lastDigit := lastTwoDigits % 10

	if lastTwoDigits >= 11 && lastTwoDigits <= 13 {
		builder.WriteString("th")
	} else {
		switch lastDigit {
		case 1:
			builder.WriteString("st")
		case 2:
			builder.WriteString("nd")
		case 3:
			builder.WriteString("rd")
		default:
			builder.WriteString("th")
		}
	}

	return builder.String()
}
