package log

import (
	"context"
	"time"
)

// Event Represents a single Log event.
// Marshalling this to JSON/Binary must have minimal allocations.
// Log entry is immutable, it has no internal state
// or context of its own. Handlers can use already implemented Marshal functions
// or implement their own, its up to the handler to decide.
type Entry struct {
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

	pc uintptr
}

// Caller returns the file, line, function of the log event.
// If the Event was created without the necessary information,
// or if the location is unavailable, it returns ("", 0, "").
// func (e Event) Caller() (string, uint, string) {
// 	return runtime.Callers()
// }
