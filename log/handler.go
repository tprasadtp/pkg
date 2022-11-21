package log

type Handler interface {
	Id() string

	Level() Level
	Enabled(Level) bool

	IncludeCallerInfo() bool

	Write(e *Entry) error

	Close() error
	Flush() error
}
