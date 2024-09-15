package strings

import (
	"fmt"
	"strconv"
)

// GoStringOf returns a string representation of the element.
//
// Parameters:
//   - elem: The element to get the string representation of.
//
// Returns:
//   - string: The string representation of the element.
//
// Behaviors:
//   - If the element is nil, the function returns "nil".
//   - If the element implements the fmt.GoStringer interface, the function
//     returns the result of the GoString method.
//   - If the element implements the fmt.Stringer interface, the function
//     returns the result of the String method.
//   - If the element is a string, the function returns the string enclosed in
//     double quotes.
//   - If the element is an error, the function returns the error message
//     enclosed in double quotes.
//   - Otherwise, the function returns the result of the %#v format specifier.
func GoStringOf(elem any) string {
	if elem == nil {
		return "nil"
	}

	switch elem := elem.(type) {
	case fmt.GoStringer:
		return elem.GoString()
	case fmt.Stringer:
		return strconv.Quote(elem.String())
	case string:
		return strconv.Quote(elem)
	case error:
		return strconv.Quote(elem.Error())
	default:
		return fmt.Sprintf("%#v", elem)
	}
}
