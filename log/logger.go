package log

import (
	"context"
	"strings"
)

// New creates a new Logger with the given Handler.
func New(h Handler) *Logger {
	return &Logger{
		handler: h,
	}
}

// Logger
type Logger struct {
	handler   Handler
	ctx       context.Context
	namespace string
	err       error
	fields    []Field
}

func (log *Logger) clone() *Logger {
	copy := *log
	return &copy
}

// Enabled reports whether Logger emits log records at the given level.
func (log *Logger) Enabled(level Level) bool {
	return log.Handler().Enabled(level)
}

// Context returns Logger's context.
func (log *Logger) Context() context.Context {
	return log.ctx
}

// Handler returns Logger's Handler.
func (log *Logger) Handler() Handler {
	return log.handler
}

// Namespace returns Logger's Namespace.
func (log *Logger) Namespace() string {
	return log.namespace
}

// Flush flushes Logger's Handler.
func (log *Logger) Flush() error {
	return log.handler.Flush()
}

// WithContext returns a new Logger with the same handler
// as the receiver and the given attribute.
func (log *Logger) WithContext(ctx context.Context) *Logger {
	clone := log.clone()
	clone.ctx = ctx
	return clone
}

// WithHandler returns a new Logger with specified handler
func (log *Logger) WithHandler(h Handler) *Logger {
	rv := log.clone()
	rv.handler = h
	return rv
}

// WithNamespace returns a new Logger with given name segment
// appended to its original Namespace.
// Segments are joined by periods.
func (log *Logger) WithNamespace(namespace string) *Logger {
	if namespace == "" {
		return log
	}
	rv := log.clone()
	if log.namespace == "" {
		log.namespace = namespace
	} else {
		log.namespace = strings.Join([]string{log.namespace, namespace}, ".")
	}
	return rv
}

// With returns a new Logger with given Key Value pair added to fields
func (log *Logger) With(key string, value any) *Logger {
	rv := log.clone()
	return rv
}

// WithKV returns a new Logger with given KV fields
func (log *Logger) WithKV(kv KV) *Logger {
	rv := log.clone()
	return rv
}

// WithError returns a new Logger with given error.
func (log *Logger) WithError(err error) *Logger {
	copy := log.clone()
	copy.err = err
	return copy
}
