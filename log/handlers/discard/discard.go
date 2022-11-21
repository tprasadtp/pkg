package discard

import (
	"sync"

	"github.com/tprasadtp/pkg/log"
)

// Discard handler. Used for testing
type Handler struct {
	mu  sync.RWMutex
	id  string
	lvl log.Level
}

func New(id string, l log.Level) *Handler {
	return &Handler{
		id:  id,
		lvl: l,
	}
}

func (h *Handler) Id() string {
	return "discard"
}

func (h *Handler) Level() log.Level {
	return h.lvl
}

func (h *Handler) Enabled(l log.Level) bool {
	return l >= h.lvl
}

func (h *Handler) Close() error {
	return nil
}

func (h *Handler) Flush() error {
	return nil
}

func (h *Handler) WriteEvent(e *log.Event) error {
	return nil
}
