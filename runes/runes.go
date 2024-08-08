package runes

import (
	"slices"
	"unicode"
	"unicode/utf8"
)

// ToInt converts a rune to an integer if possible. Conversion is case-insensitive and
// values from 0-9 and a-z are converted to 0-35.
//
// Parameters:
//   - char: The rune to convert.
//
// Returns:
//   - int: The converted integer.
//   - bool: True if the conversion was successful. False otherwise.
//
// Example:
//
//	digit, ok := ToInt('A')
//	if !ok {
//		panic("Could not convert 'A' to an integer")
//	}
//
//	fmt.Println(digit) // 10
func ToInt(char rune) (int, bool) {
	ok := unicode.IsDigit(char)
	if ok {
		return int(char - '0'), true
	}

	ok = unicode.IsLetter(char)
	if !ok {
		return 0, false
	}

	char = unicode.ToLower(char)

	return int(char - 'a' + 10), true
}

// FromInt converts an integer to a rune if possible. Conversion is case-insensitive and
// values from 0-9 and a-z are converted to 0-35.
//
// Parameters:
//   - digit: The integer to convert.
//
// Returns:
//   - rune: The converted rune.
//   - bool: True if the conversion was successful. False otherwise.
//
// Example:
//
//	char, ok := FromInt(10)
//	if !ok {
//		panic("Could not convert 10 to a rune")
//	}
//
//	fmt.Println(char) // 'A'
func FromInt(digit int) (rune, bool) {
	if digit < 0 || digit > 35 {
		return 0, false
	}

	if digit < 10 {
		return rune(digit + '0'), true
	}

	return rune(digit - 10 + 'a'), true
}

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

// Indices returns the indices of the separator in the data.
//
// Parameters:
//   - data: The data.
//   - sep: The separator.
//   - exclude_sep: Whether the separator is inclusive. If set to true, the indices will point to the character right after the
//     separator. Otherwise, the indices will point to the character right before the separator.
//
// Returns:
//   - []int: The indices.
func IndicesOf(data []rune, sep rune, exclude_sep bool) []int {
	if len(data) == 0 {
		return nil
	}

	var indices []int

	for i := 0; i < len(data); i++ {
		if data[i] == sep {
			indices = append(indices, i)
		}
	}

	if len(indices) == 0 {
		return nil
	}

	if exclude_sep {
		for i := 0; i < len(indices); i++ {
			indices[i] += 1
		}
	}

	return indices
}

// FindContentIndexes searches for the positions of opening and closing
// tokens in a slice of strings.
//
// Parameters:
//   - op_token: The string that marks the beginning of the content.
//   - cl_token: The string that marks the end of the content.
//   - tokens: The slice of strings in which to search for the tokens.
//
// Returns:
//   - result: An array of two integers representing the start and end indexes
//     of the content.
//   - err: Any error that occurred while searching for the tokens.
//
// Errors:
//   - *luc.ErrInvalidParameter: If the openingToken or closingToken is an
//     empty string.
//   - *ErrTokenNotFound: If the opening or closing token is not found in the
//     content.
//   - *ErrNeverOpened: If the closing token is found without any
//     corresponding opening token.
//
// Behaviors:
//   - The first index of the content is inclusive, while the second index is
//     exclusive.
//   - This function returns a partial result when errors occur. ([-1, -1] if
//     errors occur before finding the opening token, [index, 0] if the opening
//     token is found but the closing token is not found.
func FindContentIndexes(op_token, cl_token rune, tokens []rune) (result [2]int, err error) {
	result[0] = -1
	result[1] = -1

	op_tok_idx := slices.Index(tokens, op_token)
	if op_tok_idx < 0 {
		err = NewErrTokenNotFound(op_token, true)
		return
	} else {
		result[0] = op_tok_idx + 1
	}

	balance := 1
	cl_tok_idx := -1

	for i := result[0]; i < len(tokens) && cl_tok_idx == -1; i++ {
		curr_tok := tokens[i]

		if curr_tok == cl_token {
			balance--

			if balance == 0 {
				cl_tok_idx = i
			}
		} else if curr_tok == op_token {
			balance++
		}
	}

	if cl_tok_idx != -1 {
		result[1] = cl_tok_idx + 1
		return
	}

	if balance < 0 {
		err = NewErrNeverOpened(op_token, cl_token)
		return
	} else if balance != 1 || cl_token != '\n' {
		err = NewErrTokenNotFound(cl_token, false)
		return
	}

	result[1] = len(tokens)
	return
}
