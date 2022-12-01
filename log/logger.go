package log

import (
	"context"
	"fmt"
	"os"
	"strings"
)

// New creates a new Logger with the given Handler.
func New(h Handler, namespaces ...string) *Logger {
	return &Logger{
		handler: h,
		exit:    os.Exit,
	}
}

// Logger.
type Logger struct {
	handler   Handler
	ctx       context.Context
	namespace string
	err       error
	fields    []Field
	exit      func(int)
}

// Clone the logger.
func (log *Logger) clone() *Logger {
	clone := *log
	return &clone
}

// Context returns Logger's context.
func (log *Logger) Context() context.Context {
	return log.ctx
}

// Namespace returns Logger's Namespace.
func (log *Logger) Namespace() string {
	return log.namespace
}

// Flush flushes Logger's Handler.
func (log *Logger) Flush() error {
	if err := log.handler.Flush(); err != nil {
		return fmt.Errorf("failed to flush handler; %w", err)
	}
	return nil
}

// WithNamespace returns a new Logger with given name segment
// appended to its original Namespace.
// Segments are joined by periods.
func (log *Logger) WithNamespace(namespace string) *Logger {
	clone := log.clone()
	if namespace != "" {
		if clone.namespace == "" {
			clone.namespace = namespace
		} else {
			clone.namespace = strings.Join([]string{log.namespace, namespace}, ".")
		}
	}
	return clone
}

// With returns a new Logger with given key-value pair,
// with optionally defined namespace. Namespace specified applies
// to the kv field, not the logger. Use WithNamespace for namespaced logger.
func (log *Logger) With(key string, value any, ns ...string) *Logger {
	clone := log.clone()
	clone.fields = append(clone.fields, F(key, value, ns...))
	return clone
}

// WithFields returns a new Logger with given Fields fields appended.
func (log *Logger) WithFields(fields Field) *Logger {
	clone := log.clone()
	clone.fields = append(clone.fields, fields)
	return clone
}

// WithContext returns a new Logger with the same handler
// as the receiver and the given attribute.
func (log *Logger) WithContext(ctx context.Context) *Logger {
	clone := log.clone()
	clone.ctx = ctx
	return clone
}

// WithExitFunc returns a new Logger with specified exit function
// This is especially useful when
//   - Libraries use logger.Fatal in their
//     code where they should not or when you do not wish a dependency
//     calling logger.Fatal to exit.
//   - You wish to specify a specific exit code so that a service manager
//     like systemd can handle it properly.
//   - You wish to perform some tasks before program exits
func (log *Logger) WithExitFunc(f func(int)) *Logger {
	clone := log.clone()
	clone.exit = f
	return clone
}

// WithError returns a new Logger with given error.
// In most cases it should be used immediately with a
// message or scoped to the context of the error.
//
//	logger.WithError(err).Error("database connection lost")
func (log *Logger) WithError(err error) *Logger {
	clone := log.clone()
	clone.err = err
	return clone
}
