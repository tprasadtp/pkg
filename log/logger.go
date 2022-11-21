package log

import (
	"fmt"
	"sync"

	"go.uber.org/multierr"
)

// var (
// 	_ Interface = Logger{}
// )

// Returns a new logger
func New() *Logger {
	return &Logger{}
}

// Logger
type Logger struct {
	handlers []Handler
	mu       sync.Mutex

	names  []string
	fields []Field
	err    error

	callerInfo bool

	traceID string
	spanID  string
}

// Remove all existing handlers and set a new handler.
// This will flush and close all existing handlers
// before setting the new handler as the only handler.
func (l *Logger) SetHandler(handler Handler) error {
	l.mu.Lock()
	defer l.mu.Unlock()
	var err error
	for _, h := range l.handlers {
		multierr.Append(err, h.Close())
	}
	if err == nil {
		l.handlers = []Handler{handler}
	}
	return err
}

// Add a handler to logger. Handler must be unique.
func (l *Logger) AddHandler(handler Handler) error {
	l.mu.Lock()
	defer l.mu.Unlock()

	for _, h := range l.handlers {
		if h.Id() == handler.Id() {
			return fmt.Errorf("log: handler with name %s already exits", handler.Id())
		}
	}
	l.handlers = append(l.handlers, handler)
	return nil
}

// Remove existing handler with id specified.
// It is an error to remove a non existent handler.
func (l *Logger) RemoveHandler(id string) error {
	l.mu.Lock()
	defer l.mu.Unlock()

	var err error
	var oldHandlersCount = len(l.handlers)

	for i, h := range l.handlers {
		if h.Id() == id {
			multierr.Append(err, h.Close())
		}
		if err == nil {
			// Preserves order of handlers
			l.handlers = append(l.handlers[:i], l.handlers[i+1:]...)
		}
	}
	if err == nil {
		if oldHandlersCount == len(l.handlers) {
			return fmt.Errorf("log: handler with id %s is not present", id)
		}
	}
	return err
}

// Returns a namespaced logger. By default this is empty
// This is propagated to Fields. Useful to isolate components
func (l *Logger) WithName(name string) *Entry {
	return &Entry{
		Logger: l,
	}
}

// 	WithError(err error) *Entry
// 	WithTraceID(id TraceID) *Entry
// 	WithSpanID(id SpanID) *Entry
// 	WithFields(fields ...Field) *Entry
// 	WithNamespacedFields(name string, fields ...Field) *Entry
// 	WithName(name string) *Entry

// 	Log(level Level, message string)
// 	Logf(level Level, format string, args ...any)

// 	Debug(message string)
// 	Debugf(format string, args ...any)

// 	Verbose(message string)
// 	Verbosef(format string, args ...any)

// 	Info(message string)
// 	Infof(format string, args ...any)

// 	Success(message string)
// 	Successf(format string, args ...any)

// 	Warn(message string)
// 	Warnf(format string, args ...any)

// 	Error(message string)
// 	Errorf(format string, args ...any)

// 	Panic(message string)
// 	Panicf(format string, args ...any)

// 	Exit(code int, message string)
// 	Exitf(code int, format string, args ...any)

// 	Flush(timeout time.Duration) error
