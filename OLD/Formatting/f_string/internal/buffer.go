package internal

import (
	"errors"
	"strings"
	"unicode/utf8"

	gcf "github.com/PlayerR9/go-commons/OLD/fixer"
	gcint "github.com/PlayerR9/go-commons/OLD/ints"
	gcers "github.com/PlayerR9/go-commons/errors"
)

const (
	// NBSP is the non-breaking space rune.
	NBSP rune = '\u00A0'
)

// Buffer is a type that represents a Buffer of a document.
type Buffer struct {
	// pages are the pages of the buffer.
	pages [][]*sectionBuilder

	// buff is the in-progress section of the buffer.
	buff *sectionBuilder

	// last_page is the last page of the buffer.
	last_page int
}

// Cleanup implements the object.Cleaner interface method.
func (b *Buffer) Clean() {
	if b == nil {
		return
	}

	if len(b.pages) > 0 {
		// pages are the pages of the buffer.
		for i := 0; i < len(b.pages); i++ {
			slice := gcf.CleanSlice(b.pages[i])
			b.pages[i] = slice
			b.pages[i] = nil
		}

		b.pages = nil
	}

	if b.buff != nil {
		b.buff.Cleanup()
		b.buff = nil
	}
}

// NewBuffer creates a new buffer.
//
// Returns:
//   - *Buffer: The new buffer. Never returns nil.
func NewBuffer() *Buffer {
	b := &Buffer{
		pages:     [][]*sectionBuilder{{}},
		buff:      nil,
		last_page: 0,
	}

	return b
}

// IsFirstOfLine is a function that returns true if the current position is the first
// position of a line.
//
// Returns:
//   - bool: True if the current position is the first position of a line, false otherwise.
//
// If the receiver is nil, this function returns true.
func (b *Buffer) IsFirstOfLine() bool {
	if b.buff == nil {
		return true
	}

	ok, _ := b.buff.is_first_of_line()
	return ok
}

// ForceWriteString is a function that writes a string to the buffer without
// checking for special characters. However, it does check for empty strings
// and will not write them.
//
// Parameters:
//   - str: The string to write.
//
// Returns:
//   - bool: True if the receiver is not nil, false otherwise.
func (b *Buffer) ForceWriteString(str string) bool {
	if b == nil {
		return false
	}

	if str == "" {
		return true
	}

	if b.buff == nil {
		b.buff = new_section_builder()
	}

	_ = b.buff.write_string(str)

	return true
}

// WriteRune is a private function that appends a rune to the buffer
// without checking for special characters.
//
// Parameters:
//   - r: The rune to append.
//
// Returns:
//   - bool: True if the receiver is not nil, false otherwise.
func (b *Buffer) WriteRune(r rune) bool {
	if b == nil {
		return false
	}

	if b.buff == nil {
		b.buff = new_section_builder()
	}

	_ = b.buff.write_rune(r)

	return true
}

// Write is a private function that appends a rune to the buffer
// while dealing with special characters.
//
// Parameters:
//   - char: The rune to append.
//
// Returns:
//   - bool: True if the receiver is not nil, false otherwise.
func (b *Buffer) Write(char rune) bool {
	if b == nil {
		return false
	}

	switch char {
	case '\t':
		// Tab : Add spaces until the next tab stop
		if b.buff == nil {
			b.buff = new_section_builder()
		}

		_ = b.buff.write_rune(char) // deal with this in later stages
	case '\v':
		// vertical tab : Add vertical tabulation

		// Do nothing
	case '\r', '\n', '\u0085':
		// carriage return : Move to the start of the line (alone)
		// or move to the start of the line and down (with line feed)
		// line feed : Add a new line or move to the left edge and down

		_ = b.accept()
	case '\f':
		// form feed : Go to the next page
		_ = b.accept()

		b.last_page++
		b.pages = append(b.pages, []*sectionBuilder{})
	case ' ':
		// Space
		if b.buff != nil {
			_ = b.buff.accept_word()
		}
	case '\u0000', '\a':
		// null : Ignore this character
		// Bell : Ignore this character
	case '\b':
		// backspace : Remove the last character
		if b.buff != nil {
			ok := b.buff.remove_one()
			if ok {
				return true
			}
		}

		for i := b.last_page; i >= 0; i-- {
			sections := b.pages[i]

			for j := len(sections) - 1; j >= 0; j-- {
				section := sections[j]

				ok := section.remove_one()
				if ok {
					return true
				}
			}
		}
	case '\u001A':
		// Control-Z : End of file for Windows text-mode file i/o
		b.finalize()
	case '\u001B':
		// escape : Introduce an escape sequence (next character)
		// Do nothing
	default:
		// NBSP : Non-breaking space
		// any other normal character
		if b.buff == nil {
			b.buff = new_section_builder()
		}

		if char == NBSP {
			// Non-breaking space
			_ = b.buff.write_rune(' ')
		} else {
			_ = b.buff.write_rune(char)
		}
	}

	return true
}

// AcceptLine is a function that accepts the current line of the formatted string.
// However, it does not accept the line if the line is empty.
//
// Parameters:
//   - right_delim: The right delimiter to use for the line.
//
// Returns:
//   - bool: True if the receiver is not nil, false otherwise.
func (b *Buffer) AcceptLine(right_delim string) bool {
	if b == nil {
		return false
	}

	if b.buff != nil {
		_ = b.buff.may_accept(right_delim)
	}

	return true
}

// AcceptLine is a function that accepts the current line of the formatted string.
// However, it does not accept the line if the line is empty.
//
// Parameters:
//   - right_delim: The right delimiter to use for the line.
//
// Returns:
//   - bool: True if the receiver is not nil, false otherwise.
func (b *Buffer) ForceAcceptLine(right_delim string) bool {
	if b == nil {
		return false
	}

	if b.buff == nil {
		b.buff = new_section_builder()
	}

	_ = b.buff.accept(right_delim)

	return true
}

// WriteEmptyLine is a function that accepts the current line
// regardless of the whether the line is empty or not.
//
// Parameters:
//   - right_delim: The right delimiter to use for the line.
//
// Returns:
//   - bool: True if the receiver is not nil, false otherwise.
func (b *Buffer) WriteEmptyLine(right_delim string) bool {
	if b == nil {
		return false
	}

	if b.buff == nil {
		b.buff = new_section_builder()
	}

	_ = b.buff.accept(right_delim)

	return true
}

// AcceptWord is a function that accepts the current word of the formatted string
// when the word is not empty.
//
// Returns:
//   - bool: True if the receiver is not nil, false otherwise.
func (b *Buffer) AcceptWord() bool {
	if b == nil {
		return false
	}

	if b.buff != nil {
		_ = b.buff.accept_word()
	}

	return true
}

// WriteBytes is a function that writes bytes to the formatted string.
//
// Parameters:
//   - b: The bytes to write.
//
// Returns:
//   - int: The number of bytes written.
//   - error: An error if one occurred.
//
// Errors:
//   - errors.NilReceiver if the receiver is nil.
//   - *ints.ErrAt if the data is not properly UTF-8 encoded.
func (b *Buffer) WriteBytes(data []byte) (int, error) {
	if b == nil {
		return 0, gcers.NilReceiver
	}

	if len(data) == 0 {
		return 0, nil
	}

	var count int

	for count = 0; len(data) > 0; count++ {
		r, size := utf8.DecodeRune(data)
		if r == utf8.RuneError {
			return count, gcint.NewErrAt(count+1, "byte", errors.New("invalid UTF-8 encoding"))
		}

		_ = b.Write(r)

		data = data[size:]
	}

	return count, nil
}

// GetPages returns the pages that are in the buffer.
//
// Parameters:
//   - tabSize: The size of the tab.
//   - fieldSpacing: The spacing to use for the tab stop.
//
// Returns:
//   - [][][][]string: The pages of the StdPrinter.
func (b *Buffer) GetPages(tabSize int, fieldSpacing int) [][][][]string {
	if b == nil {
		return nil
	}

	b.finalize()
	spacing := strings.Repeat(" ", fieldSpacing)

	pages := b.pages

	allStrings := make([][][][]string, 0, len(pages))

	for _, page := range pages {
		sectionLines := make([][][]string, 0)

		for _, section := range page {
			lines := section.get_lines()

			for i := 0; i < len(lines); i++ {
				line := lines[i]

				var level int
				for j := 0; j < len(line); j++ {
					line[j], level = fix_tab_stop(level, tabSize, spacing, line[j])
				}
			}

			sectionLines = append(sectionLines, lines)
		}

		allStrings = append(allStrings, sectionLines)
	}

	return allStrings
}

////////////////////////////////////////////////////////////////////////////////

// fix_tab_stop is a private function that fixes the tab stops in a string.
//
// The initial level and the integer return are used for
// chaining multiple calls to this function.
//
// Parameters:
//   - init: The initial level of the tab stop.
//   - tabSize: The size of the tab.
//   - spacing: The spacing to use for the tab stop.
//   - str: The string to fix the tab stops for.
//
// Returns:
//   - string: The string with the tab stops fixed.
//   - int: The total number of characters in the string.
//
// Behaviors:
//   - If the tabSize is less than 1, it is set to 1.
//   - If the initial tab stop is less than 0, it is set to 0.
//   - If the string is empty, an empty string is returned.
func fix_tab_stop(init, tab_size int, spacing, str string) (string, int) {
	total := init
	init = tab_size - (init % tab_size)

	if tab_size < 1 {
		tab_size = 1
	}

	var builder strings.Builder

	if init > 0 {
		init = tab_size - (init % tab_size)

		for i := init; i > 0; i-- {
			r, size := utf8.DecodeRuneInString(str)
			str = str[size:]

			if r == '\t' {
				repStr := strings.Repeat(spacing, i)
				total += i

				builder.WriteString(repStr)
				break
			} else {
				total++
				builder.WriteRune(r)
			}

			if len(str) == 0 {
				return builder.String(), total
			}
		}
	}

	for {
		for i := tab_size; i > 0; i-- {
			r, size := utf8.DecodeRuneInString(str)
			str = str[size:]

			if r == '\t' {
				repStr := strings.Repeat(spacing, i)
				total += i

				builder.WriteString(repStr)
				break
			} else {
				total++
				builder.WriteRune(r)
			}

			if len(str) == 0 {
				return builder.String(), total
			}
		}
	}
}

// Accept is a function that accepts the current in-progress buffer
// by converting it to the specified section type. Lastly, the section
// is added to the page.
//
// Parameters:
//   - sectionType: The section type to convert the buffer to.
//
// Returns:
//   - bool: True if the receiver is nil, false otherwise.
//
// Behaviors:
//   - Even when the buffer is empty, the section is still added to the page.
//     To avoid this, use the Finalize function.
func (b *Buffer) accept() bool {
	if b == nil {
		return false
	}

	if b.buff != nil {
		_ = b.buff.accept_word()
	}

	b.pages[b.last_page] = append(b.pages[b.last_page], b.buff)

	b.buff = nil

	return true
}

// finalize is a private function that finalizes the buffer.
func (b *Buffer) finalize() {
	if b.buff == nil {
		return
	}

	_ = b.buff.accept_word()

	b.pages[b.last_page] = append(b.pages[b.last_page], b.buff)

	b.buff = nil
}
