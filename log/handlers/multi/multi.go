// Package multi implements a Handler which wraps multiple handlers.
package multi

import (
	"fmt"

	"github.com/tprasadtp/pkg/log"
)

// Compile time check for handler.
// This will fail if multi.Handler does not
// implement log.Handler interface.
var _ log.Handler = &Handler{}

// No-Op Handler.
type Handler struct {
	handlers []log.Handler
}

// New returns a handler which wraps other handlers.
func New(handlers ...log.Handler) *Handler {
	return &Handler{
		handlers: handlers,
	}
}

// Enabled Checks if given level is enabled on ANY of the handlers.
func (m *Handler) Enabled(level log.Level) bool {
	for _, h := range m.handlers {
		if h.Enabled(level) {
			return true
		}
	}
	return false
}

// Passes the event to all the handlers and let them handle it.
func (m *Handler) Handle(event log.Event) error {
	var err error
	for _, h := range m.handlers {
		if h.Enabled(event.Level) {
			if eloop := h.Handle(event); eloop != nil {
				if err != nil {
					err = fmt.Errorf("%w", eloop)
				} else {
					err = eloop
				}
			}
		}
	}
	return err
}

// Flushes all the handlers.
func (m *Handler) Flush() error {
	var err error
	for _, h := range m.handlers {
		if eloop := h.Flush(); eloop != nil {
			if err != nil {
				err = fmt.Errorf("%w", eloop)
			} else {
				err = eloop
			}
		}
	}
	return err
}
