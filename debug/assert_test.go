package debug

import (
	"errors"
	"testing"
)

func TestAssert(t *testing.T) {
	defer func() {
		r := recover()

		if r == nil {
			t.Error("Expected an error but got none")
		}

		err, ok := r.(*ErrAssertionFailed)
		if !ok {
			t.Errorf("Expected an error of type *ErrAssertionFailed but got %T", r)
		}

		if err.Msg != "foo is not bar" {
			t.Errorf("Expected 'foo is not bar' but got %s", err.Msg)
		}
	}()

	foo := "foo"

	Assert(foo == "bar", "foo is not bar")
}

func TestAssertF(t *testing.T) {
	defer func() {
		r := recover()

		if r == nil {
			t.Error("Expected an error but got none")
		}

		err, ok := r.(*ErrAssertionFailed)
		if !ok {
			t.Errorf("Expected an error of type *ErrAssertionFailed but got %T", r)
		}

		if err.Msg != "foo is not bar" {
			t.Errorf("Expected 'foo is not bar' but got %s", err.Msg)
		}
	}()

	foo := "foo"
	bar := "bar"

	AssertF(foo == bar, "%s is not %s", foo, bar)
}

func TestAssertErr(t *testing.T) {
	defer func() {
		r := recover()

		if r == nil {
			t.Error("Expected an error but got none")
		}

		err, ok := r.(*ErrAssertionFailed)
		if !ok {
			t.Errorf("Expected an error of type *ErrAssertionFailed but got %T", r)
		}

		if err.Msg != "function my_function(foo, bar) returned the error: <err>" {
			t.Errorf("Expected 'function my_function(foo, bar) returned the error: <err>' but got %s", err.Msg)
		}
	}()

	foo := "foo"

	my_function := func(left, right string) error {
		t.Logf("my_function(%s, %s)", left, right)

		return errors.New("<err>")
	}

	err := my_function(foo, "bar")
	AssertErr(err, "my_function(%s, %s)", foo, "bar")
}

func TestAssertOk(t *testing.T) {
	defer func() {
		r := recover()

		if r == nil {
			t.Error("Expected an error but got none")
		}

		err, ok := r.(*ErrAssertionFailed)
		if !ok {
			t.Errorf("Expected an error of type *ErrAssertionFailed but got %T", r)
		}

		if err.Msg != "function my_function(foo, bar) returned false while true was expected" {
			t.Errorf("Expected 'function my_function(foo, bar) returned false while true was expected' but got %s", err.Msg)
		}
	}()

	foo := "foo"

	my_function := func(left, right string) bool {
		t.Logf("my_function(%s, %s)", left, right)

		return false
	}

	ok := my_function(foo, "bar")
	AssertOk(ok, "my_function(%s, %s)", foo, "bar")
}

///////////////////////////////

/*
func AssertDerefNil[T any](elem *T, param_name string) T {
	if elem != nil {
		return *elem
	}

	var builder strings.Builder

	builder.WriteString("Parameter (")
	builder.WriteString(strconv.Quote(param_name))
	builder.WriteString(") must not be nil")

	panic(builder.String())
}

func AssertNil(elem any, param_name string) {
	if elem != nil {
		return
	}

	var builder strings.Builder

	builder.WriteString("Parameter (")
	builder.WriteString(strconv.Quote(param_name))
	builder.WriteString(") must not be nil")

	panic(builder.String())
}

func AssertType(elem any, expected string, allow_nil bool, var_name string) {
	if elem == nil {
		if !allow_nil {
			panic(fmt.Sprintf("expected %q to be of type %s, got nil instead", var_name, expected))
		}

		return
	}

	to := reflect.TypeOf(elem)
	if to.String() != expected {
		panic(fmt.Sprintf("expected %q to be of type %s, got %T instead", var_name, expected, elem))
	}
}

func AssertConv[T any](elem any, var_name string) T {
	if elem == nil {
		panic(fmt.Sprintf("expected %q to be of type %T, got nil instead", var_name, *new(T)))
	}

	res, ok := elem.(T)
	if !ok {
		panic(fmt.Sprintf("expected %q to be of type %T, got %T instead", var_name, *new(T), elem))
	}

	return res
}
*/
