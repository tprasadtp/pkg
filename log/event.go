package log

import (
	"time"
)

// Field is logger fields. Logger contains a slice of [Fields]
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

// Event represents a single Log event. Event should be considered immutable.
// Unlike many logging libraries this does not use [sync.Pool]
// as handler is an interface, and it would have to depend on implementation
// to release the Event back to the pool. Instead, it uses a pre-allocated
// fixed size array which should be sufficient for most cases.
// However, in cases where a log event has more than 20 fields,
// this will allocate.
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
	// Typically has less than 20 or  so fields.
	fieldsPrealloc     [fieldsBucketSize]Field
	fileldsPreallocLen uint

	fieldsOverflow []Field
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
