package log

import (
	"fmt"
)

// Compile time check for handler.
// This will fail if discard.Handler does not
// implement Handler interface.
var _ Handler = &ConsoleHandler{}

func NewConsole() *ConsoleHandler {
	return &ConsoleHandler{
		level: DebugLevel,
	}
}

type ConsoleHandler struct {
	level  Level
	closed bool
}

// Enabled Checks if given level is enabled.
func (h *ConsoleHandler) Enabled(level Level) bool {
	return true
}

// Write the Event.
func (h *ConsoleHandler) Write(event Event) error {
	if h.closed {
		return ErrHandlerClosed
	}
	fmt.Printf("%+v\n\n", event)
	return nil
}

// Flushes the handler.
func (h *ConsoleHandler) Flush() error {
	if h.closed {
		return ErrHandlerClosed
	}
	return nil
}

// Closes the handler.
func (h *ConsoleHandler) Close() error {
	if h.closed {
		return ErrHandlerClosed
	}
	h.closed = true
	return nil
}
