package log

import (
	"time"
)

// Field is logger fields.
// Value can be any (even another Field or []Field).
type Field struct {
	Key   string
	Value any
}

// Includes caller info if available.
type CallerInfo struct {
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

// Event Represents a single Log event.
// Marshalling this to JSON/Binary must have minimal allocations.
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

	// Caller
	Caller CallerInfo

	// Allocation optimization for fields
	// Typically has around 10 fields.
	Fields []Field
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
