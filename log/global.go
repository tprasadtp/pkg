package log

import "os"

type AutomaticOptions struct {
	// By default logs are written to Windows event log
	// when running as a windows service.
	WindowsEventLogSource string

	// By default linux services running under systemd
	// log to journald by default, or when running via
	// Terminal use the [os.Stderr].
	// This is a fallback when not running as systemd unit
	// and not running in terminal.
	File *os.File
}

// Returns a logger with automatic handler selection.
func Automatic(o *AutomaticOptions) (*Logger, error) {
	return &Logger{}, nil
}
