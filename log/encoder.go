package log

import "encoding/json"

// Ensures Event type implements
// json.Marshaler interface.
var _ json.Marshaler = Event{}

const (
	// Default level key for Entry encoder.
	DefaultLevelKey = "level"

	// Default error key for Entry encoder.
	DefaultErrorKey = "error"

	// Default time key (timestamp) for Entry encoder.
	DefaultTimeKey = "time"

	// Default Trace ID key for Entry encoder.
	DefaultTraceKey = "trace_id"

	// Default Span Key for Entry encoder.
	DefaultSpanKey = "span_id"

	// Default app Version Key for Entry encoder.
	DefaultVersionKey = "version"

	// Default Caller function key for for Entry encoder.
	DefaultCallerFileKey = "file"

	// Default Caller function key for for Entry encoder.
	DefaultCallerFuncKey = "function"

	// Default Caller function key for for Entry encoder.
	DefaultCallerLineKey = "line"

	// Default Stacktrace key for for Entry encoder.
	DefaultStacktraceKey = "stacktrace"
)

// EncoderConfig is encode configuration for encoder.
type EncoderConfig struct {
	CallerFileKey string
	CallerLineKey string
	CallerFuncKey string
	StacktraceKey string

	ErrorKey   string
	SpanKey    string
	TraceKey   string
	LevelKey   string
	TimeKey    string
	VersionKey string
}

func (e Event) Encode(c EncoderConfig) ([]byte, error) {
	return nil, nil
}

// MarshalJSON implements json.Marshaler interface.
func (e Event) MarshalJSON() ([]byte, error) {
	return nil, nil
}
