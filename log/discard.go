package log

// Compile time check for handler.
// This will fail if discard.Handler does not
// implement Handler interface.
var _ Handler = &NoOpHandler{}

// Discard Handler.
type NoOpHandler struct {
	level  Level
	closed bool
}

// New  returns a a new discard Handler. Unlike most handler constructors,
// this DOES NOT have a [io.Writer] as argument.
func NewNoOpHandler(l Level) *NoOpHandler {
	return &NoOpHandler{
		level: l,
	}
}

// Enabled Checks if given level is enabled.
func (h *NoOpHandler) Enabled(level Level) bool {
	return level >= h.level
}

// Write the Event.
func (h *NoOpHandler) Write(event Event) error {
	if h.closed {
		return ErrHandlerClosed
	}
	return nil
}

// Flushes the handler.
func (h *NoOpHandler) Flush() error {
	if h.closed {
		return ErrHandlerClosed
	}
	return nil
}

// Closes the handler.
func (h *NoOpHandler) Close() error {
	if h.closed {
		return ErrHandlerClosed
	}
	h.closed = true
	return nil
}
