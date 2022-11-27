package log

import (
	"context"
	"fmt"
)

// Logger
type Logger struct {
	handler   Handler
	ctx       context.Context
	namespace string
	err       error
	fields    []Field
}

// New creates a new Logger with the given Handler.
func New(h Handler) *Logger {
	return &Logger{
		handler: h,
	}
}

// Enabled reports whether Logger emits log records at the given level.
func (l *Logger) Enabled(level Level) bool {
	return l.Handler().Enabled(level)
}

// Context returns Logger's context.
func (l *Logger) Context() context.Context {
	return l.ctx
}

// Handler returns Logger's Handler.
func (l *Logger) Handler() Handler {
	return l.handler
}

// Namespace returns Logger's Namespace.
func (l *Logger) Namespace() string {
	return l.namespace
}

// Flush flushes Logger's Handler.
func (l *Logger) Flush() error {
	return l.handler.Flush()
}

// Close flushes Logger's Handler and closes it.
// This also flushes the handler.
func (l *Logger) Close() error {
	err := l.handler.Flush()
	if err != nil {
		return fmt.Errorf("%w;%e", err, l.handler.Close())
	} else {
		return l.handler.Close()
	}
}

// WithContext returns a new Logger with the same handler
// as the receiver and the given context.
func (l *Logger) WithContext(ctx context.Context) *Logger {
	rv := *l
	rv.ctx = ctx
	return &rv
}
