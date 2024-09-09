package assert

import ers "github.com/PlayerR9/go-commons/errors"

// Assertion is the struct that is used to perform assertions.
type Assertion[T Asserter] struct {
	// target is the target to assert.
	target Target

	// assert is the assertion to perform.
	assert T

	// negative is true if the assertion should be negated.
	negative bool
}

// AssertThat is a constructor for the Assertion struct.
//
// Parameters:
//   - target: the target to assert.
//   - assert: the assertion to perform.
//
// Returns:
//   - *Assertion: the assertion. Never returns nil.
//
// A normal construction is a chain of AssertThat() function followed by
// the conditions and the action to perform.
//
// Example:
//
//	foo := "foo"
//	AssertThat(NewVariable("foo"), NewOrderedAssert(foo).In("bar", "fooo", "baz")).Not().Panic()
//	// Does not panic since foo is not in ["bar", "fooo", "baz"]
func AssertThat[T Asserter](target Target, assert T) *Assertion[T] {
	return &Assertion[T]{
		target: target,
		assert: assert,
	}
}

// Not negates the assertion. Useful for checking the negation of an assertion.
//
// However, if the positive check is more expensive than its negative counterpart,
// it is suggested to create the negative assertion rather than negating the positive one.
//
// Furthermore, if more than one Not() function is called on the same assertion,
// then if the count of the Not() functions is odd, the assertion will be negated. Otherwise,
// the assertion will be positive.
//
// For example, doing .Not().Not().Not() is the same as .Not().
//
// Returns:
//   - *Assertion: the assertion for chaining. Never returns nil.
func (a *Assertion[T]) Not() *Assertion[T] {
	a.negative = !a.negative
	return a
}

// Panic will panic if the condition is not met.
//
// The error message is "expected <name> to <message>; got <value> instead" where
// <name> is the name of the assertion, <message> is the message of the condition
// and <value> is the value of the assertion. Finally, this error message is used
// within the *ErrAssertionFailed error.
func (a *Assertion[T]) Panic() {
	ok := a.assert.Verify()
	if ok != a.negative {
		return
	}

	panic(ers.NewErrAssertFailed(a.assert.Message(a.target, a.negative)))
}

// PanicWithMessage is the same as Panic but with a custom error message.
// This error message is overrides the default error message of the Assertion.
//
// Of course, the message is still used within the *ErrAssertionFailed error.
func (a *Assertion[T]) PanicWithMessage(msg string) {
	ok := a.assert.Verify()
	if ok != a.negative {
		return
	}

	panic(ers.NewErrAssertFailed(msg))
}

// Error same as Panic but returns the *ErrAssertionFailed error instead of a panic.
//
// The error message is "expected <name> to <message>; got <value> instead" where
// <name> is the name of the assertion, <message> is the message of the condition
// and <value> is the value of the assertion.
//
// Returns:
//   - error: the error. Nil iff the condition is met.
func (a *Assertion[T]) Error() error {
	ok := a.assert.Verify()
	if ok != a.negative {
		return nil
	}

	return ers.NewErrAssertFailed(a.assert.Message(a.target, a.negative))
}

// ErrorWithMessage is the same as PanicWithMessage but returns the *ErrAssertionFailed error instead of a panic.
// This error message is overrides the default error message of the Assertion.
//
// Of course, the message is still used within the *ErrAssertionFailed error.
//
// Returns:
//   - error: the error. Nil iff the condition is met.
func (a *Assertion[T]) ErrorWithMessage(msg string) error {
	ok := a.assert.Verify()
	if ok != a.negative {
		return nil
	}

	return ers.NewErrAssertFailed(msg)
}

// Check checks whether the condition is met.
//
// Returns:
//   - bool: true if the condition is met. false otherwise.
func (a *Assertion[T]) Check() bool {
	ok := a.assert.Verify()
	return ok != a.negative
}
