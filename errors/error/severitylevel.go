package error

//go:generate stringer -type=SeverityLevel

// SeverityLevel represents the severity level of an error.
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
