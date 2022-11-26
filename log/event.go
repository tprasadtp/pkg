package log

import (
	"context"
	"time"
)

// Field is Key value pair
type Field struct {
	Namespace string
	Key       string
	Value     any
}

// Event Represents a single Log event.
// Marshalling this to JSON/Binary must have minimal allocations.
// Log entry is immutable, it has no internal state
// or context of its own. Handlers can use already implemented Marshal functions
// or implement their own, its up to the handler to decide.
type Event struct {
	// Namespace
	Namespace string

	// Time (Global)
	Time time.Time

	// Context
	// The context of the Logger that created the Record. Present
	// solely to provide Handlers access to the context's values.
	// Canceling the context MUST NOT affect processing of log event.
	Context context.Context

	// Log Level (Global)
	Level Level

	// Log Message (Global)
	Message string

	// Error (Global)
	Error error

	// OpenTracing/OpenTelemetry Trace ID (Global)
	TraceID string

	// OpenTracing/OpenTelemetry Span ID (Global)
	SpanID string

	// Contextual fields
	Fields []Field

	// CallerInfo is only populated if one of the handlers in
	// root logger has WithCallerInfo() returns true (Global)

	// Version
	// This is extremely useful for A/B deployed logs
	// Unlike Time, Level and similar fields, this is NOT Global
	// and inherits namespace of the Entry.
	Version string
}
