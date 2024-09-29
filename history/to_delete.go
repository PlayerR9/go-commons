package history

import (
	gers "github.com/PlayerR9/go-errors"
	gerr "github.com/PlayerR9/go-errors/error"
)

// TODO: Remove this once go-errors is updated.

// NewErrNilReceiver creates a new error.Err error with the code
// OperationFail.
//
// Returns:
//   - *error.Err: The new error. Never returns nil.
func NewErrNilReceiver() *gerr.Err {
	err := gerr.New(gers.OperationFail, "receiver must not be nil")
	err.AddSuggestion("Did you forget to initialize the receiver?")

	return err
}
