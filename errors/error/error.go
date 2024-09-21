package error

import (
	"fmt"
	"strings"
	"time"

	"github.com/PlayerR9/go-commons/errors/error/internal"
)

// ErrorCoder is an interface that all error codes must implement.
type ErrorCoder interface {
	~int

	fmt.Stringer
}

// Err represents a generalized error.
type Err[C ErrorCoder] struct {
	// Code is the error code.
	Code C

	// Message is the error message.
	Message string

	// Suggestions is a list of suggestions for the user.
	Suggestions []string

	// Severity is the severity level of the error.
	Severity SeverityLevel

	// Timestamp is the timestamp of the error.
	Timestamp time.Time

	// Context is the context of the error.
	Context map[string]any

	// StackTrace is the stack trace of the error.
	StackTrace *internal.StackTrace

	// Inner is the inner error of the error.
	Inner error
}

// Error implements the error interface.
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

// NewErr creates a new error.
//
// Parameters:
//   - severity: The severity level of the error.
//   - code: The error code.
//   - message: The error message.
//
// Returns:
//   - *Err: A pointer to the new error. Never returns nil.
func NewErr[C ErrorCoder](severity SeverityLevel, code C, message string) *Err[C] {
	return &Err[C]{
		Code:        code,
		Message:     message,
		Suggestions: nil,
		Severity:    severity,
		Timestamp:   time.Now(),
		Context:     nil,
		StackTrace:  nil,
	}
}

// ChangeSeverity changes the severity level of the error. Does
// nothing if the receiver is nil.
//
// Parameters:
//   - new_severity: The new severity level of the error.
func (e *Err[C]) ChangeSeverity(new_severity SeverityLevel) {
	if e == nil {
		return
	}

	e.Severity = new_severity
}

// AddSuggestion adds a suggestion to the error. Does nothing
// if the receiver is nil.
//
// Parameters:
//   - suggestion: The suggestion to add.
func (e *Err[C]) AddSuggestion(suggestion string) {
	if e == nil {
		return
	}

	e.Suggestions = append(e.Suggestions, suggestion)
}

// AddFrame prepends a frame to the stack trace. Does nothing
// if the receiver is nil or the trace is empty.
//
// Parameters:
//   - prefix: The prefix of the frame.
//   - call: The call of the frame.
//
// If prefix is empty, the call is used as the frame. Otherwise a dot is
// added between the prefix and the call.
func (e *Err[C]) AddFrame(prefix, call string) {
	if e == nil {
		return
	}

	var frame string

	if prefix == "" {
		frame = call
	} else {
		frame = prefix + "." + call
	}

	if e.StackTrace == nil {
		e.StackTrace = internal.NewStackTrace(frame)
	} else {
		e.StackTrace.AddFrame(frame)
	}
}
