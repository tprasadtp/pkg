package log

// Package level global logger.
// Defaults to automatic detection.
var GlobalLogger Logger = Automatic()

// Returns a logger with automatic handler selection.
// In rare cases this can return a logger with more than one
// handler. However all handlers are guaranteed to be unique.
// By default all handlers are set to their default levels.
// Please note that this returns a Logger not a SubLogger.
func Automatic() Logger {
	return Logger{}
}
