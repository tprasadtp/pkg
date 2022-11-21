package log

type Interface interface {
	WithError(err error) *Entry

	WithTraceID(id TraceID) *Entry
	WithSpanID(id TraceID) *Entry
	WithNewSpanID() *Entry
	WithNewTraceID() *Entry

	WithFields(fields ...Field) *Entry
	WithNamespace(namespace string, fields ...Field) *Entry

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

	Exit(code int, message string)
	Exitf(code int, format string, args ...any)
}
