package log

import (
	"context"
	"os"
)

// fieldsBucketSize represents step size when growing capacity of
// fields slice. Capacity of Fields slice is always a integer multiple
// of this number.
const fieldsBucketSize = 25

// NewLogger creates a new Logger with the given Handler.
// This pre allocates some storage for storing fields.
func NewLogger(handler Handler) Logger {
	return Logger{
		handler: handler,
		fields:  make([]Field, 0, fieldsBucketSize),
	}
}

// Logger.
type Logger struct {
	handler   Handler
	namespace string
	ctx       context.Context
	err       error
	fields    []Field
	caller    bool
}

// Handler exposes logger's underlying handler.
func (log Logger) Handler() Handler {
	return log.handler
}

// Namespace returns Logger's Namespace.
func (log Logger) Namespace() string {
	return log.namespace
}

// Err returns Logger's error context.
func (log Logger) Err() error {
	return log.err
}

// Fields returns Logger's fields context.
func (log Logger) Fields() []Field {
	return log.fields
}

// Caller returns whether of not caller tracing is enabled.
func (log Logger) Caller() bool {
	return log.caller
}

// Fields returns Logger's embedded context.
func (log Logger) Ctx() context.Context {
	return log.ctx
}

// WithCaller enables tracing caller info like,
// file name and line number. This costs 2 allocations (when logger is enabled)
// due to [runtime.Frames] implementation, which this depends on.
func (log Logger) WithCaller() Logger {
	log.caller = true
	return log
}

// WithoutCaller disables tracing caller info. (default).
// This is typically used to disable caller tracing,
// which can be bit expensive and costs an allocation.
func (log Logger) WithoutCaller() Logger {
	log.caller = false
	return log
}

// WithNamespace returns a new Logger with given name segment
// appended to its original Namespace. Segments are joined by periods.
// This is useful if you want to pass the logger to a library, especially
// the one which you don't control. This will always return a new logger
// even when specified namespace is empty. Its is recommended that namespaces
// start with a letter and only include alphanumerics in snake case or
// camel case.
func (log Logger) WithNamespace(namespace string) Logger {
	if namespace != "" {
		if log.namespace == "" {
			log.namespace = namespace
		} else {
			log.namespace = log.namespace + "." + namespace
		}
	}
	return log
}

// WithErr returns a new Logger with given error.
// In most cases it should be used immediately with a
// message or scoped to the context of the error.
// If logger already contains an error, it is replaced.
//
//	logger.WithErr(err).Error("Failed to update metadata cache file.")
func (log Logger) WithErr(err error) Logger {
	log.err = err
	return log
}

// WithCtx returns a new logger with specified context embedded.
// This is useful for handlers which process the context directly
// and enrich log entries. Cancelling this context has no effect
// whatsoever on the log entry. Context might include
// request scoped information and tracing data.
// If logger already contains a context,
// it is replaced as context is considered immutable.
func (log Logger) WithCtx(ctx context.Context) Logger {
	log.ctx = ctx
	return log
}

// With returns a new Logger with given fields.
// This will allocate if logger's underlying fields slice capacity is
// smaller than required.
func (log Logger) With(fields ...Field) Logger {
	m := len(log.fields)
	n := m + len(fields)

	// Check if fields slice can store all the fields.
	// if not re-allocate in fieldsBucketSize increments.
	if n > cap(log.fields) {
		buckets := (n / fieldsBucketSize) + 1
		newSlice := make([]Field, m, fieldsBucketSize*buckets)
		// If log.fields has elements, copy them to new slice.
		if m > 0 {
			copy(newSlice[:m], log.fields)
		}
		log.fields = newSlice
	}
	// log.fields's backing array has enough capacity,
	// so append wont allocate.
	log.fields = append(log.fields, fields...)
	return log
}

// Write a log message with custom level.
// Prefer using one of the named log levels instead.
func (log Logger) Log(level Level, message string) {
	log.write(level, message)
}

// Log at TraceLevel.
func (log Logger) Trace(message string) {
	log.write(LevelTrace, message)
}

// Log at DebugLevel.
func (log Logger) Debug(message string) {
	log.write(LevelDebug, message)
}

// Log at VerboseLevel.
func (log Logger) Verbose(message string) {
	log.write(LevelVerbose, message)
}

// Log at InfoLevel.
func (log Logger) Info(message string) {
	log.write(LevelInfo, message)
}

// Log at SuccessLevel.
func (log Logger) Success(message string) {
	log.write(LevelSuccess, message)
}

// Log at NoticeLevel.
func (log Logger) Notice(message string) {
	log.write(LevelNotice, message)
}

// Log at WarningLevel.
func (log Logger) Warning(message string) {
	log.write(LevelWarning, message)
}

// Log at ErrorLevel.
func (log Logger) Error(message string) {
	log.write(LevelError, message)
}

// Log at CriticalLevel AND flush the handler.
func (log Logger) Critical(message string) {
	log.write(LevelCritical, message)
	log.handler.Flush()
}

// Log at FatalLevel and exit.
func (log Logger) Fatal(message string) {
	log.write(LevelFatal, message)
	log.handler.Flush()
	log.handler.Close()
	os.Exit(1)
}
