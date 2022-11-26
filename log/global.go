package log

// Returns a logger with automatic handler selection.
func Automatic() (*Logger, error) {
	return &Logger{}, nil
}
