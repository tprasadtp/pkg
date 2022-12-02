// Package discard implements a no-op [github.com/tprasadtp/pkg/log.Handler].
// This [Handler] simply discards the log events. This handler is safe for
// concurrent usage.
package discard

import (
	"sync"

	"github.com/tprasadtp/pkg/log"
)

// Compile time check for handler.
// This will fail if discard.Handler does not
// implement log.Handler interface.
var _ log.Handler = &Handler{}

// Discard Handler.
type Handler struct {
	mu     sync.Mutex
	level  log.Level
	closed bool
}

// New  returns a a new discard Handler. Unlike most handler constructors,
// this DOES NOT have a [io.Writer] as argument.
func New(l log.Level) *Handler {
	return &Handler{
		level: l,
	}
}

// Enabled Checks if given level is enabled.
func (h *Handler) Enabled(level log.Level) bool {
	return level >= h.level
}

// Write the Event.
func (h *Handler) Write(event log.Event) error {
	h.mu.Lock()
	defer h.mu.Unlock()
	if h.closed {
		return log.ErrHandlerClosed
	}
	return nil
}

// Flushes the handler.
func (h *Handler) Flush() error {
	h.mu.Lock()
	defer h.mu.Unlock()
	if h.closed {
		return log.ErrHandlerClosed
	}
	return nil
}

// Closes the handler.
func (h *Handler) Close() error {
	h.mu.Lock()
	defer h.mu.Unlock()
	if h.closed {
		return log.ErrHandlerClosed
	}
	h.closed = true
	return nil
}
