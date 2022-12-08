package log

// Ensures all custom errors implement the error interface.
var (
	_ error = handlerError("")
	_ error = loggerError("")
)

const (
	// Error returned when logger is invalid.
	ErrLoggerInvalid = loggerError("logger is invalid or nil")
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

// var (
// 	// Error returned when logger is invalid.
// 	ErrLoggerInvalid = errors.New("logger is invalid or nil")
// 	// Error returned when write or flush methods fail.
// 	ErrHandlerWrite = errors.New("handler write failed")
// 	// Error returned when writing, flushing or closing an
// 	// already closed handler.
// 	ErrHandlerClosed = errors.New("handler is closed")
// )
