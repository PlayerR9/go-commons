package strings

import (
	"slices"
	"strings"
	"text/tabwriter"

	gcers "github.com/PlayerR9/go-commons/errors"
)

// TableAligner is a struct to help align tables.
type TableAligner struct {
	// head is the head of the table.
	head string

	// table is the table.
	table [][]string

	// idxs is the indexes to align.
	idxs []int
}

// NewTableAligner creates a new table aligner.
func NewTableAligner() *TableAligner {
	return &TableAligner{
		table: make([][]string, 0),
		idxs:  make([]int, 0),
	}
}

// SetHead sets the head of the table.
//
// Parameters:
//   - head: The head of the table.
//
// Returns:
//   - bool: True if the receiver is not nil, false otherwise.
func (ta *TableAligner) SetHead(head string) bool {
	if ta == nil {
		return false
	}

	ta.head = head

	return true
}

// AddRow adds a row to the table.
//
// Parameters:
//   - elems: The elements of the row.
//
// Returns:
//   - bool: True if the receiver is not nil, false otherwise.
func (ta *TableAligner) AddRow(elems ...string) bool {
	if ta == nil {
		return false
	}

	if len(elems) == 0 {
		ta.table = append(ta.table, []string{""})
	} else {
		ta.table = append(ta.table, elems)
	}

	return true
}

// AlignColumn specifies a row to align.
//
// Parameters:
//   - idx: The index of the row to align.
//
// Returns:
//   - bool: True if the receiver is not nil, false otherwise.
//
// If 'idx' is less than zero, this function returns false. If 'idx' already exists
// this function does nothing.
func (ta *TableAligner) AlignColumn(idx int) bool {
	if ta == nil || idx < 0 {
		return false
	}

	pos, ok := slices.BinarySearch(ta.idxs, idx)
	if ok {
		return true
	}

	ta.idxs = slices.Insert(ta.idxs, pos, idx)

	return true
}

// Reset resets the table aligner for reuse.
func (ta *TableAligner) Reset() {
	if ta == nil {
		return
	}

	if len(ta.table) > 0 {
		for i := 0; i < len(ta.table); i++ {
			for j := 0; j < len(ta.table[i]); j++ {
				ta.table[i][j] = ""
			}
			ta.table[i] = ta.table[i][:0]
			ta.table[i] = nil
		}

		ta.table = ta.table[:0]
	}

	ta.head = ""

	if len(ta.idxs) > 0 {
		for i := 0; i < len(ta.idxs); i++ {
			ta.idxs[i] = 0
		}
		ta.idxs = ta.idxs[:0]
	}
}

// Accept accepts the table aligner.
//
// Parameters:
//   - tab_size: The size of the tab.
//   - table_indent: Whether to indent the table.
//
// Returns:
//   - error: An error of type *errors.ErrInvalidParameter if the
//     tabSize is less than 1.
func (ta TableAligner) Build(tab_size int, table_indent bool) ([]string, error) {
	if tab_size < 1 {
		return nil, gcers.NewErrInvalidParameter("tab_size", gcers.NewErrGT(0))
	}

	// Add the table indent if needed.
	if table_indent {
		var builder strings.Builder

		for i := 0; i < len(ta.table); i++ {
			if len(ta.table[i]) == 0 {
				ta.table[i] = append(ta.table[i], "\t")
			} else {
				builder.WriteString("\t")
				builder.WriteString(ta.table[i][0])

				ta.table[i][0] = builder.String()
				builder.Reset()
			}
		}
	}

	var builder strings.Builder

	w := tabwriter.NewWriter(&builder, tab_size+1, tab_size, 1, ' ', 0)

	for _, row := range ta.table {
		str := strings.Join(row, "\t")

		_, err := w.Write([]byte(str))
		if err != nil {
			return nil, err
		}
	}

	err := w.Flush()
	if err != nil {
		return nil, err
	}

	lines := strings.Split(builder.String(), "\n")
	return lines, nil
}
