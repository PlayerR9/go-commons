package f_string

import (
	"fmt"
	"strings"

	"github.com/PlayerR9/go-commons/OLD/Formatting/f_string/internal"
	gcers "github.com/PlayerR9/go-commons/errors"
)

// ApplyMany applies a format to a stringer.
//
// Parameters:
//   - p: The StdPrinter to use.
//   - elems: The elements to format.
//
// Returns:
//   - error: An error if the formatting fails.
//
// Errors:
//   - *ErrInvalidParameter: If the StdPrinter is nil.
//   - *ErrFinalization: If the finalization of the page fails.
//   - *Errors.ErrAt: If an error occurs on a specific element.
//
// Behaviors:
//   - If the formatter is nil, the function uses the nil formatter.
//   - If an element is nil, the function skips the element.
//   - If all elements are nil, the function does nothing.
func ApplyMany[T FStringer](trav *Traversor, elems []T) error {
	if len(elems) == 0 {
		return nil
	}

	for i, elem := range elems {
		err := elem.FString(trav)
		if err != nil {
			return gcers.NewErrAt(i+1, "FStringer element", err)
		}
	}

	return nil
}

// ApplyFuncMany applies a format function to the StdPrinter.
//
// Parameters:
//   - p: The StdPrinter to use.
//   - f: The function to apply.
//   - elems: The elements to apply the function to.
//
// Returns:
//   - error: An error if the function fails.
//
// Errors:
//   - *ErrInvalidParameter: If the StdPrinter is nil.
//   - *Errors.ErrAt: If an error occurs on a specific element.
//   - any error returned by the function.
func ApplyFuncMany[T any](trav *Traversor, f FStringFunc[T], elems []T) error {
	if len(elems) == 0 {
		return nil
	}

	for i, elem := range elems {
		err := f(trav, elem)
		if err != nil {
			return gcers.NewErrAt(i+1, "element", err)
		}
	}

	return nil
}

// StdPrinter is a type that represents a formatted string.
type StdPrinter struct {
	// buffer is the buffer of the document.
	buff *internal.Buffer

	// formatter is the formatter of the document.
	formatter *FormatConfig
}

// Cleanup implements the Cleaner interface.
func (p *StdPrinter) Clean() {
	if p == nil || p.buff == nil {
		return
	}

	p.buff.Clean()
	p.buff = nil
}

// NewStdPrinter creates a new StdPrinter.
//
// Parameters:
//   - form: The formatter to use.
//
// Returns:
//   - *StdPrinter: The new StdPrinter.
//   - *Traversor: The traversor of the StdPrinter.
//
// Behaviors:
//   - If the formatter is nil, the function uses the default formatter.
func NewStdPrinter(form *FormatConfig) (*StdPrinter, *Traversor) {
	if form == nil {
		form = DefaultFormatter.Copy()
	}

	p := &StdPrinter{
		buff:      internal.NewBuffer(),
		formatter: form,
	}

	trav, _ := newTraversor(p.formatter, p.buff)

	return p, trav
}

// NewStdPrinterFromConfig creates a new StdPrinter from a configuration.
//
// Parameters:
//   - opts: The configuration to use.
//
// Returns:
//   - *StdPrinter: The new StdPrinter.
//   - *Traversor: The traversor of the StdPrinter.
//
// Behaviors:
//   - If the configuration is nil, the function uses the default configuration.
//   - Panics if an invalid configuration type is given (i.e., not IndentConfig, DelimiterConfig,
//     or SeparatorConfig).
func NewStdPrinterFromConfig(opts ...any) (*StdPrinter, *Traversor) {
	p := &StdPrinter{
		buff:      internal.NewBuffer(),
		formatter: NewFormatter(opts...),
	}

	trav, _ := newTraversor(p.formatter, p.buff)

	return p, trav
}

// GetPages returns the pages of the StdPrinter.
//
// Returns:
//   - [][][][]string: The pages of the StdPrinter. Nil if the receiver is nil.
func (p *StdPrinter) GetPages() [][][][]string {
	if p == nil {
		return nil
	}

	tabSize, fieldSpacing := p.formatter.GetTabSize(), p.formatter.GetSpacingSize()

	pages := p.buff.GetPages(tabSize, fieldSpacing)

	// Reset the buffer
	p.buff = internal.NewBuffer()

	return pages
}

// SprintFString prints a formatted string.
//
// Parameters:
//   - form: The formatter to use.
//   - elem: The element to print.
//
// Returns:
//   - [][][][]string: The pages of the formatted string.
//   - error: An error if the printing fails.
//
// Behaviors:
//   - If the formatter is nil, the function uses the default formatter.
func SprintFString[T FStringer](form *FormatConfig, elem T) ([][][][]string, error) {
	if form == nil {
		form = DefaultFormatter.Copy()
	}

	buff := internal.NewBuffer()

	trav, _ := newTraversor(form, buff)

	err := elem.FString(trav)
	if err != nil {
		return nil, err
	}

	tabSize, fieldSpacing := form.GetTabSize(), form.GetSpacingSize()

	pages := buff.GetPages(tabSize, fieldSpacing)

	return pages, nil
}

// Sprint prints strings.
//
// Parameters:
//   - form: The formatter to use.
//   - strs: The strings to print.
//
// Returns:
//   - [][][][]string: The pages of the formatted strings.
//   - error: An error of type *common.ErrAt if the printing fails because of an invalid UTF-8 encoding.
//
// Behaviors:
//   - If the formatter is nil, the function uses the default formatter.
func Sprint(form *FormatConfig, strs ...string) ([][][][]string, error) {
	if len(strs) == 0 {
		return nil, nil
	}

	if form == nil {
		form = DefaultFormatter.Copy()
	}

	buff := internal.NewBuffer()
	trav, _ := newTraversor(form, buff)

	for i, str := range strs {
		err := trav.writeString(str)
		if err != nil {
			return nil, gcers.NewErrAt(i, "string", err)
		}
	}

	tabSize, fieldSpacing := form.GetTabSize(), form.GetSpacingSize()

	pages := buff.GetPages(tabSize, fieldSpacing)

	return pages, nil
}

// Sprintj prints a joined string.
//
// Parameters:
//   - form: The formatter to use.
//   - sep: The separator to use.
//   - strs: The strings to join.
//
// Returns:
//   - [][][][]string: The pages of the formatted strings.
//   - error: An error of type *runes.ErrInvalidUTF8Encoding if the printing fails because of an invalid UTF-8 encoding.
//
// Behaviors:
//   - If the formatter is nil, the function uses the default formatter.
func Sprintj(form *FormatConfig, sep string, strs ...string) ([][][][]string, error) {
	if len(strs) == 0 {
		return nil, nil
	}

	if form == nil {
		form = DefaultFormatter.Copy()
	}

	buff := internal.NewBuffer()
	trav, _ := newTraversor(form, buff)

	str := strings.Join(strs, sep)

	err := trav.writeString(str)
	if err != nil {
		return nil, err
	}

	tabSize, fieldSpacing := form.GetTabSize(), form.GetSpacingSize()

	pages := buff.GetPages(tabSize, fieldSpacing)

	return pages, nil
}

// Sfprint prints a formatted string.
//
// Parameters:
//   - form: The formatter to use.
//   - a: The elements to print.
//
// Returns:
//   - [][][][]string: The pages of the formatted strings.
//   - error: An error if the printing fails.
//
// Behaviors:
//   - If the formatter is nil, the function uses the default formatter.
func Sfprint(form *FormatConfig, a ...interface{}) ([][][][]string, error) {
	if len(a) == 0 {
		return nil, nil
	}

	if form == nil {
		form = DefaultFormatter.Copy()
	}

	buff := internal.NewBuffer()
	trav, _ := newTraversor(form, buff)

	_, err := fmt.Fprint(trav, a...)
	if err != nil {
		return nil, err
	}

	tabSize, fieldSpacing := form.GetTabSize(), form.GetSpacingSize()

	pages := buff.GetPages(tabSize, fieldSpacing)

	return pages, nil
}

// Sfprintf prints a formatted string.
//
// Parameters:
//   - form: The formatter to use.
//   - format: The format string.
//   - a: The elements to print.
//
// Returns:
//   - [][][][]string: The pages of the formatted strings.
//   - error: An error if the printing fails.
//
// Behaviors:
//   - If the formatter is nil, the function uses the default formatter.
func Sfprintf(form *FormatConfig, format string, a ...interface{}) ([][][][]string, error) {
	if len(a) == 0 {
		return nil, nil
	}

	if form == nil {
		form = DefaultFormatter.Copy()
	}

	buff := internal.NewBuffer()
	trav, _ := newTraversor(form, buff)

	_, err := fmt.Fprintf(trav, format, a...)
	if err != nil {
		return nil, err
	}

	tabSize, fieldSpacing := form.GetTabSize(), form.GetSpacingSize()

	pages := buff.GetPages(tabSize, fieldSpacing)

	return pages, nil
}

// Sprintln prints a string with a newline.
//
// Parameters:
//   - form: The formatter to use.
//   - lines: The lines to print.
//
// Returns:
//   - [][][][]string: The pages of the formatted strings.
//   - error: An error if the printing fails.
//
// Behaviors:
//   - If the formatter is nil, the function uses the default formatter.
func Sprintln(form *FormatConfig, lines ...string) ([][][][]string, error) {
	if len(lines) == 0 {
		return nil, nil
	}

	if form == nil {
		form = DefaultFormatter.Copy()
	}

	buff := internal.NewBuffer()
	trav, _ := newTraversor(form, buff)

	for i, line := range lines {
		err := trav.writeLine(line)
		if err != nil {
			return nil, gcers.NewErrAt(i, "line", err)
		}
	}

	tabSize, fieldSpacing := form.GetTabSize(), form.GetSpacingSize()

	pages := buff.GetPages(tabSize, fieldSpacing)

	return pages, nil
}

// Sprintjln prints a joined string with a newline.
//
// Parameters:
//   - form: The formatter to use.
//   - sep: The separator to use.
//   - lines: The lines to join.
//
// Returns:
//   - [][][][]string: The pages of the formatted strings.
//   - error: An error if the printing fails.
//
// Behaviors:
//   - If the formatter is nil, the function uses the default formatter.
func Sprintjln(form *FormatConfig, sep string, lines ...string) ([][][][]string, error) {
	if len(lines) == 0 {
		return nil, nil
	}

	if form == nil {
		form = DefaultFormatter.Copy()
	}

	buff := internal.NewBuffer()
	trav, _ := newTraversor(form, buff)

	str := strings.Join(lines, sep)

	err := trav.writeLine(str)
	if err != nil {
		return nil, err
	}

	tabSize, fieldSpacing := form.GetTabSize(), form.GetSpacingSize()

	pages := buff.GetPages(tabSize, fieldSpacing)

	return pages, nil
}

// Sfprintln prints a formatted string with a newline.
//
// Parameters:
//   - form: The formatter to use.
//   - a: The elements to print.
//
// Returns:
//   - [][][][]string: The pages of the formatted strings.
//   - error: An error if the printing fails.
//
// Behaviors:
//   - If the formatter is nil, the function uses the default formatter.
func Sfprintln(form *FormatConfig, a ...interface{}) ([][][][]string, error) {
	if len(a) == 0 {
		return nil, nil
	}

	if form == nil {
		form = DefaultFormatter.Copy()
	}

	buff := internal.NewBuffer()
	trav, _ := newTraversor(form, buff)

	_, err := fmt.Fprintln(trav, a...)
	if err != nil {
		return nil, err
	}

	tabSize, fieldSpacing := form.GetTabSize(), form.GetSpacingSize()

	pages := buff.GetPages(tabSize, fieldSpacing)

	return pages, nil
}

/*
// FilePrinter is a type that represents a formatted string.
type FilePrinter struct {
	// buffer is the buffer of the document.
	buff *buffer

	// formatter is the formatter of the document.
	formatter FormatConfig

	// file is the file to write to.
	out io.Writer
}

// Cleanup implements the Cleaner interface.
func (p *FilePrinter) Clean() {
	p.buff.Clean()

	p.buff = nil
}

// NewFilePrinter creates a new FilePrinter.
//
// Parameters:
//   - out: The writer to use.
//   - form: The formatter to use.
//
// Returns:
//   - *FilePrinter: The new FilePrinter.
//
// Behaviors:
//   - If the writer is nil, the function uses os.Stdout.
func NewFilePrinter(out io.Writer, form FormatConfig) (*FilePrinter, *Traversor) {
	fp := &FilePrinter{
		buff:      pkg.NewBuffer(),
		formatter: form,
	}

	if out != nil {
		fp.out = out
	} else {
		fp.out = os.Stdout
	}

	trav := newTraversor(form, fp.buff)

	return fp, trav
}

// NewFilePrinterFromConfig creates a new FilePrinter from a configuration.
//
// Parameters:
//   - out: The writer to use.
//   - opts: The configuration to use.
//
// Returns:
//   - *FilePrinter: The new FilePrinter.
//   - *Traversor: The traversor of the FilePrinter.
//
// Behaviors:
//   - If the writer is nil, the function uses os.Stdout.
func NewFilePrinterFromConfig(out io.Writer, opts ...uc.Copier) (*FilePrinter, *Traversor) {
	fp := &FilePrinter{
		buff:      pkg.NewBuffer(),
		formatter: NewFormatter(opts...),
	}

	if out != nil {
		fp.out = out
	} else {
		fp.out = os.Stdout
	}

	trav := newTraversor(fp.formatter, fp.buff)

	return fp, trav
}

// Update updates the FilePrinter by writing the buffer to the file.
func (p *FilePrinter) Update() {
	tabSize, fieldSpacing := p.formatter.GetTabSize(), p.formatter.GetSpacingSize()

	pages := p.buff.getPages(tabSize, fieldSpacing)

	// Reset the buffer
	p.buff = pkg.NewBuffer()

	p.out.Write([]byte(strings.Join(Stringfy(pages), "\f")))
}

// FprintFString prints a formatted string.
//
// Parameters:
//   - out: The writer to use.
//   - form: The formatter to use.
//   - elem: The element to print.
//
// Returns:
//   - error: An error if the printing fails.
//
// Behaviors:
//   - If the writer is nil, the function uses os.Stdout.
func FprintFString[T FStringer](out io.Writer, form FormatConfig, elem T) error {
	buff := pkg.NewBuffer()
	trav := newTraversor(form, buff)

	err := elem.FString(trav)
	if err != nil {
		return err
	}

	if out == nil {
		out = os.Stdout
	}

	tabSize, fieldSpacing := form.GetTabSize(), form.GetSpacingSize()

	pages := buff.getPages(tabSize, fieldSpacing)

	out.Write([]byte(strings.Join(Stringfy(pages), "\f")))

	return nil
}

// Fprint prints strings.
//
// Parameters:
//   - out: The writer to use.
//   - form: The formatter to use.
//   - strs: The strings to print.
//
// Returns:
//   - error: An error if the printing fails.
//
// Behaviors:
//   - If the writer is nil, the function uses os.Stdout.
func Fprint(out io.Writer, form FormatConfig, strs ...string) error {
	if len(strs) == 0 {
		return nil
	}

	buff := pkg.NewBuffer()
	trav := newTraversor(form, buff)

	for i, str := range strs {
		err := trav.writeString(str)
		if err != nil {
			return luint.NewErrAt(i, "string", err)
		}
	}

	if out == nil {
		out = os.Stdout
	}

	tabSize, fieldSpacing := form.GetTabSize(), form.GetSpacingSize()

	pages := buff.getPages(tabSize, fieldSpacing)

	out.Write([]byte(strings.Join(Stringfy(pages), "\f")))

	return nil
}

// Fprintj prints a joined string.
//
// Parameters:
//   - out: The writer to use.
//   - form: The formatter to use.
//   - sep: The separator to use.
//   - strs: The strings to join.
//
// Returns:
//   - error: An error if the printing fails.
//
// Behaviors:
//   - If the writer is nil, the function uses os.Stdout.
func Fprintj(out io.Writer, form FormatConfig, sep string, strs ...string) error {
	if len(strs) == 0 {
		return nil
	}

	buff := pkg.NewBuffer()
	trav := newTraversor(form, buff)

	str := strings.Join(strs, sep)

	err := trav.writeString(str)
	if err != nil {
		return err
	}

	if out == nil {
		out = os.Stdout
	}

	tabSize, fieldSpacing := form.GetTabSize(), form.GetSpacingSize()

	pages := buff.getPages(tabSize, fieldSpacing)

	out.Write([]byte(strings.Join(Stringfy(pages), "\f")))

	return nil
}

// Ffprint prints a formatted string.
//
// Parameters:
//   - out: The writer to use.
//   - form: The formatter to use.
//   - a: The elements to print.
//
// Returns:
//   - error: An error if the printing fails.
//
// Behaviors:
//   - If the writer is nil, the function uses os.Stdout.
func Ffprint(out io.Writer, form FormatConfig, a ...interface{}) error {
	if len(a) == 0 {
		return nil
	}

	buff := pkg.NewBuffer()
	trav := newTraversor(form, buff)

	_, err := fmt.Fprint(trav, a...)
	if err != nil {
		return err
	}

	if out == nil {
		out = os.Stdout
	}

	tabSize, fieldSpacing := form.GetTabSize(), form.GetSpacingSize()

	pages := buff.getPages(tabSize, fieldSpacing)

	out.Write([]byte(strings.Join(Stringfy(pages), "\f")))

	return nil
}

// Ffprintf prints a formatted string.
//
// Parameters:
//   - out: The writer to use.
//   - form: The formatter to use.
//   - format: The format string.
//   - a: The elements to print.
//
// Returns:
//   - error: An error if the printing fails.
//
// Behaviors:
//   - If the writer is nil, the function uses os.Stdout.
func Ffprintf(out io.Writer, form FormatConfig, format string, a ...interface{}) error {
	if len(a) == 0 {
		return nil
	}

	buff := pkg.NewBuffer()
	trav := newTraversor(form, buff)

	_, err := fmt.Fprintf(trav, format, a...)
	if err != nil {
		return err
	}

	if out == nil {
		out = os.Stdout
	}

	tabSize, fieldSpacing := form.GetTabSize(), form.GetSpacingSize()

	pages := buff.getPages(tabSize, fieldSpacing)

	out.Write([]byte(strings.Join(Stringfy(pages), "\f")))

	return nil
}

// Fprintln prints a string with a newline.
//
// Parameters:
//   - out: The writer to use.
//   - form: The formatter to use.
//   - lines: The lines to print.
//
// Returns:
//   - error: An error if the printing fails.
//
// Behaviors:
//   - If the writer is nil, the function uses os.Stdout.
func Fprintln(out io.Writer, form FormatConfig, lines ...string) error {
	if len(lines) == 0 {
		return nil
	}

	buff := pkg.NewBuffer()
	trav := newTraversor(form, buff)

	for i, line := range lines {
		err := trav.writeLine(line)
		if err != nil {
			return luint.NewErrAt(i, "line", err)
		}
	}

	if out == nil {
		out = os.Stdout
	}

	tabSize, fieldSpacing := form.GetTabSize(), form.GetSpacingSize()

	pages := buff.getPages(tabSize, fieldSpacing)

	out.Write([]byte(strings.Join(Stringfy(pages), "\f")))

	return nil
}

// Fprintjln prints a joined string with a newline.
//
// Parameters:
//   - out: The writer to use.
//   - form: The formatter to use.
//   - sep: The separator to use.
//   - lines: The lines to join.
//
// Returns:
//   - error: An error if the printing fails.
//
// Behaviors:
//   - If the writer is nil, the function uses os.Stdout.
func Fprintjln(out io.Writer, form FormatConfig, sep string, lines ...string) error {
	if len(lines) == 0 {
		return nil
	}

	buff := pkg.NewBuffer()
	trav := newTraversor(form, buff)

	str := strings.Join(lines, sep)

	err := trav.writeLine(str)
	if err != nil {
		return err
	}

	if out == nil {
		out = os.Stdout
	}

	tabSize, fieldSpacing := form.GetTabSize(), form.GetSpacingSize()

	pages := buff.getPages(tabSize, fieldSpacing)

	out.Write([]byte(strings.Join(Stringfy(pages), "\f")))

	return nil
}

// Ffprintln prints a formatted string with a newline.
//
// Parameters:
//   - out: The writer to use.
//   - form: The formatter to use.
//   - a: The elements to print.
//
// Returns:
//   - error: An error if the printing fails.
//
// Behaviors:
//   - If the writer is nil, the function uses os.Stdout.
func Ffprintln(out io.Writer, form FormatConfig, a ...interface{}) error {
	if len(a) == 0 {
		return nil
	}

	buff := pkg.NewBuffer()
	trav := newTraversor(form, buff)

	_, err := fmt.Fprintln(trav, a...)
	if err != nil {
		return err
	}

	if out == nil {
		out = os.Stdout
	}

	tabSize, fieldSpacing := form.GetTabSize(), form.GetSpacingSize()

	pages := buff.getPages(tabSize, fieldSpacing)

	out.Write([]byte(strings.Join(Stringfy(pages), "\f")))

	return nil
}
*/

/////////////////////////////////////////////////

//////////////////////////////////////////////////////////////
/*
const (
	// Hellip is the ellipsis character.
	Hellip string = "..."

	// HellipLen is the length of the ellipsis character.
	HellipLen int = len(Hellip)

	// MarginLeft is the left margin of the content box.
	MarginLeft int = 1
)


// addLine is a private function that adds a line to the formatted string.
//
// Parameters:
//   - mlt: The line to add.
func (p *StdPrinterSource) addLine(mlt *cb.MultiLineText) {
	if mlt == nil {
		return
	}

	p.lines = append(p.lines, mlt)
}

// GetLines returns the lines of the formatted string.
//
// Returns:
//   - []*MultiLineText: The lines of the formatted string.
func (p *StdPrinterSource) GetLines() []*cb.MultiLineText {
	return p.lines
}

/*
func (p *StdPrinterSource) Boxed(width, height int) ([]string, error) {
	p.fix()

	all_fields := p.getAllFields()

	fss := make([]*StdPrinterSource, 0, len(all_fields))

	for _, fields := range all_fields {
		p := &StdPrinterSource{
			lines: fields,
		}

		fss = append(fss, p)
	}

	lines := make([]string, 0)

	for _, p := range fss {
		ts, err := p.generateContentBox(width, height)
		if err != nil {
			return nil, err
		}

		leftLimit, ok := ts.GetFurthestRightEdge()
		if !ok {
			panic("could not get furthest right edge")
		}

		for _, line := range ts.GetLines() {
			fitted, err := sext.FitString(line.String(), leftLimit)
			if err != nil {
				return nil, err
			}

			lines = append(lines, fitted)
		}
	}

	return lines, nil
}


func (p *StdPrinterSource) fix() {
	// 1. Fix newline boundaries
	newLines := make([]string, 0)

	for _, line := range p.lines {
		newFields := strings.Split(line, "\n")

		newLines = append(newLines, newFields...)
	}

	p.lines = newLines
}

// Must call Fix() before calling this function.
func (p *StdPrinterSource) getAllFields() [][]string {
	// TO DO: Handle special WHITESPACE characters

	fieldList := make([][]string, 0)

	for _, content := range p.lines {
		fields := strings.Fields(content)

		if len(fields) != 0 {
			fieldList = append(fieldList, fields)
		}
	}

	return fieldList
}
*/

/*
// GetDocument returns the content of the FieldSplitter as a Document.
//
// Returns:
//   - *Document: The content of the FieldSplitter.
func (p *FieldSplitter) GetDocument() *FieldSplitter {
	return p.content
}


// Build is a function that builds the document.
//
// Returns:
//   - *tld.Document: The built document.
func (p *FieldSplitter) Build() *tld.Document {
	doc := tld.NewDocument()

	for _, page := range p.content.pages {
		iter := page.Iterator()

		for {
			section, err := iter.Consume()
			if err != nil {
				break
			}
		}
	}

	return doc
}
*/
