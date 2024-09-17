package runes

import (
	"slices"
	"strconv"
	"strings"
	"unicode/utf8"
)

// JoinSize returns the number of runes in the data.
//
// Parameters:
//   - data: The data to join.
//
// Returns:
//   - int: The number of runes.
func JoinSize(data [][]rune) int {
	if len(data) == 0 {
		return 0
	}

	var size int

	for _, line := range data {
		size += len(line)
	}

	size += len(data) - 1

	return size
}

// Join is a function that joins the data. Returns nil if the data is empty.
//
// Parameters:
//   - data: The data to join.
//   - sep: The separator to use.
//
// Returns:
//   - []rune: The joined data.
func Join(data [][]rune, sep rune) []rune {
	if len(data) == 0 {
		return nil
	}

	size := JoinSize(data)

	result := make([]rune, 0, size)

	result = append(result, data[0]...)

	for _, line := range data[1:] {
		result = append(result, sep)
		result = append(result, line...)
	}

	return result
}

// split_size returns the number of lines and the maximum line length.
//
// Parameters:
//   - data: The data to split.
//   - sep: The separator to use.
//
// Returns:
//   - int: The number of lines.
//   - int: The maximum line length.
func split_size(data []rune, sep rune) (int, int) {
	var count int
	var max int
	var current int

	for _, c := range data {
		if c == sep {
			count++

			if current > max {
				max = current
			}

			current = 0
		} else {
			current++
		}
	}

	if current != 0 {
		count++

		if current > max {
			max = current
		}
	}

	return count, max
}

// Split is a function that splits the data into lines. Returns nil if the data is empty.
//
// Parameters:
//   - data: The data to split.
//   - sep: The separator to use.
//
// Returns:
//   - [][]rune: The lines.
func Split(data []rune, sep rune) [][]rune {
	if len(data) == 0 {
		return nil
	}

	count, max := split_size(data, sep)

	lines := make([][]rune, 0, count)
	current_line := make([]rune, 0, max)

	for i := 0; i < len(data); i++ {
		if data[i] != sep {
			current_line = append(current_line, data[i])

			continue
		}

		lines = append(lines, current_line[:len(current_line):len(current_line)])

		current_line = make([]rune, 0, max)
	}

	if len(current_line) > 0 {
		lines = append(lines, current_line)
	}

	return lines
}

// LimitReverseLines is a function that limits the lines of the data in reverse order.
//
// Parameters:
//   - data: The data to limit.
//   - limit: The limit of the lines.
//
// Returns:
//   - []byte: The limited data.
func LimitReverseLines(data []rune, limit int) []rune {
	if len(data) == 0 {
		return nil
	}

	lines := Split(data, '\n')

	if limit == -1 || limit > len(lines) {
		limit = len(lines)
	}

	start_idx := len(lines) - limit

	lines = lines[start_idx:]

	return Join(lines, '\n')
}

// LimitLines is a function that limits the lines of the data.
//
// Parameters:
//   - data: The data to limit.
//   - limit: The limit of the lines.
//
// Returns:
//   - []byte: The limited data.
func LimitLines(data []rune, limit int) []rune {
	if len(data) == 0 {
		return nil
	}

	lines := Split(data, '\n')

	if limit == -1 || limit > len(lines) {
		limit = len(lines)
	}

	lines = lines[:limit]

	return Join(lines, '\n')
}

// Repeat is a function that repeats the character.
//
// Parameters:
//   - char: The character to repeat.
//   - count: The number of times to repeat the character.
//
// Returns:
//   - []rune: The repeated character. Returns nil if count is less than 0.
func Repeat(char rune, count int) []rune {
	if count < 0 {
		return nil
	} else if count == 0 {
		return []rune{}
	}

	chars := make([]rune, 0, count)

	for i := 0; i < count; i++ {
		chars = append(chars, char)
	}

	return chars
}

// FixTabSize fixes the tab size by replacing it with a specified rune iff
// the tab size is greater than 0. The replacement rune is repeated for the
// specified number of times.
//
// Parameters:
//   - size: The size of the tab.
//   - rep: The replacement rune.
//
// Returns:
//   - []rune: The fixed tab size.
func FixTabSize(size int, rep rune) []rune {
	if size <= 0 {
		return []rune{'\t'}
	}

	return Repeat(rep, size)
}

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
