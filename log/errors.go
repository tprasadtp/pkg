package log

// Ensure all custom errors implement error interface.
var (
	_ error = ErrLoggerInvalid
	_ error = ErrHandlerWrite
	_ error = ErrHandlerClosed
)

const (
	// Error returned when invalid logger or nil is used.
	ErrLoggerInvalid = handlerError("logger is invalid or nil")
	// Error returned when write or flush methods fail.
	ErrHandlerWrite = handlerError("handler write failed")
	// Error returned when writing, flushing or closing an
	// already closed handler.
	ErrHandlerClosed = handlerError("handler is closed")
)

// handlerError error.
type handlerError string

// Implements Error() interface on handlerError.
func (m handlerError) Error() string {
	return string(m)
}

// loggerError error.
type loggerError string

// Implements Error() interface on handlerError.
func (m loggerError) Error() string {
	return string(m)
}
