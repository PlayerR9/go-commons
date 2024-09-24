package f_string

import (
	"fmt"
	"io"
	"strings"
	"sync"

	"github.com/PlayerR9/go-commons/OLD/Formatting/f_string/internal"
	gcch "github.com/PlayerR9/go-commons/runes"
	gcers "github.com/PlayerR9/go-errors"
	"github.com/dustin/go-humanize"
)

var (
	// NBSP is the non-breaking space rune.
	NBSP rune
)

func init() {
	NBSP = internal.NBSP
}

/////////////////////////////////////////////////

// Traversor is a type that represents a traversor for a formatted string.
type Traversor struct {
	// config is the configuration of the traversor.
	config *FormatConfig

	// indentation is the string that is used for indentation
	// on the left side of the traversor.
	indentation string

	// hasIndent is a flag that indicates if the traversor has indentation.
	hasIndent bool

	// leftConfig is the configuration for the left symbol of the traversor.
	leftConfig *DelimiterConfig

	// rightDelim is the right delimiter of the traversor.
	rightDelim string

	// source is the buffer of the traversor.
	source *internal.Buffer

	// mu is the mutex of the traversor.
	mu *sync.Mutex
}

// Cleanup implements the Cleaner interface.
func (trav *Traversor) Clean() {
	if trav == nil {
		return
	}

	trav.source = nil
}

// newTraversor creates a new traversor.
//
// Parameters:
//   - config: The configuration of the traversor.
//   - source: The source of the traversor.
//
// Returns:
//   - *Traversor: The new traversor.
//   - errors: An error of type *errors.ErrInvalidParameter if 'config' or 'source' are nil.
func newTraversor(config *FormatConfig, source *internal.Buffer) (*Traversor, error) {
	if config == nil {
		return nil, gcers.NewErrNilParameter("config")
	} else if source == nil {
		return nil, gcers.NewErrNilParameter("source")
	}

	trav := &Traversor{
		config:      config,
		source:      source, // shared source
		hasIndent:   false,
		indentation: "",
		leftConfig:  nil,
		rightDelim:  "",
	}

	indentConfig := config.indentation
	if indentConfig != nil {
		trav.indentation = indentConfig.GetIndentation()
		trav.hasIndent = true
	}

	leftConfig := config.delimiterLeft
	if leftConfig != nil {
		trav.leftConfig = leftConfig
	}

	rightConfig := config.delimiterRight
	if rightConfig != nil {
		trav.rightDelim = rightConfig.str
	}

	return trav, nil
}

// writeIndent writes the indentation string to the traversor if
// the traversor has indentation and the traversor is at the first
// of the line.
//
// Returns:
//   - bool: True if the receiver is not nil, false otherwise.
func (trav *Traversor) writeIndent() bool {
	if trav == nil {
		return false
	}

	ok := trav.source.IsFirstOfLine()
	if !ok {
		return true
	}

	if trav.hasIndent {
		_ = trav.source.ForceWriteString(trav.indentation)
	}

	if trav.leftConfig != nil {
		_ = trav.source.ForceWriteString(trav.leftConfig.str)
	}

	return false
}

// writeRune appends a rune to the current, in-progress line of the traversor.
//
// Parameters:
//   - r: The rune to append.
//
// Returns:
//   - bool: True if the receiver is not nil, false otherwise.
func (trav *Traversor) writeRune(r rune) bool {
	if trav == nil {
		return false
	}

	_ = trav.writeIndent()

	if r == NBSP {
		_ = trav.source.WriteRune(r)
	} else {
		_ = trav.source.Write(r)
	}

	return true
}

// writeString appends a string to the current, in-progress line of the traversor.
//
// Parameters:
//   - str: The string to append.
//
// Returns:
//   - error: An error of this function failed.
//
// Errors:
//   - errors.NilReceiver if the receiver is not nil.
//   - *runes.ErrInvalidUTF8Encoding if the string is not valid UTF-8.
func (trav *Traversor) writeString(str string) error {
	if trav == nil {
		return nil
	}

	_ = trav.writeIndent()

	if str == "" {
		return nil
	}

	chars, err := gcch.StringToUtf8(str)
	if err != nil {
		return err
	}

	for _, r := range chars {
		_ = trav.source.Write(r)
	}

	return err
}

// writeLine writes a line to the traversor. If there is any in-progress line,
// then the line is appended to the line before accepting it. Otherwise, a new line
// with the line is added to the source.
//
// Parameters:
//   - line: The line to write.
//
// Returns:
//   - error: An error if the function failed.
//
// Errors:
//   - errors.NilReceiver if the receiver is nil.
//   - *runes.ErrInvalidUTF8Encoding if the string is not valid UTF-8.
//
// Behaviors:
//   - If line is empty, then an empty line is added to the source.
func (trav *Traversor) writeLine(line string) error {
	if trav == nil {
		return nil
	}

	_ = trav.source.AcceptLine(trav.rightDelim) // Accept the current line if any.

	_ = trav.writeIndent()

	if line == "" {
		_ = trav.source.WriteEmptyLine(trav.rightDelim)
	} else {
		chars, err := gcch.StringToUtf8(line)
		if err != nil {
			return err
		}

		for _, r := range chars {
			_ = trav.source.Write(r)
		}
	}

	_ = trav.source.AcceptLine(trav.rightDelim) // Accept the line.

	return nil
}

// AppendRune appends a rune to the half-line of the traversor.
//
// Parameters:
//   - r: The rune to append.
//
// Returns:
//   - bool: True if the receiver is not nil, false otherwise.
//
// Behaviors:
//   - If the half-line is nil, then a new half-line is created.
func (trav *Traversor) AppendRune(r rune) bool {
	if trav == nil {
		return false
	}

	if trav.source != nil {
		_ = trav.writeRune(r)
	}

	return true
}

// AppendString appends a string to the half-line of the traversor.
//
// Parameters:
//   - str: The string to append.
//
// Returns:
//   - error: An error if this function fails.
//
// Errors:
//   - errors.NilReceiver if the receiver is nil.
//   - *runes.ErrInvalidUTF8Encoding if the string is not valid UTF-8.
//
// Behaviors:
//   - IF str is empty: nothing is done.
func (trav *Traversor) AppendString(str string) error {
	if trav == nil {
		return nil
	}

	if trav.source == nil {
		return nil
	}

	err := trav.writeString(str)
	if err != nil {
		return err
	}

	return nil
}

// AppendStrings appends multiple strings to the half-line of the traversor.
//
// Parameters:
//   - strs: The strings to append.
//
// Returns:
//   - error: An error if this function fails.
//
// Errors:
//   - errors.NilReceiver if the receiver is nil.
//   - *ints.ErrAt if there is an error appending a string.
//
// Behaviors:
//   - This is equivalent to calling AppendString for each string in strs but more efficient.
func (trav *Traversor) AppendStrings(strs []string) error {
	if trav == nil {
		return nil
	}

	if trav.source == nil || len(strs) == 0 {
		return nil
	}

	for i, str := range strs {
		err := trav.writeString(str)
		if err != nil {
			return gcers.NewErrAt(humanize.Ordinal(i+1)+" string", err)
		}
	}

	return nil
}

// AppendJoinedString appends a joined string to the half-line of the traversor.
//
// Parameters:
//   - sep: The separator to use.
//   - fields: The fields to join.
//
// Returns:
//   - error: An error if this function fails.
//
// Errors:
//   - errors.NilReceiver if the receiver is nil.
//   - *runes.ErrInvalidUTF8Encoding if some field or the separator is not valid UTF-8 encoding.
//
// Behaviors:
//   - This is equivalent to calling AppendString(strings.Join(fields, sep)).
func (trav *Traversor) AppendJoinedString(sep string, fields ...string) error {
	if trav == nil {
		return nil
	}

	if trav.source == nil || len(fields) == 0 {
		return nil
	}

	str := strings.Join(fields, sep)

	err := trav.writeString(str)
	if err != nil {
		return err
	}

	return nil
}

// AcceptWord is a function that, if there is any in-progress word, then said word is added
// to the source.
//
// Returns:
//   - bool: True if the receiver is not nil, false otherwise.
func (trav *Traversor) AcceptWord() bool {
	if trav == nil {
		return false
	}

	if trav.source == nil {
		return true
	}

	_ = trav.source.AcceptWord()

	return true
}

// AcceptLine is a function that accepts the current line of the traversor.
//
// Returns:
//   - bool: True if the receiver is not nil, false otherwise.
//
// Behaviors:
//   - This also accepts the current word if any.
func (trav *Traversor) AcceptLine() bool {
	if trav == nil {
		return false
	}

	if trav.source == nil {
		return true
	}

	_ = trav.source.AcceptLine(trav.rightDelim)

	return true
}

// AddLine adds a line to the traversor. If there is any in-progress line, then the line is
// appended to the line before accepting it. Otherwise, a new line with the line is added to
// the source.
//
// Parameters:
//   - line: The line to add.
//
// Returns:
//   - error: An error if this function fails.
//
// Errors:
//   - errors.NilReceiver if the receiver is nil.
//   - *ints.ErrAt if there is an error adding the line.
//
// Behaviors:
//   - If line is empty, then an empty line is added to the source.
func (trav *Traversor) AddLine(line string) error {
	if trav == nil {
		return nil
	}

	if trav.source == nil {
		return nil
	}

	err := trav.writeLine(line)
	if err != nil {
		return err
	}

	return nil
}

// AddLines adds multiple lines to the traversor in a more efficient way than
// adding each line individually.
//
// Parameters:
//   - lines: The lines to add.
//
// Returns:
//   - error: An error if this function fails.
//
// Errors:
//   - errors.NilReceiver if the receiver is nil.
//   - *ints.ErrAt if there is an error adding a line.
//
// Behaviors:
//   - If there are no lines, then nothing is done.
func (trav *Traversor) AddLines(lines []string) error {
	if trav == nil {
		return nil
	}

	if trav.source == nil || len(lines) == 0 {
		return nil
	}

	for i, line := range lines {
		err := trav.writeLine(line)
		if err != nil {
			return gcers.NewErrAt(humanize.Ordinal(i+1)+" line", err)
		}
	}

	return nil
}

// AddJoinedLine adds a joined line to the traversor. This is a more efficient way to do
// the same as AddLine(strings.Join(fields, sep)).
//
// Parameters:
//   - sep: The separator to use.
//   - fields: The fields to join.
//
// Returns:
//   - error: An error if this function fails.
//
// Error:
//   - errors.NilReceiver if the receiver is nil.
//   - *ints.ErrInvalidRuneAt if there is an invalid rune in the line.
//
// Behaviors:
//   - If fields is empty, then nothing is done.
func (trav *Traversor) AddJoinedLine(sep string, fields ...string) error {
	if trav == nil {
		return nil
	}

	if trav.source == nil || len(fields) == 0 {
		return nil
	}

	str := strings.Join(fields, sep)

	err := trav.writeLine(str)
	if err != nil {
		return err
	}

	return nil
}

// EmptyLine adds an empty line to the traversor. This is a more efficient way to do
// the same as AddLine("") or AddLines([]string{""}).
//
// Returns:
//   - bool: True if the receiver is not nil, false otherwise.
//
// Behaviors:
//   - If the half-line is not empty, then the half-line is added to the source
//     (half-line is reset) and an empty line is added to the source.
func (trav *Traversor) EmptyLine() bool {
	if trav == nil {
		return false
	}

	if trav.source == nil {
		return true
	}

	_ = trav.source.AcceptLine(trav.rightDelim) // Accept the current line if any.

	_ = trav.writeIndent()

	_ = trav.source.ForceAcceptLine(trav.rightDelim) // Accept the line.

	return true
}

// Write implements the io.Writer interface for the traversor.
func (trav *Traversor) Write(p []byte) (int, error) {
	if len(p) == 0 {
		return 0, nil
	} else if trav == nil {
		return 0, io.ErrShortWrite
	}

	if trav.source == nil {
		return 0, nil
	}

	n, err := trav.source.WriteBytes(p)
	if err != nil {
		return n, err
	}

	return len(p), nil
}

// Print is a function that writes to the traversor using the fmt.Fprint function.
//
// Parameters:
//   - a: The arguments to write.
//
// Returns:
//   - error: An error if this function fails.
//
// Error:
//   - errors.NilReceiver if the receiver is nil.
//   - any other error returned by the fmt.Fprint function.
func (trav *Traversor) Print(a ...interface{}) error {
	if trav == nil {
		return nil
	}

	if trav.source == nil {
		return nil
	}

	_, err := fmt.Fprint(trav, a...)
	return err
}

// Printf is a function that writes to the traversor using the fmt.Fprintf function.
//
// Parameters:
//   - format: The format string.
//   - a: The arguments to write.
//
// Returns:
//   - error: An error if this function fails.
//
// Error:
//   - errors.NilReceiver if the receiver is nil.
//   - any other error returned by the fmt.Fprintf function.
func (trav *Traversor) Printf(format string, a ...interface{}) error {
	if trav == nil {
		return nil
	}

	if trav.source == nil {
		return nil
	}

	_, err := fmt.Fprintf(trav, format, a...)
	return err
}

// Println is a function that writes to the traversor using the fmt.Fprintln function.
//
// Parameters:
//   - a: The arguments to write.
//
// Returns:
//   - error: An error if this function fails.
//
// Error:
//   - errors.NilReceiver if the receiver is nil.
//   - any other error returned by the fmt.Fprintln function.
func (trav *Traversor) Println(a ...interface{}) error {
	if trav == nil {
		return nil
	}

	if trav.source == nil {
		return nil
	}

	_, err := fmt.Fprintln(trav, a...)
	return err
}

// GetConfig is a method that returns a copy of the configuration of the traversor.
//
// Parameters:
//   - options: The options to apply to the configuration.
//
// Returns:
//   - *FormatConfig: A pointer to the copy of the configuration of the traversor.
//   - error: An error of type errors.NilReceiver if the receiver is nil.
func (trav Traversor) GetConfig(options ...ConfigOption) (*FormatConfig, error) {
	configCopy := trav.config.Copy()

	for _, option := range options {
		option(configCopy)
	}

	return configCopy, nil
}

// Lock locks the traversor. Be aware of deadlocks.
//
// Returns:
//   - bool: True if the receiver is not nil, false otherwise.
func (trav *Traversor) Lock() bool {
	if trav == nil {
		return false
	}

	trav.mu.Lock()

	return true
}

// Unlock unlocks the traversor. Be aware of deadlocks.
//
// Returns:
//   - bool: True if the receiver is not nil, false otherwise.
func (trav *Traversor) Unlock() bool {
	if trav == nil {
		return false
	}

	trav.mu.Unlock()

	return true
}

//////////////////////////////////////////////////////////////

/*
// GetIndent returns the indentation string of the traversor.
//
// Returns:
//   - string: The indentation string of the traversor.
func (trav *Traversor) GetIndent() string {
	if trav.indent == nil {
		return ""
	} else {
		return trav.indentStr
	}
}

// ApplyIndent applies the indentation configuration to a specified string.
//
// Parameters:
//   - str: The string to apply the indentation to.
//
// Returns:
//   - string: The string with the indentation applied.
func (trav *Traversor) ApplyIndent(isFirstLine bool, str string) string {
	if trav.indent == nil || !trav.source.isFirstOfLine() {
		return str
	}

	var builder strings.Builder

	builder.WriteString(trav.indentStr)
	builder.WriteString(str)

	return builder.String()
}
*/

/*
// AddMultiline adds a multiline to the traversor. But first, it accepts any in-progress
// half-line.
//
// Parameters:
//   - mlt: The multiline to add.
//
// Behaviors:
//   - If the multiline is nil, then nothing is done.
func (trav *Traversor) AddMultiline(mlt *cb.MultiLineText) {
	if mlt == nil {
		return
	}

	trav.AcceptHalfLine()
	trav.source.addLine(mlt)
}
*/
