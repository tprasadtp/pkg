package log

// A Handler describes the logging backend.
// It handles log records produced by a Logger.
// A typical handler may
//   - print log records to standard error,
//   - write them to a file, database or network service etc.
type Handler interface {
	// Init is usually used to initialize buffer pools,
	// channels do some sanity checks on the handler.
	Init() error

	Id() string

	Level() Level
	Enabled(level Level) bool

	IncludeCallerInfo() bool

	Write(e Entry) error

	Close() error
	Flush() error
}
