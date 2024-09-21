package bytes

import (
	"bytes"

	gcers "github.com/PlayerR9/go-commons/errors"
)

var (
	// Newline is the newline character in bytes.
	Newline []byte
)

func init() {
	Newline = []byte("\n")
}

// LimitReverseLines is a function that limits the lines of the data in reverse order.
//
// Parameters:
//   - data: The data to limit.
//   - limit: The limit of the lines.
//
// Returns:
//   - []byte: The limited data.
func LimitReverseLines(data []byte, limit int) []byte {
	if len(data) == 0 {
		return nil
	}

	lines := bytes.Split(data, Newline)

	if limit == -1 || limit > len(lines) {
		limit = len(lines)
	}

	start_idx := len(lines) - limit

	lines = lines[start_idx:]

	return bytes.Join(lines, Newline)
}

// LimitLines is a function that limits the lines of the data.
//
// Parameters:
//   - data: The data to limit.
//   - limit: The limit of the lines.
//
// Returns:
//   - []byte: The limited data.
func LimitLines(data []byte, limit int) []byte {
	if len(data) == 0 {
		return nil
	}

	lines := bytes.Split(data, Newline)

	if limit == -1 || limit > len(lines) {
		limit = len(lines)
	}

	lines = lines[:limit]

	return bytes.Join(lines, Newline)
}

// FirstIndex returns the first index of the target in the tokens.
//
// Parameters:
//   - tokens: The slice of tokens in which to search for the target.
//   - target: The target to search for.
//
// Returns:
//   - int: The index of the target. -1 if the target is not found.
//
// If either tokens or the target are empty, it returns -1.
func FirstIndex(tokens [][]byte, target []byte) int {
	if len(tokens) == 0 || len(target) == 0 {
		return -1
	}

	for i, token := range tokens {
		if bytes.Equal(token, target) {
			return i
		}
	}

	return -1
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
//   - *luc.ErrInvalidParameter: If the closingToken is an empty string.
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
func FindContentIndexes(op_token, cl_token []byte, tokens [][]byte) (result [2]int, err error) {
	result[0] = -1
	result[1] = -1

	if len(cl_token) == 0 {
		err := gcers.NewErrInvalidParameter("cl_token must not be empty")
		err.AddFrame("bytes", "FindContentIndexes()")

		return result, err
	}

	op_tok_idx := FirstIndex(tokens, op_token)
	if op_tok_idx == -1 {
		err = NewErrTokenNotFound(op_token, true)
		return
	}

	result[0] = op_tok_idx + 1

	balance := 1
	cl_tok_idx := -1

	for i := result[0]; i < len(tokens) && cl_tok_idx == -1; i++ {
		curr_tok := tokens[i]

		if bytes.Equal(curr_tok, cl_token) {
			balance--

			if balance == 0 {
				cl_tok_idx = i
			}
		} else if bytes.Equal(curr_tok, op_token) {
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
	} else if balance != 1 || bytes.Equal(cl_token, Newline) {
		err = NewErrTokenNotFound(cl_token, false)
		return
	}

	result[1] = len(tokens)
	return
}

// TrimEmpty removes empty bytes from a slice of bytes; including any empty
// spaces at the beginning and end of the bytes.
//
// Parameters:
//   - values: The values to trim.
//
// Returns:
//   - [][]byte: The trimmed values.
func TrimEmpty(values [][]byte) [][]byte {
	if len(values) == 0 {
		return values
	}

	var top int

	for i := 0; i < len(values); i++ {
		current_value := values[i]

		res := bytes.TrimSpace(current_value)
		if len(res) > 0 {
			values[top] = res
			top++
		}
	}

	values = values[:top]

	return values
}

// FixTabSize fixes the tab size by replacing it with a specified bytes iff
// the tab size is greater than 0. The replacement bytes are repeated for the
// specified number of times.
//
// Parameters:
//   - size: The size of the tab.
//   - rep: The replacement bytes.
//
// Returns:
//   - []byte: The fixed tab size.
func FixTabSize(size int, rep []byte) []byte {
	if size <= 0 {
		return []byte{'\t'}
	}

	return bytes.Repeat(rep, size)
}

// FilterNonEmpty removes empty bytes from a slice of bytes.
//
// Parameters:
//   - values: The slice of bytes to trim.
//
// Returns:
//   - [][]byte: The slice of bytes with empty bytes removed.
func FilterNonEmpty(values [][]byte) [][]byte {
	if len(values) == 0 {
		return nil
	}

	var top int

	for i := 0; i < len(values); i++ {
		if len(values[i]) > 0 {
			values[top] = values[i]
			top++
		}
	}

	return values[:top:top]
}

// DetermineCoords is a helper function that determines the coordinates of the given position.
//
// Parameters:
//   - data: The data read from the input stream.
//   - pos: The position of the faulty token.
//
// Returns:
//   - int: The x coordinate of the faulty token.
//   - int: The y coordinate of the faulty token.
func DetermineCoords(data []byte, pos int) (int, int) {
	if len(data) == 0 {
		return 0, 0
	}

	if pos < 0 {
		pos = len(data) + pos
	} else if pos >= len(data) {
		pos = len(data) - 1
	}

	var x int
	var y int

	for i := 0; i < pos; i++ {
		if data[i] == '\n' {
			x = 0
			y++
		} else {
			x++
		}
	}

	return x, y
}
