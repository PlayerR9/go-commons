package runes

import "unicode/utf8"

// BytesToUtf8 is a function that converts bytes to runes.
//
// Parameters:
//   - data: The bytes to convert.
//
// Returns:
//   - []rune: The runes.
//   - error: An error of type *ErrInvalidUTF8Encoding if the bytes are not
//     valid UTF-8.
//
// This function also converts '\r\n' to '\n'. Plus, whenever an error occurs, it returns the runes
// decoded so far and the index of the error rune.
func BytesToUtf8(data []byte) ([]rune, error) {
	if len(data) == 0 {
		return nil, nil
	}

	var chars []rune
	var i int

	for len(data) > 0 {
		c, size := utf8.DecodeRune(data)
		if c == utf8.RuneError {
			return chars, NewErrInvalidUTF8Encoding(i)
		}

		data = data[size:]
		i += size

		if c != '\r' {
			chars = append(chars, c)
			continue
		}

		if len(data) == 0 {
			return chars, NewErrInvalidUTF8Encoding(i)
		}

		c, size = utf8.DecodeRune(data)
		if c == utf8.RuneError {
			return chars, NewErrInvalidUTF8Encoding(i)
		}

		data = data[size:]
		i += size

		if c != '\n' {
			return chars, NewErrInvalidUTF8Encoding(i)
		}

		chars = append(chars, '\n')
	}

	return chars, nil
}

// StringToUtf8 converts a string to a slice of runes.
//
// Parameters:
//   - str: The string to convert.
//
// Returns:
//   - runes: The slice of runes.
//   - error: An error of type *ErrInvalidUTF8Encoding if the string is not
//     valid UTF-8.
//
// Behaviors:
//   - An empty string returns a nil slice with no errors.
//   - The function stops at the first invalid UTF-8 encoding; returning an
//     error and the runes found up to that point.
//   - The function converts '\r\n' to '\n'.
func StringToUtf8(str string) ([]rune, error) {
	if str == "" {
		return nil, nil
	}

	var chars []rune
	var i int

	for len(str) > 0 {
		c, size := utf8.DecodeRuneInString(str)
		if c == utf8.RuneError {
			return chars, NewErrInvalidUTF8Encoding(i)
		}

		str = str[size:]
		i += size

		if c != '\r' {
			chars = append(chars, c)
			continue
		}

		if len(str) == 0 {
			return chars, NewErrInvalidUTF8Encoding(i)
		}

		c, size = utf8.DecodeRuneInString(str)
		if c == utf8.RuneError {
			return chars, NewErrInvalidUTF8Encoding(i)
		}

		str = str[size:]
		i += size

		if c != '\n' {
			return chars, NewErrInvalidUTF8Encoding(i)
		}

		chars = append(chars, '\n')
	}

	return chars, nil
}
