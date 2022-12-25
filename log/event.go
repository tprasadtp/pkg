package log

import (
	"context"
	"strings"
	"time"
)

// Field is logger fields. Optionally with a namespace.
// Namespace is distinct from Logger's namespace.
// Handler MUST consider both logger's and Field's namespaces.
type Field struct {
	Namespace string
	Key       string
	Value     Value
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

// Creates a new Field. Optionally with specified namespace segments.
// Namespace segments are joined by '.', Ideally namespace segments
// should be alphanumeric and start with a letter. Namespace on fields
// are distinct from namespaces in logger. Handler MUST consider both
// namespaces.
func F(key string, value any, namespaces ...string) Field {
	switch len(namespaces) {
	case 0:
		return Field{
			Key:   key,
			Value: ToValue(value),
		}
	case 1:
		return Field{
			Namespace: namespaces[0],
			Key:       key,
			Value:     ToValue(value),
		}
	default:
		return Field{
			Namespace: strings.Join(namespaces, "."),
			Key:       key,
			Value:     ToValue(value),
		}
	}
}
