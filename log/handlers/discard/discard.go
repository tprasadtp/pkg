// Package discard implements a no-op [github.com/tprasadtp/pkg/log.Handler].
// This [Handler] simply discards the log events.
package discard

import (
	"sync"

	"github.com/tprasadtp/pkg/log"
)

// Compile time check for handler.
// This will fail if discard.Handler does not
// implement log.Handler interface.
var _ log.Handler = &Handler{}

// No-Op Handler.
type Handler struct {
	mu    sync.Mutex
	level log.Level
}

// New  returns a a new no-op Handler. Unlike most handler constructors,
// this DOES NOT have a [io.Writer] as argument.
func New(l log.Level) *Handler {
	return &Handler{
		level: l,
	}
}

// Enabled Checks if given level is enabled.
func (h *Handler) Enabled(level log.Level) bool {
	return h.level >= level
}

// Handle the Event.
func (h *Handler) Handle(e log.Event) error {
	h.mu.Lock()
	defer h.mu.Unlock()
	return nil
}

// Flushes the handler, because this handler is no-op, flush in also a no-op.
func (h *Handler) Flush() error {
	h.mu.Lock()
	defer h.mu.Unlock()
	return nil
}
