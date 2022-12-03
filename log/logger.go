package log

import (
	"context"
	"fmt"
	"runtime"
	"strings"
	"time"
)

// New creates a new Logger with the given Handler.
// Its is recommended that namespaces start with
// a letter and only include alphanumerics in snake case or
// camel case.
func New(h Handler, namespaces ...string) *Logger {
	switch len(namespaces) {
	case 0:
		return &Logger{
			handler: h,
		}
	case 1:
		return &Logger{
			handler:   h,
			namespace: namespaces[0],
		}
	default:
		return &Logger{
			handler:   h,
			namespace: strings.Join(namespaces, "."),
		}
	}
}

// Logger.
type Logger struct {
	// Disable caller tracing.
	// This also means caller tracing is not
	// available in stacktraces.
	NoCallerTracing bool

	handler   Handler
	ctx       context.Context
	namespace string
	err       error
	fields    []Field

	exit func(int)
}

// Clone the logger.
func (log *Logger) clone() *Logger {
	clone := *log
	return &clone
}

// Enabled checks if underlying handler is enabled
// at the specified log level.
func (log *Logger) Enabled(level Level) bool {
	return log.handler.Enabled(level)
}

// Flush flushes Logger's Handler.
func (log *Logger) Flush() error {
	if err := log.handler.Flush(); err != nil {
		return fmt.Errorf("failed to flush handler; %w", err)
	}
	return nil
}

// Write the event directly to the handler if it is enabled.
// This should not be used by normal library users.
//   - This is intended by plugins.
//   - For example stdlib plugin uses this to write
//     Event generated to handler.
func (log *Logger) WriteEvent(event Event) error {
	if log.handler.Enabled(event.Level) {
		return log.handler.Write(event)
	}
	return nil
}

// Context returns Logger's context.
func (log *Logger) Context() context.Context {
	return log.ctx
}

// Namespace returns Logger's Namespace.
func (log *Logger) Namespace() string {
	return log.namespace
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

// WithCtx returns a new Logger with the same handler
// as the receiver and the given attribute.
func (log *Logger) WithCtx(ctx context.Context) *Logger {
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

// Internal wrapper which writes event to log.Handler.
// All other named levels and methods use this with some form or other.
func (log *Logger) write(level Level, message string, depth uint) error {
	if log.handler != nil {
		if log.handler.Enabled(level) {
			// build log Event
			event := Event{
				Level:   level,
				Message: message,
				Error:   log.err,
				Time:    time.Now(),
				Fields:  log.fields,
			}

			// Build caller info
			if log.NoCallerTracing {
				// depth + 1 (this function)
				if pc, file, line, ok := runtime.Caller(int(depth + 1)); ok {
					if fn := runtime.FuncForPC(pc); fn != nil {

					}
				}

			}
			return log.handler.Write(event)
		}
	}
	return nil
}
