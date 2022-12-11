package log

import (
	"time"
)

// Field is logger fields. Logger contains a slice of [Fields]
// Optionally with a namespace.
type Field struct {
	Namespace string
	Key       string
	Value     any
}

// Includes caller info if available.
type Caller struct {
	// Defined represents whether caller entry is defined.
	Defined bool
	// Line number of the caller
	// If not available, this is 0.
	Line uint
	// File containing the code
	// If not available this is empty string.
	File string
	// Function name of the caller.
	// this includes full path of the package.
	// except for main package.
	// This is empty if information is not available.
	Func string
}

// Event represents a single Log event.
// Event should be considered immutable.
type Event struct {
	// Namespace is namespace of the logger that generated this event.
	Namespace string

	// Time (Global)
	Time time.Time

	// Log Level (Global)
	Level Level

	// Log Message (Global)
	Message string

	// Error (Global)
	Error error

	// Caller
	Caller Caller

	preAllocFs [fieldsBucketSize]Field
	overflowFs []Field
}
