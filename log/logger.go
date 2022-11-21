package log

import (
	"time"
)

type Event struct {
	Level     Level
	Timestamp time.Time
	Message   string
	Fields    []Field
}

func (e *Event) Write() error {
	return nil
}

type Handler interface {
	Id() string
	Level() Level
	Enabled(Level) bool

	WriteEvent(e *Event) error

	Close() error
	Flush() error
}

type Logger interface {
	SetHandler(handler *Handler) error

	AddHandler(handler *Handler) error
	RemoveHandler(id string) error

	WithError(err error) Logger

	WithFields(fields []Field) Logger
	WithField(field Field) Logger

	Log(level Level, message string)
	Logf(level Level, format string, args ...any)

	Debug(message string)
	Debugf(format string, args ...any)

	Verbose(message string)
	Verbosef(format string, args ...any)

	Info(message string)
	Infof(format string, args ...any)

	Success(message string)
	Successf(format string, args ...any)

	Warn(message string)
	Warnf(format string, args ...any)

	Error(message string)
	Errorf(format string, args ...any)

	Panic(message string)
	Panicf(format string, args ...any)

	Exit(code int, message string)
	Exitf(code int, format string, args ...any)
}
