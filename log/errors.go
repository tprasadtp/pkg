package log

// Ensures all custom errors implement the error interface.
var (
	_ error = handlerError("")
	_ error = loggerError("")
)

const (
	// Error returned when logger is invalid.
	ErrLoggerInvalid = loggerError("log: logger is invalid or nil")
	// LoggerInvalidKind
	ErrInvalidKind = loggerError("log: kind mismatch")
	// Error returned when write or flush methods fail.
	ErrHandlerWrite = handlerError("log: handler write failed")
	// Error returned when writing, flushing or closing an
	// already closed handler.
	ErrHandlerClosed = handlerError("log: handler is closed")
)

// handlerError error.
type handlerError string

// Implements Error() interface for handlerError.
func (h handlerError) Error() string {
	return string(h)
}

// loggerError error.
type loggerError string

// Implements Error() interface for handlerError.
func (l loggerError) Error() string {
	return string(l)
}

// var (
// 	// Error returned when logger is invalid.
// 	ErrLoggerInvalid = errors.New("logger is invalid or nil")
// 	// Error returned when write or flush methods fail.
// 	ErrHandlerWrite = errors.New("handler write failed")
// 	// Error returned when writing, flushing or closing an
// 	// already closed handler.
// 	ErrHandlerClosed = errors.New("handler is closed")
// )
