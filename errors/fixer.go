package errors

import (
	"fmt"

	gcers "github.com/PlayerR9/go-commons/errors/error"
)

// Fixer is defines the behavior of an object that can be fixed. This
// must not have a non-pointer receiver.
type Fixer interface {
	// Fix fixes the object.
	//
	// Returns:
	//   - error: An error that occurred while fixing the object.
	Fix() error
}

// Fix fixes the object.
//
// Parameters:
//   - name: The name of the object.
//   - obj: The object to fix.
//   - allow_nil: Whether to allow the object to be nil.
//
// Returns:
//   - error: An error that occurred while fixing the object.
func Fix(name string, obj Fixer, allow_nil bool) error {
	if name == "" {
		name = "struct{}"
	}

	if obj == nil && !allow_nil {
		err := NewErrFix(fmt.Sprintf("%q must not be nil", name))

		return err
	}

	err := obj.Fix()
	if err == nil {
		return nil
	}

	sub_err, ok := err.(*gcers.Err[ErrorCode])
	if !ok {
		sub_err = NewErrFix(err.Error())
	}

	sub_err.AddFrame(name, "Fix()")

	return sub_err
}
