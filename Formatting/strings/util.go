package strings

import (
	"strings"
	"unicode/utf8"

	gcers "github.com/PlayerR9/go-commons/errors"
)

// find_tab_stop is a helper function to find the next tab stop for a string.
//
// Parameters:
//   - s: The string to find the tab stop for.
//   - tab_size: The size of the tab.
//
// Returns:
//   - int: The tab stop.
func find_tab_stop(s string, tab_size int) int {
	s = strings.TrimRight(s, " ")

	count := utf8.RuneCountInString(s)

	return tab_size * ((count / tab_size) + 1)
}

// pad_right is a helper function to pad a string to the right.
//
// Parameters:
//   - s: The string to pad.
//   - length: The length to pad the string to.
//
// Returns:
//   - string: The padded string.
func pad_right(s string, length int) string {
	var builder strings.Builder

	builder.WriteString(s)
	builder.WriteString(strings.Repeat(" ", length-utf8.RuneCountInString(s)))

	return builder.String()
}

// TabAlign aligns the tabs of a table's column.
//
// Parameters:
//   - table: The table to align.
//   - column: The column to align.
//   - tab_size: The size of the tab.
//
// Returns:
//   - [][]string: The aligned table.
//   - error: An error of type *errors.ErrInvalidParameter if the tab_size is less than 1
//     or the column is less than 0.
//
// Behaviors:
//   - If the column is not found in the table, the table is returned as is.
func TabAlign(table [][]string, column int, tab_size int) ([][]string, error) {
	if tab_size < 1 {
		return nil, gcers.NewErrInvalidParameter("tab_size", gcers.NewErrGT(0))
	} else if column < 0 {
		return nil, gcers.NewErrInvalidParameter("column", gcers.NewErrGTE(0))
	}

	seen := make(map[int]bool)

	for i := 0; i < len(table); i++ {
		if len(table[i]) > column {
			seen[i] = true
		}
	}

	if len(seen) == 0 {
		return table, nil
	}

	stops := make(map[int]int)

	for k := range seen {
		table[k][column] = strings.TrimRight(table[k][column], " ")

		stops[k] = find_tab_stop(table[k][column], tab_size)
	}

	max := -1

	for _, val := range stops {
		if max == -1 || val > max {
			max = val
		}
	}

	for k := range seen {
		table[k][column] = pad_right(table[k][column], max)
	}

	return table, nil
}

// LongestLine finds the longest line in a table.
//
// Parameters:
//   - table: The table to find the longest line in.
//
// Returns:
//   - int: The length of the longest line. -1 if the table is empty.
func LongestLine[T any](table [][]T) int {
	if len(table) == 0 {
		return -1
	}

	max := -1

	for i := 0; i < len(table); i++ {
		if len(table[i]) > max {
			max = len(table[i])
		}
	}

	return max
}

// TableEntriesAlign aligns the entries of a table.
//
// Parameters:
//   - table: The table to align.
//   - tabSize: The size of the tab.
//
// Returns:
//   - [][]string: The aligned table.
//   - error: An error if there was an issue aligning the table.
//
// Errors:
//   - *errors.ErrAt: If there was an issue aligning a specific column.
//   - *errors.ErrInvalidParameter: If the tabSize is less than 1.
func TableEntriesAlign(table [][]string, tabSize int) ([][]string, error) {
	if tabSize < 1 {
		return nil, gcers.NewErrInvalidParameter("tabSize", gcers.NewErrGT(0))
	}

	width := LongestLine(table)
	if width == -1 {
		return table, nil
	}

	var err error

	for i := 0; i < width; i++ {
		table, err = TabAlign(table, i, tabSize)
		if err != nil {
			return nil, gcers.NewErrAt(i+1, "column", err)
		}
	}

	return table, nil
}
