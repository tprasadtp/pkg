package log

type Handler interface {
	Id() string
	Level() Level
	Enabled(Level) bool
	Write(e *Event) error
	Close() error
	Flush() error
}
