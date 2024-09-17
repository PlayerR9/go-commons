package runes

import (
	"slices"
	"strconv"
	"strings"
	"unicode/utf8"
)

// EitherOrString is a function that returns a string representation of a slice
// of strings. Empty strings are ignored.
//
// Parameters:
//   - values: The values to convert to a string.
//
// Returns:
//   - string: The string representation.
//
// Example:
//
//	EitherOrString([]rune{'a', 'b', 'c'}, false) // "either a, b or c"
func EitherOrString(elems []rune, quote bool) string {
	if len(elems) == 0 {
		return ""
	} else if len(elems) == 1 {
		if quote {
			return strconv.QuoteRune(elems[0])
		} else {
			return string(elems[0])
		}
	}

	values := make([]string, 0, len(elems))

	if quote {
		for _, elem := range elems {
			values = append(values, strconv.QuoteRune(elem))
		}
	} else {
		for _, elem := range elems {
			values = append(values, string(elem))
		}
	}

	var builder strings.Builder

	builder.WriteString("either ")

	if len(values) > 2 {
		builder.WriteString(strings.Join(values[:len(values)-1], ", "))
		builder.WriteRune(',')
	} else {
		builder.WriteString(values[0])
	}

	builder.WriteString(" or ")
	builder.WriteString(values[len(values)-1])

	return builder.String()
}

// NormalizeRunes is a function that converts '\r\n' to '\n'.
//
// Parameters:
//   - chars: The characters to convert.
//
// Returns:
//   - []rune: The normalized characters.
//   - error: An error if the characters are not valid UTF-8.
//
// Errors:
//   - *ErrUnexpectedChar: If the characters are not valid UTF-8.
func NormalizeRunes(chars []rune) ([]rune, error) {
	if len(chars) == 0 {
		return chars, nil
	}

	indices := IndicesOf(chars, '\r', false)

	for _, idx := range indices {
		if idx+1 >= len(chars) {
			return chars, NewErrUnexpectedChar(chars[idx], []rune{'\r'}, nil)
		}

		next := chars[idx+1]
		if next != '\n' {
			return chars, NewErrUnexpectedChar(chars[idx], []rune{'\r'}, &next)
		}
	}

	for len(indices) > 0 {
		idx := indices[0]
		indices = indices[1:]

		chars = slices.Delete(chars, idx, idx+1)

		// Update the indices.
		for i := 0; i < len(indices); i++ {
			indices[i] -= 1
		}
	}

	return chars, nil
}

// BytesToUtf8 is a function that converts bytes to runes. When error occurs,
// the function returns the runes decoded so far and the error.
//
// Parameters:
//   - data: The bytes to convert.
//
// Returns:
//   - []rune: The runes.
//   - error: An error if the bytes are not valid UTF-8.
//
// Errors:
//   - *ErrInvalidUTF8Encoding: If the bytes are not valid UTF-8.
func BytesToUtf8(data []byte) ([]rune, error) {
	if len(data) == 0 {
		return nil, nil
	}

	var chars []rune
	var i int

	for len(data) > 0 {
		c, size := utf8.DecodeRune(data)
		data = data[size:]

		if c == utf8.RuneError {
			return chars, NewErrInvalidUTF8Encoding(i)
		}

		i += size
		chars = append(chars, c)
	}

	return chars, nil
}

// StringToUtf8 converts a string to a slice of runes. When error occurs, the
// function returns the runes decoded so far and the error.
//
// Parameters:
//   - str: The string to convert.
//
// Returns:
//   - runes: The slice of runes.
//   - error: An error of if the string is not valid UTF-8.
//
// Errors:
//   - *ErrInvalidUTF8Encoding: If the string is not valid UTF-8.
func StringToUtf8(str string) ([]rune, error) {
	if str == "" {
		return nil, nil
	}

	var chars []rune
	var i int

	for len(str) > 0 {
		c, size := utf8.DecodeRuneInString(str)
		str = str[size:]

		if c == utf8.RuneError {
			return chars, NewErrInvalidUTF8Encoding(i)
		}

		i += size
		chars = append(chars, c)
	}

	return chars, nil
}
