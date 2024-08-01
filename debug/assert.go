package debug

import (
	"fmt"
	"strings"
)

// Assert panics iff the condition is false. The panic is not a string
// but an error of type *ErrAssertionFailed.
//
// Parameters:
//   - cond: the condition to check.
//   - msg: the message to show if the condition is not met.
//
// Example:
//
//	foo := "foo"
//	Assert(foo == "bar", "foo is not bar") // panics: "assertion failed: foo is not bar"
func Assert(cond bool, msg string) {
	if cond {
		return
	}

	panic(NewErrAssertionFailed(msg))
}

// AssertF same as Assert but with a format string and arguments that are in
// accordance with fmt.Printf.
//
// Parameters:
//   - cond: the condition to check.
//   - format: the format string to show if the condition is not met.
//   - args: the arguments to pass to the format string.
//
// Example:
//
//	foo := "foo"
//	bar := "bar"
//	AssertF(foo == bar, "%s is not %s", foo, bar) // panics: "assertion failed: foo is not bar"
func AssertF(cond bool, format string, args ...any) {
	if cond {
		return
	}

	panic(NewErrAssertionFailed(fmt.Sprintf(format, args...)))
}

// AssertErr is the same as Assert but for errors. Best used for ensuring that a function
// does not return an unexpected error.
//
// Parameters:
//   - err: the error to check.
//   - format: the format describing the function's signature.
//   - args: the arguments passed to the function.
//
// Example:
//
//	foo := "foo"
//	err := my_function(foo, "bar")
//	AssertErr(err, "my_function(%s, %s)", foo, "bar")
//	// panics: "assertion failed: function my_function(foo, bar) returned the error: <err>"
func AssertErr(err error, format string, args ...any) {
	if err == nil {
		return
	}

	var builder strings.Builder

	builder.WriteString("function ")
	fmt.Fprintf(&builder, format, args...)
	builder.WriteString(" returned the error: ")
	builder.WriteString(err.Error())

	panic(NewErrAssertionFailed(builder.String()))
}

// AssertOk is the same as Assert but for booleans. Best used for ensuring that a function that
// are supposed to return the boolean `true` does not return `false`.
//
// Parameters:
//   - cond: the result of the function.
//   - format: the format describing the function's signature.
//   - args: the arguments passed to the function.
//
// Example:
//
//	foo := "foo"
//	ok := my_function(foo, "bar")
//	AssertOk(ok, "my_function(%s, %s)", foo, "bar")
//	// panics: "assertion failed: function my_function(foo, bar) returned false while true was expected"
func AssertOk(cond bool, format string, args ...any) {
	if cond {
		return
	}

	var builder strings.Builder

	builder.WriteString("function ")
	fmt.Fprintf(&builder, format, args...)
	builder.WriteString(" returned false while true was expected")

	panic(NewErrAssertionFailed(builder.String()))
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
