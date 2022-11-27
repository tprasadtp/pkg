// Package noop implements a no-op [github.com/tprasadtp/pkg/log.Handler].
// This [Handler] simply discards the log events.
package noop

import (
	"io"
	"sync"

	"github.com/tprasadtp/pkg/log"
)

// No-Op Handler
type Handler struct {
	mu    sync.Mutex
	w     io.Writer
	level log.Level
}

// Creates a new No-Op Handler. Unlike most handler constructors,
// this DOES NOT have a [io.Writer] as argument.
func New(l log.Level) *Handler {
	return &Handler{
		w:     io.Discard,
		level: l,
	}
}

// Enabled Checks if given level is enabled.
func (h *Handler) Enabled(level log.Level) bool {
	return h.level >= level
}

// Handle the Event
func (h *Handler) Handle(e log.Entry) error {
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
