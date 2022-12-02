// Package multi implements a Handler which wraps multiple handlers.
package multi

import (
	"fmt"

	"github.com/tprasadtp/pkg/log"
	"go.uber.org/multierr"
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
			errPerH := h.Handle(event)
			fmt.Printf("left=%s, right=%s\n", err, errPerH)
			err = multierr.Append(err, errPerH)
		}
	}
	return err
}

// Flushes the handler, because this handler is no-op, flush in also a no-op.
func (m *Handler) Flush() error {
	var err error
	for _, h := range m.handlers {
		err = multierr.Append(err, h.Flush())
	}
	return err
}
