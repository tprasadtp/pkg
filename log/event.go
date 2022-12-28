package log

import (
	"context"
	"time"
)

// Creates a new Field.
func F(key string, value any) Field {
	return Field{
		Key:   key,
		Value: AnyValue(value),
	}
}

// Field is logger fields.
type Field struct {
	Key   string
	Value Value
}

// Includes caller info if available.
type CallerInfo struct {
	// Defined represents whether caller info is defined.
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

	// Context Cancelling this context has no effect on the logging operations.
	Ctx context.Context

	// Caller
	Caller CallerInfo

	// Fields
	Fields []Field
}
