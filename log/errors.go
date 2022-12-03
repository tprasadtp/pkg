package log

// handlerError error.
type handlerError string

// Implements Error() interface on handlerError.
func (m handlerError) Error() string {
	return string(m)
}

const (
	// Error returned when write or flush methods fail.
	ErrHandlerWrite = handlerError("handler write failed")
	// Error returned when writing, flushing or closing an
	// already closed handler.
	ErrHandlerClosed = handlerError("handler is closed")
)
