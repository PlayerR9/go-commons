package errors

import (
	"fmt"
	"strings"
	"time"

	"github.com/PlayerR9/go-commons/errors/internal"
)

type ErrorCoder interface {
	~int

	String() string
}

//go:generate stringer -type=SeverityLevel

type SeverityLevel int

const (
	// INFO is the severity level for informational messages.
	// (i.e., errors that are not critical nor fatal)
	//
	// Mostly used for message passing.
	INFO SeverityLevel = iota

	// WARNING is the severity level for warning messages.
	// (i.e., errors that are not critical nor fatal, yet worthy of attention).
	WARNING

	// ERROR is the severity level for error messages.
	// (i.e., the standard error level).
	ERROR

	// FATAL is the severity level for fatal errors.
	// (i.e., the highest severity level).
	//
	// These are usually panic-level of errors.
	FATAL
)

type Err[C ErrorCoder] struct {
	Code        C
	Message     string
	Suggestions []string
	Severity    SeverityLevel
	Timestamp   time.Time
	Context     map[string]any
	StackTrace  *internal.StackTrace
	Inner       error
}

func (e Err[C]) Error() string {
	var builder strings.Builder

	builder.WriteRune('[')
	builder.WriteString(e.Severity.String())
	builder.WriteString("] Error ")
	builder.WriteString(e.Code.String())
	builder.WriteString(": ")

	if e.Message == "" {
		builder.WriteString("[no message was provided]")
	} else {
		builder.WriteString(e.Message)
	}

	if !e.Timestamp.IsZero() {
		builder.WriteString("\noccurred at: ")
		builder.WriteString(e.Timestamp.String())
	}

	if len(e.Suggestions) > 0 {
		builder.WriteString("\n\nsuggestion: ")

		for _, suggestion := range e.Suggestions {
			builder.WriteString("\n- ")
			builder.WriteString(suggestion)
		}
	}

	if len(e.Context) > 0 {
		builder.WriteString("\n\ncontext: ")

		for k, v := range e.Context {
			fmt.Fprintf(&builder, "\n- %s: %v", k, v)
		}
	}

	if e.StackTrace != nil {
		builder.WriteString("\nstack trace:\n\t")
		builder.WriteString(e.StackTrace.String())
	}

	if e.Inner != nil {
		builder.WriteString("\n\ncaused by: ")
		builder.WriteString(e.Inner.Error())
	}

	return builder.String()
}

func NewErr[C ErrorCoder](code C, message string) *Err[C] {
	return &Err[C]{
		Code:        code,
		Message:     message,
		Suggestions: nil,
		Severity:    ERROR,
		Timestamp:   time.Now(),
		Context:     nil,
		StackTrace:  nil,
	}
}

func NewErrF[C ErrorCoder](code C, format string, args ...any) *Err[C] {
	return &Err[C]{
		Code:        code,
		Message:     fmt.Sprintf(format, args...),
		Suggestions: nil,
		Severity:    ERROR,
		Timestamp:   time.Now(),
		Context:     nil,
		StackTrace:  nil,
	}
}

func (e *Err[C]) ChangeSeverity(new_severity SeverityLevel) {
	if e == nil {
		return
	}

	e.Severity = new_severity
}

func (e *Err[C]) AddSuggestion(suggestion string) {
	if e == nil {
		return
	}

	e.Suggestions = append(e.Suggestions, suggestion)
}
