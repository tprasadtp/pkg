package log

import (
	"context"
	"time"
)

const fieldsFront = 10

// Field is logger fields.
// Value can be any (even another Field or []Field).
type Field struct {
	Key   string
	Value any
}

// Includes caller info if available.
type CallerInfo struct {
	// Line number of the caller
	// If not available, this is 0.
	Line uint
	// File containing the code
	// If not available this is empty string.
	File string
	// Function name of the caller.
	// this includes full path of the package.
	// except for main package. This is limitation
	// of [runtime.FuncForPC]
	// This is empty if information is not available.
	Func string
}

// Event Represents a single Log event.
// Marshalling this to JSON/Binary must have minimal allocations.
// Log entry is immutable, it has no internal state
// or context of its own. Handlers can use already implemented Marshal functions
// or implement their own, its up to the handler to decide.
type Event struct {
	// Event Namespace
	Namespace string

	// Time (Global)
	Time time.Time

	// Log Level (Global)
	Level Level

	// Log Message (Global)
	Message string

	// Error (Global)
	Error error

	// Context
	// The context of the Logger that created the Record.
	// Present solely to provide Handlers access to context bound
	// fields like Trace ID and Span ID.
	// Canceling the context MUST NOT affect processing of log event.
	Context context.Context

	// Caller
	Caller CallerInfo

	// Allocation optimization for fields
	// Typically has around 10 fields.
	Fields [50]Field
}

// Returns a new Field.
func F(key string, value any) Field {
	return Field{
		Key:   key,
		Value: value,
	}
}

// Returns a new Map field.
func M(key string, fields ...Field) Field {
	return Field{
		Key:   key,
		Value: fields,
	}
}
