package log

import (
	"context"
	"strings"
	"time"
)

// Field is logger fields.
type Field struct {
	Namespace string
	Key       string
	Value     Value
}

// Stacktrace includes stacktrace of the error.
type StackTrace struct {
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

	// Contextual fields
	Fields []Field

	// Callerinfo this can be nil, of caller info collection
	// is disabled.
	Caller CallerInfo

	// StackTrace is stacktrace. This is nil
	// if stacktracing is disabled.
	StackTrace StackTrace
}

// Returns a new Field. Namespace is optional,
// if multiple namespaces are specified, it is joined with a '.'
// as the separator.
func F(key string, value any, namespace ...string) Field {
	switch len(namespace) {
	case 0:
		return Field{
			Key:   key,
			Value: NewValue(value),
		}
	case 1:
		return Field{
			Namespace: namespace[0],
			Key:       key,
			Value:     NewValue(value),
		}
	default:
		return Field{
			Namespace: strings.Join(namespace, "."),
			Key:       key,
			Value:     NewValue(value),
		}
	}
}
