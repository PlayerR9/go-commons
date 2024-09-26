package strings

import (
	"strings"
	"text/tabwriter"

	gcstr "github.com/PlayerR9/go-commons/strings"
	gers "github.com/PlayerR9/go-errors"
	gerr "github.com/PlayerR9/go-errors/error"
)

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
func TableEntriesAlign(table [][]string, tab_size int) ([]string, error) {
	if tab_size < 1 {
		return nil, gerr.New(gers.BadParameter, "tab_size must be positive")
	}

	var lb gcstr.LineBuffer

	w := tabwriter.NewWriter(&lb, tab_size+1, tab_size, 1, ' ', 0)

	for _, row := range table {
		data := strings.Join(row, "\t")

		_, err := w.Write([]byte(data))
		if err != nil {
			return nil, err
		}
	}

	err := w.Flush()
	if err != nil {
		return nil, err
	}

	lines := lb.LinesString()
	return lines, nil
}
