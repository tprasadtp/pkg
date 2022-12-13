package log

// Compile time check for handler.
// This will fail if discard.Handler does not
// implement Handler interface.
var _ Handler = &DiscardHandler{}

// Discard Handler.
type DiscardHandler struct {
	level  Level
	closed bool
}

// New  returns a a new discard Handler. Unlike most handler constructors,
// this DOES NOT have a [io.Writer] as argument.
func NewDiscardHandler(l Level) *DiscardHandler {
	return &DiscardHandler{
		level: l,
	}
}

// Enabled Checks if given level is enabled.
func (h *DiscardHandler) Enabled(level Level) bool {
	return level >= h.level
}

// Write the Event.
func (h *DiscardHandler) Write(event Event) error {
	if h.closed {
		return ErrHandlerClosed
	}
	return nil
}

// Flushes the handler.
func (h *DiscardHandler) Flush() error {
	if h.closed {
		return ErrHandlerClosed
	}
	return nil
}

// Closes the handler.
func (h *DiscardHandler) Close() error {
	if h.closed {
		return ErrHandlerClosed
	}
	h.closed = true
	return nil
}
