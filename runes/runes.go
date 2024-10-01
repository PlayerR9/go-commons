package runes

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
