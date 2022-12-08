//nolint:unparam,errcheck // These two are pretty useless in this file.
package log

import (
	"context"
	"fmt"
	"os"
	"strings"
)

// New creates a new Logger with the given Handler.
// Its is recommended that namespaces start with
// a letter and only include alphanumerics in snake case or
// camel case.
func New(handler Handler, namespaces ...string) Logger {
	switch len(namespaces) {
	case 0:
		return Logger{
			handler: handler,
		}
	case 1:
		return Logger{
			handler:   handler,
			namespace: namespaces[0],
		}
	default:
		return Logger{
			handler:   handler,
			namespace: strings.Join(namespaces, "."),
		}
	}
}

// Logger.
type Logger struct {
	// Non exported fields.
	handler   Handler
	ctx       context.Context
	namespace string
	err       error
	fields    []Field
	exit      func()
}

// Enabled checks if underlying handler is enabled
// at the specified log level.
func (log Logger) Enabled(level Level) bool {
	return log.handler.Enabled(level)
}

// Flush flushes Logger's Handler.
func (log Logger) Flush() error {
	if err := log.handler.Flush(); err != nil {
		return fmt.Errorf("log: failed to flush handler; %w", err)
	}
	return nil
}

// Namespace returns Logger's Namespace.
func (log Logger) Namespace() string {
	return log.namespace
}

// Context returns Logger's context.
// Use [github.com/tprasadtp/pkg/log/Logger.WithCtx]
// for passing the context to the logger.
func (log Logger) Context() context.Context {
	return log.ctx
}

// WithCtx returns a new Logger with the same handler
// as the receiver and the given attribute.
func (log Logger) WithCtx(ctx context.Context) Logger {
	log.ctx = ctx
	return log
}

// WithNamespace returns a new Logger with given name segment
// appended to its original Namespace. Segments are joined by periods.
// This is useful if you want to pass the logger to a library, especially
// the one which you don't control. This will always return a new logger
// even when specified namespace is empty.
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

// WithExitFunc returns a new Logger with specified exit function
// This is especially useful when
//   - Libraries use logger.Fatal in their
//     code where they should not or when you do not wish a dependency
//     calling logger.Fatal to exit.
//   - You wish to specify a specific exit code so that a service manager
//     like systemd can handle it properly.
//   - You wish to perform some tasks before program exits
func (log Logger) WithExitFunc(fn func()) Logger {
	log.exit = fn
	return log
}

// WithError returns a new Logger with given error.
// In most cases it should be used immediately with a
// message or scoped to the context of the error.
//
//	logger.WithError(err).Error("database connection lost")
func (log Logger) WithError(err error) Logger {
	log.err = err
	return log
}

// With returns a new Logger with given key-value pair,
// with optionally defined namespace. Namespace specified applies
// to the kv field, not the logger. Use WithNamespace for namespaced logger.
func (log Logger) With(fields ...Field) Logger {
	fs := make([]Field, len(log.fields)+len(fields))
	copy(fs, log.fields)
	log.fields = fs
	return log
}

// Write Log message with custom level, Usually you do not need this
// unless you are using custom logging levels. Use one of the named log
// levels instead.
func (log Logger) Log(level Level, message string) {
	log.write(level, message, 1)
}

// Log at TraceLevel.
func (log Logger) Trace(message string) {
	log.write(TraceLevel, message, 1)
}

// Log at DebugLevel.
func (log Logger) Debug(message string) {
	log.write(DebugLevel, message, 1)
}

// Log at VerboseLevel.
func (log Logger) Verbose(message string) {
	log.write(VerboseLevel, message, 1)
}

// Log at InfoLevel.
func (log Logger) Info(message string) {
	log.write(InfoLevel, message, 1)
}

// Log at SuccessLevel.
func (log Logger) Success(message string) {
	log.write(SuccessLevel, message, 1)
}

// Log at NoticeLevel.
func (log Logger) Notice(message string) {
	log.write(NoticeLevel, message, 1)
}

// Log at WarningLevel.
func (log Logger) Warning(message string) {
	log.write(WarningLevel, message, 1)
}

// Log at ErrorLevel.
func (log Logger) Error(message string) {
	log.write(ErrorLevel, message, 1)
}

// Log at CriticalLevel AND flush the handler.
func (log Logger) Critical(message string) {
	log.write(CriticalLevel, message, 1)
	log.handler.Flush()
}

// Log at FatalLevel AND flush the handler.
func (log Logger) Fatal(message string) {
	log.write(FatalLevel, message, 1)
	log.handler.Flush()
	if log.exit == nil {
		os.Exit(1)
	} else {
		log.exit()
	}
}

// Write Log message with custom level, Usually you do not need this,
// unless you need custom logging levels.
// Prefer using one of the named log levels instead.
func (log Logger) Logf(level Level, format string, args ...any) {
	log.write(level, fmt.Sprintf(format, args...), 1)
}

// Log at TraceLevel.
func (log Logger) Tracef(format string, args ...any) {
	log.write(TraceLevel, fmt.Sprintf(format, args...), 1)
}

// Log at DebugLevel.
func (log Logger) Debugf(format string, args ...any) {
	log.write(DebugLevel, fmt.Sprintf(format, args...), 1)
}

// Log at VerboseLevel.
func (log Logger) Verbosef(format string, args ...any) {
	log.write(VerboseLevel, fmt.Sprintf(format, args...), 1)
}

// Log at InfoLevel.
func (log Logger) Infof(format string, args ...any) {
	log.write(InfoLevel, fmt.Sprintf(format, args...), 1)
}

// Log at SuccessLevel.
func (log Logger) Successf(format string, args ...any) {
	log.write(SuccessLevel, fmt.Sprintf(format, args...), 1)
}

// Log at NoticeLevel.
func (log Logger) Noticef(format string, args ...any) {
	log.write(NoticeLevel, fmt.Sprintf(format, args...), 1)
}

// Log at WarningLevel.
func (log Logger) Warningf(format string, args ...any) {
	log.write(WarningLevel, fmt.Sprintf(format, args...), 1)
}

// Log at ErrorLevel.
func (log Logger) Errorf(format string, args ...any) {
	log.write(ErrorLevel, fmt.Sprintf(format, args...), 1)
}

// Log at CriticalLevel AND flush the handler.
func (log Logger) Criticalf(format string, args ...any) {
	log.write(CriticalLevel, fmt.Sprintf(format, args...), 1)
	log.handler.Flush()
}

// Log at FatalLevel AND flush the handler.
func (log Logger) Fatalf(format string, args ...any) {
	log.write(FatalLevel, fmt.Sprintf(format, args...), 1)
	log.handler.Flush()
	if log.exit == nil {
		os.Exit(1)
	} else {
		log.exit()
	}
}
