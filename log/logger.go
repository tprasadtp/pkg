package log

import "context"

// Logger
type Logger struct {
	handler   Handler
	ctx       context.Context
	namespace string
	err       error
	fields    []Field
	span      string
	trace     string
}

// NewLogger returns a new Logger with specified handler
func NewLogger(h Handler) *Logger {
	return &Logger{
		handler: h,
	}
}
