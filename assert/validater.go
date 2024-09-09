package assert

import (
	"errors"
	"strings"

	ers "github.com/PlayerR9/go-commons/errors"
)

// Validater is an interface that can be validated.
type Validater interface {
	// Validate checks the validity of the object (i.e., the integrity of the object).
	//
	// Returns:
	//   - error: An error if the object is invalid. Nil if the object is valid.
	Validate() error
}

// Validate is a function that checks the validity of the object.
//
// Parameters:
// - obj: The object to validate.
// - allow_nil: True if nil is an allowed value. False otherwise.
//
// Returns:
//   - error: An error if the object is invalid. Nil if the object is valid.
func Validate(obj Validater, allow_nil bool) error {
	if obj == nil {
		if !allow_nil {
			return errors.New("expected value to not be nil")
		} else {
			return nil
		}
	}

	err := obj.Validate()
	if err != nil {
		return err
	}

	return nil
}

// AssertValidate panics if the object is invalid.
//
// Parameters:
// - context: the context of the call. If empty, "call to method" is used.
// - target: the target of the validation.
// - obj: the object to validate.
// - allow_nil: true if nil is an allowed value. False otherwise.
//
// Panics:
//   - *ErrAssertionFailed: if the object is invalid.
func AssertValidate(context string, target Target, obj Validater, allow_nil bool) {
	err := Validate(obj, allow_nil)
	if err == nil {
		return
	}

	var builder strings.Builder

	builder.WriteString("call to method ")

	if context != "" {
		builder.WriteString(context)
		builder.WriteRune(' ')
	}

	builder.WriteString("failed as ")
	builder.WriteString(target.String())
	builder.WriteString("is invalid: ")
	builder.WriteString(err.Error())

	panic(ers.NewErrAssertFailed(builder.String()).Error())
}
