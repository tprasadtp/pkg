package log

// Handler implements a log handler.
// Its up to the handler to be thread-safe.
type Handler interface {
	Init() error

	Id() string

	Level() Level
	Enabled(level Level) bool

	WithCallerInfo() bool

	Write(e Entry) error

	Close() error
	Flush() error
}
