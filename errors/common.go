package errors

import (
	gcers "github.com/PlayerR9/go-commons/errors/error"
)

// Is is function that checks if an error is of type T.
//
// Parameters:
//   - err: The error to check.
//   - code: The error code to check.
//
// Returns:
//   - bool: true if the error is of type T, false otherwise (including if the error is nil).
func Is[T gcers.ErrorCoder](err error, code T) bool {
	if err == nil {
		return false
	}

	sub_err, ok := err.(*gcers.Err[T])
	if !ok {
		return false
	}

	return sub_err.Code == code
}
