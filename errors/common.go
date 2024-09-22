package errors

import (
	"errors"
)

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
	ChangeReason(new_reason error)
}

// LimitErrorMsg is a function that limits the number of errors in an error chain.
//
// Parameters:
//   - err: The error to limit.
//   - limit: The maximum number of errors to limit.
//
// Returns:
//   - error: The limited error.
//
// If the error is nil or the limit is less than or equal to 0, the function returns nil.
func LimitErrorMsg(err error, limit int) error {
	if err == nil || limit <= 0 {
		return nil
	}

	target := err

	for i := 0; i < limit; i++ {
		w, ok := target.(Unwrapper)
		if !ok {
			return err
		}

		reason := w.Unwrap()
		if reason == nil {
			return err
		}

		target = reason
	}

	if target == nil {
		return err
	}

	w, ok := target.(Unwrapper)
	if !ok {
		return err
	}

	w.ChangeReason(nil)

	return err
}

// Error returns the error message for an error.
//
// Parameters:
//   - err: The error to get the message for.
//
// Returns:
//   - string: The error message.
//
// If the error is nil, the function returns "something went wrong".
func Reason(err error) string {
	if err == nil {
		return "something went wrong"
	}

	return err.Error()
}

// Is is function that checks if an error is of type T.
//
// Parameters:
//   - err: The error to check.
//
// Returns:
//   - bool: true if the error is of type T, false otherwise (including if the error is nil).
func Is[T any](err error) bool {
	if err == nil {
		return false
	}

	var target T

	return errors.As(err, &target)
}

/* func Wrap(outer, inner error) error {
	if outer == nil {
		if inner == nil {

		} else {

		}
	} else {
		if outer == nil {

		} else {

		}
	}
} */
