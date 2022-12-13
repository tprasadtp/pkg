package log

import (
	"os"
	"strings"
)

// fieldsBucketSize represents step size when growing capacity of
// fields slice. Capacity of Fields slice is always a integer multiple
// of this number.
const fieldsBucketSize = 25

// New creates a new Logger with the given Handler.
// This pre allocates some storage for storing fields.
func New(handler Handler) Logger {
	return Logger{
		handler: handler,
		fields:  make([]Field, 0, fieldsBucketSize),
	}
}

// Logger.
type Logger struct {
	handler   Handler
	namespace string
	err       error
	fields    []Field
	exit      func()
}

// Handler exposes the underlying handler.
func (log Logger) Handler(level Level) Handler {
	return log.handler
}

// Namespace returns Logger's Namespace.
func (log Logger) Namespace() string {
	return log.namespace
}

// Err returns Logger's error context.
func (log Logger) Err() []Field {
	return log.fields
}

// Fields returns Logger's fields context.
func (log Logger) Fields() []Field {
	return log.fields
}

// WithNamespace returns a new Logger with given name segment
// appended to its original Namespace. Segments are joined by periods.
// This is useful if you want to pass the logger to a library, especially
// the one which you don't control. This will always return a new logger
// even when specified namespace is empty.Its is recommended that namespaces
// start with a letter and only include alphanumerics in snake case or
// camel case.
func (log Logger) WithNamespace(namespace string) Logger {
	if namespace != "" {
		if log.namespace == "" {
			log.namespace = namespace
		} else {
			log.namespace = strings.Join([]string{log.namespace, namespace}, ".")
		}
	}
	return log
}

// WithExitFunc returns a new Logger with specified exit function.
// By default logger uses [os.Exit](1) if exit function is not specified or nil.
func (log Logger) WithExitFunc(fn func()) Logger {
	log.exit = fn
	return log
}

// WithErr returns a new Logger with given error.
// In most cases it should be used immediately with a
// message or scoped to the context of the error.
// If logger already contains an error, it is replaced.
// It is job of the package author to wrap error not the logger.
//
//	logger.WithError(err).Error("package metadata database connection lost")
func (log Logger) WithErr(err error) Logger {
	log.err = err
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
	log.write(level, message, 1)
}

// Log at TraceLevel.
func (log Logger) Trace(message string) {
	log.write(LevelTrace, message, 1)
}

// Log at DebugLevel.
func (log Logger) Debug(message string) {
	log.write(LevelDebug, message, 1)
}

// Log at VerboseLevel.
func (log Logger) Verbose(message string) {
	log.write(LevelVerbose, message, 1)
}

// Log at InfoLevel.
func (log Logger) Info(message string) {
	log.write(LevelInfo, message, 1)
}

// Log at SuccessLevel.
func (log Logger) Success(message string) {
	log.write(LevelSuccess, message, 1)
}

// Log at NoticeLevel.
func (log Logger) Notice(message string) {
	log.write(LevelNotice, message, 1)
}

// Log at WarningLevel.
func (log Logger) Warning(message string) {
	log.write(LevelWarning, message, 1)
}

// Log at ErrorLevel.
func (log Logger) Error(message string) {
	log.write(LevelError, message, 1)
}

// Log at CriticalLevel AND flush the handler.
func (log Logger) Critical(message string) {
	log.write(LevelCritical, message, 1)
	log.handler.Flush()
}

// Log at FatalLevel flush and close the handler.
func (log Logger) Fatal(message string) {
	log.write(LevelFatal, message, 1)
	log.handler.Close()
	if log.exit == nil {
		os.Exit(1)
	} else {
		log.exit()
	}
}
