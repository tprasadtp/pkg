package log

import (
	"fmt"
	"strings"
	"sync"

	"go.uber.org/multierr"
)

// core holds core stateful elements of Logger
// like handlers, mutexes and other data. Once initialized,
// it MUST ALWAYS be passed as by reference(pointer)
// and NEVER copied.
type core struct {
	handlers []Handler
	mu       sync.Mutex
}

// Logger Logger includes handlers
type Logger struct {
	namespace string
	core      *core
	fields    []Field
	err       error
}

func (l *Logger) clone() Logger {
	l.core.mu.Lock()
	defer l.core.mu.Unlock()

	return Logger{
		core:      l.core,
		fields:    l.fields,
		err:       l.err,
		namespace: l.namespace,
	}
}

// Remove all existing handlers and set a new handler.
// This will flush and close all existing handlers
// before setting the new handler as the only handler.
func (l *Logger) SetHandler(handler Handler) error {
	l.core.mu.Lock()
	defer l.core.mu.Unlock()
	var err error
	for _, h := range l.core.handlers {
		multierr.Append(err, h.Close())
	}
	if err == nil {
		l.core.handlers = []Handler{handler}
	}
	return err
}

// Add a handler to logger. Handler must be unique.
func (l *Logger) AddHandler(handler Handler) error {
	l.core.mu.Lock()
	defer l.core.mu.Unlock()

	var err error
	for _, h := range l.core.handlers {
		if h.Id() == handler.Id() {
			multierr.Append(err, fmt.Errorf("log: handler with name %s already exits", handler.Id()))
		}
	}

	l.core.handlers = append(l.core.handlers, handler)
	return nil
}

// Remove existing handler with id specified.
// It is an error to remove a non existent handler.
func (l *Logger) RemoveHandler(id string) error {
	l.core.mu.Lock()
	defer l.core.mu.Unlock()

	var err error
	var oldHandlersCount = len(l.core.handlers)

	for i, h := range l.core.handlers {
		if h.Id() == id {
			multierr.Append(err, h.Close())
		}
		if err == nil {
			// Preserves order of handlers
			l.core.handlers = append(l.core.handlers[:i], l.core.handlers[i+1:]...)
		}
	}
	if err == nil {
		if oldHandlersCount == len(l.core.handlers) {
			return fmt.Errorf("log: handler with id %s is not present", id)
		}
	}
	return err
}

// WithNamespace returns copy of the logger with added namespace
func (l *Logger) WithNamespace(name string) Logger {
	l.core.mu.Lock()
	defer l.core.mu.Unlock()
	rv := l.clone()

	// If no namespace is provides, we create a random namespace
	if strings.TrimSpace(strings.ToLower(name)) == "" {
		name = "ns_invalid"
	}

	if len(rv.namespace) > 0 {
		rv.namespace = rv.namespace + "." + name
	} else {
		rv.namespace = name
	}
	return rv
}
