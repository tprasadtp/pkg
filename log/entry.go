package log

import (
	"time"
)

// Caller info
type CallerInfo struct {
	File string
	Line uint
}

// Field is Key value pair
type Field struct {
	Namespace string
	Key       string
	Value     any
}

// Entry Represents a single Log entry.
// Marshalling this to JSON/Binary must have minimal allocations.
// Log entry is immutable, it has no internal state
// or context of its own Handlers can use already implemented Marshal functions
// or implement their own, its up to the handler to decide.
type Entry struct {
	// Logger Namespace
	// This can be different than filed namespaces.
	// Filed namespaces will always append logger namespace
	// if their Namespace filed is empty.
	Namespace string

	// Time (Global)
	Time time.Time

	// Log Level (Global)
	Level Level

	// Log Message (Global)
	Msg string

	// Error (Global)
	Error error

	// OpenTracing/OpenTelemetry Trace ID (Global)
	Trace string

	// OpenTracing/OpenTelemetry Span ID (Global)
	Span string

	// Contextual fields
	// Marshaller MUST consider namespaces on both Entry AND Fields.
	Fields []Field

	// CallerInfo is only populated if one of the handlers in
	// root logger has WithCallerInfo() returns true (Global)
	Caller CallerInfo
}

// JSONEncoderConfig is encode configuration for JSON encoder
type JSONEncoderConfig struct {
	CallerFileKey string
	CallerLineKey string
	ErrorKey      string
	LevelKey      string
	SpanKey       string
	TimeKey       string
	TraceKey      string
}

var DefaultJSONEncoderConfig = JSONEncoderConfig{
	CallerFileKey: "file",
	CallerLineKey: "line",
	ErrorKey:      "error",
	LevelKey:      "level",
	SpanKey:       "span_id",
	TimeKey:       "time",
	TraceKey:      "trace_id",
}

// MarshalJSON implements json.Marshaler interface
// This used default JsonEncoderConfig.
func (e Entry) MarshalJSON() ([]byte, error) {
	return nil, nil
}

// MarshalJSON implements json.Marshaler interface
// This is same as MarshalJSON but used custom JSONEncoderConfig.
// Useful when custom keys are required
func (e Entry) MarshalJSONWithConfig(c JSONEncoderConfig) ([]byte, error) {
	return nil, nil
}

// MarshalLogFmt marshalls Event to Logfmt format
func (e Entry) MarshalLogFmt() ([]byte, error) {
	return nil, nil
}

// MarshalLogFmt marshalls Event to Journald format
func (e Entry) MarshalJournald() ([]byte, error) {
	return nil, nil
}

// MarshalLogFmt marshalls Event to syslog format
func (e Entry) MarshalSyslog() ([]byte, error) {
	return nil, nil
}
