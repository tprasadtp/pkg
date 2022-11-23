package log

// Handler implements a log handler.
// Its up to the handler to be thread-safe.
// Please note that handlers may not be able to preserve
// Event Levels as not all levels are supported on all handlers.
type Handler interface {
	// Init is usually used to initialize buffer pools,
	// channels do some sanity checks on the handler.
	Init() error

	Id() string

	Level() Level
	Enabled(level Level) bool

	WithCallerInfo() bool

	Write(e Entry) error

	Close() error
	Flush() error
}
