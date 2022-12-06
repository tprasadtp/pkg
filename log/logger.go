package log

import (
	"context"
	"fmt"
	"os"
	"runtime"
	"strings"
	"time"

	"github.com/tprasadtp/pkg/log/internal/helpers"
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

// Write the event directly to the handler if it is enabled.
// This should not be used by normal library users.
//   - This is intended to be used by plugins.
//   - For example stdlib plugin uses this to write
//     Event generated to handler.
func (log Logger) WriteEvent(event Event) error {
	if log.handler.Enabled(event.Level) {
		return log.handler.Write(event)
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

// With returns a new Logger with given key-value pair,
// with optionally defined namespace. Namespace specified applies
// to the kv field, not the logger. Use WithNamespace for namespaced logger.
func (log Logger) With(fields ...Field) Logger {
	fs := make([]Field, len(log.fields)+len(fields))
	copy(fs, log.fields)
	log.fields = fs
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

// write is an internal wrapper which writes event to log.Handler.
// All other named levels and methods use this with some form or other.
func (log Logger) write(level Level, message string, depth uint) {
	// // logger must not be nil.
	// if log.handler == nil {
	// 	panic(ErrLoggerInvalid)
	// }

	// Skip if handler is not enabled on the level.
	if !log.handler.Enabled(level) {
		return
	}

	// build log Event
	event := Event{
		Level:   level,
		Context: log.ctx,
		Message: message,
		Error:   log.err,
		Time:    time.Now(),
		Fields:  log.fields,
	}

	// log.handler.Write(event)

	// // If caller tracing is disabled, skip caller info and write to handler.
	if event.NoCallerTracing {
		log.handler.Write(event)
	}

	// Caller Tracing

	const maxStackLen = 50
	pc := make([]uintptr, maxStackLen)

	// Skip two extra frames to account for this function
	// and runtime.Callers itself.
	//nolint:gomnd // ignore this magic number.
	n := runtime.Callers(int(depth+2), pc)
	frames := runtime.CallersFrames(pc[:n])
	for i := 0; i < maxStackLen; i++ {
		frame, more := frames.Next()
		_, helper := helpers.Map.Load(frame.Function)
		// We ran out of frames (This implies bug in log package)
		if !more {
			event.Caller = CallerInfo{
				Line: 0,
				File: "INVALID_FRAME",
				Func: "INVALID_FRAME",
			}
			break
		}

		if !helper {
			event.Caller = CallerInfo{
				Line: uint(frame.Line),
				File: frame.File,
				Func: frame.Function,
			}
			break
		}
	}
	log.handler.Write(event)
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

// Log at WarningLevel (This is an alias for log.Warning).
func (log Logger) Warn(message string) {
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
