package errors

import "errors"

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
