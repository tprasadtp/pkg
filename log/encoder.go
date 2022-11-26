package log

// EntryEncoderConfig is encode configuration for encoder
type EntryEncoderConfig struct {
	CallerFileKey string
	CallerLineKey string
	CallerFuncKey string

	ErrorKey   string
	SpanKey    string
	TraceKey   string
	LevelKey   string
	TimeKey    string
	VersionKey string
}

const (
	// Default level key for Entry encoder
	DefaultLevelKey = "level"

	// Default error key for Entry encoder
	DefaultErrorKey = "error"

	// Default time key (timestamp) for Entry encoder
	DefaultTimeKey = "time"

	// Default Trace ID key for Entry encoder
	DefaultTraceKey = "trace"

	// Default Span Key for Entry encoder
	DefaultSpanKey = "spanID"

	// Default app Version Key for Entry encoder
	DefaultVersionKey = "version"

	// Default Caller function key for for Entry encoder
	DefaultCallerFileKey = "file"

	// Default Caller function key for for Entry encoder
	DefaultCallerFuncKey = "function"

	// Default Caller function key for for Entry encoder
	DefaultCallerLineKey = "line"
)

// DefaultEntryEncoderConf is default Entry encoder configuration.
// It is not recommended to change this.
// Use Marshal*WithConfig functions instead of modifying the default.
var DefaultEntryEncoderConf = EntryEncoderConfig{
	CallerFileKey: DefaultCallerFileKey,
	CallerLineKey: DefaultCallerLineKey,
	CallerFuncKey: DefaultCallerFuncKey,
	ErrorKey:      DefaultErrorKey,
	SpanKey:       DefaultSpanKey,
	TraceKey:      DefaultTraceKey,
	LevelKey:      DefaultLevelKey,
	TimeKey:       DefaultLevelKey,
	VersionKey:    DefaultVersionKey,
}

// MarshalJSON implements json.Marshaler interface
// This uses DefaultEntryEncoderConf.
func (e Event) MarshalJSON() ([]byte, error) {
	return nil, nil
}

// MarshalJSON implements json.Marshaler interface
// This is same as MarshalJSON but uses custom EncoderConfig.
// Useful when custom keys are required.
func (e Event) MarshalJSONWithConfig(c EntryEncoderConfig) ([]byte, error) {
	return nil, nil
}

// MarshalLogFmt marshalls Event to Journald format
func (e Event) MarshalJournald() ([]byte, error) {
	return nil, nil
}

// MarshalLogFmt marshalls Event to syslog format
func (e Event) MarshalSyslog() ([]byte, error) {
	return nil, nil
}
