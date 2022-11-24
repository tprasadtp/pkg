package log

import (
	"fmt"
	"strings"
	"sync"

	"go.uber.org/multierr"
)

// Logger
type Logger struct {
	mu       sync.Mutex
	handlers []Handler
}

// SubLogger inherits handlers from Logger,
// It implements all methods of Logger
// with exception of managing handlers.
type SubLogger struct {
	handlers  []Handler
	namespace string
	err       error
	fields    []Field
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

	var err error
	for _, h := range l.handlers {
		if h.Id() == handler.Id() {
			multierr.Append(err, fmt.Errorf("log: handler with name %s already exits", handler.Id()))
		}
	}

	if err == nil {
		multierr.Append(err, handler.Init())
		if err == nil {
			l.handlers = append(l.handlers, handler)
		}
	}
	return err
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

// Flushes all handlers synchronously.
func (l *Logger) Flush() error {
	var err error
	for _, h := range l.handlers {
		multierr.Append(err, h.Flush())
	}
	return err
}

// WithNamespace returns a SubLogger with added namespace
func (l *Logger) WithNamespace(name string) SubLogger {
	l.mu.Lock()
	defer l.mu.Unlock()

	// If no namespace is provides, we create a random namespace
	if strings.TrimSpace(strings.ToLower(name)) == "" {
		name = "ns_invalid"
	}

	return SubLogger{
		handlers:  l.handlers,
		namespace: name,
	}
}

// WithNamespace returns a SubLogger with added fields
func (l *Logger) WithFields(fields ...Field) SubLogger {
	l.mu.Lock()
	defer l.mu.Unlock()

	return SubLogger{
		handlers: l.handlers,
		fields:   fields,
	}
}

// WithField returns a SubLogger with added field
func (l *Logger) WithField(key string, value any) SubLogger {
	l.mu.Lock()
	defer l.mu.Unlock()

	return SubLogger{
		handlers: l.handlers,
		fields: []Field{
			{
				Key:   key,
				Value: value,
			},
		},
	}
}

// WithError returns a SubLogger with error field
func (l *Logger) WithError(err error) SubLogger {
	l.mu.Lock()
	defer l.mu.Unlock()

	return SubLogger{
		handlers: l.handlers,
		err:      err,
	}
}
