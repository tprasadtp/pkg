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
type Event struct {
	// Event Namespace
	Namespace string

	// Time (Global)
	Time time.Time

	// Context
	// The context of the Logger that created the Record.
	// Present solely to provide Handlers access to context bound
	// fields like Trace ID and Span ID.
	// Canceling the context MUST NOT affect processing of log event.
	Context context.Context

	// Log Level (Global)
	Level Level

	// Log Message (Global)
	Message string

	// Error (Global)
	Error error

	// Contextual fields
	Fields []Field

	// Caller Info

}
