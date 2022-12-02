package log

// handlerError error.
type handlerError string

// Implements Error() interface on handlerError.
func (m handlerError) Error() string {
	return string(m)
}

const (
	// Error returned by MockHandler.
	ErrHandlerWrite = handlerError("handler write failed")
	// Error returned by MockHandler when writing or closing already closed handler.
	ErrHandlerClosed = handlerError("handler is closed")
)
