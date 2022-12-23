package log

import (
	"sync"

	"go.opentelemetry.io/otel/trace"
)

// Compile time check for DiscardHandler.
var _ Handler = &DiscardHandler{}

// NewDiscardHandler returns a a new discard Handler.
func NewDiscardHandler(l Level) *DiscardHandler {
	return &DiscardHandler{
		level: l,
	}
}

// Discard Handler discards all events written to it.
type DiscardHandler struct {
	mu     sync.Mutex
	level  Level
	closed bool
}

// Enabled Checks if given level is enabled.
func (h *DiscardHandler) Enabled(level Level) bool {
	return level >= h.level
}

// Write the Event.
func (h *DiscardHandler) Write(event *Event) error {
	h.mu.Lock()
	defer h.mu.Unlock()

	_ = trace.SpanContextFromContext(event.Ctx)

	if h.closed {
		return ErrHandlerClosed
	}
	return nil
}

// Flushes the handler.
func (h *DiscardHandler) Flush() error {
	h.mu.Lock()
	defer h.mu.Unlock()
	if h.closed {
		return ErrHandlerClosed
	}
	return nil
}

// Closes the handler.
func (h *DiscardHandler) Close() error {
	h.mu.Lock()
	defer h.mu.Unlock()
	if h.closed {
		return ErrHandlerClosed
	}
	h.closed = true
	return nil
}
