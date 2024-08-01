package debug

import (
	"testing"
)

func TestAssertion(t *testing.T) {
	foo := "foo"

	var actual_err ErrAssertionFailed

	res := AssertThat("foo", foo).Not().In("bar", "foo", "baz").Error()
	if res == nil {
		t.Errorf("expected an error but got none")
	} else {
		actual_err = *res
	}

	msg_to_check := "expected \"foo\" to not be one of {bar, baz, foo}; got foo instead"

	if actual_err.Msg != msg_to_check {
		t.Errorf("expected error %q but got %q instead", msg_to_check, actual_err.Msg)
	}
}
