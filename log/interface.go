package log

// Interface represents the API of both Logger and Entry.
type Interface interface {
	WithFields(Fielder) *Entry
	WithField(string, any) *Entry
	WithError(error) *Entry

	Debug(string)
	Info(string)
	Warn(string)
	Error(string)
	Fatal(string)
	// Panic(string)

	Debugf(string, ...any)
	Infof(string, ...any)
	Warnf(string, ...any)
	Errorf(string, ...any)
	Fatalf(string, ...any)
	// Panicf(string, ...any)

	WithoutPadding() *Entry
	// WithPadding(int) *Entry
	ResetPadding()
	IncreasePadding()
	DecreasePadding()
}
