package runes

import (
	"slices"
	"unicode/utf8"
)

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
