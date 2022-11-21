package discard

import (
	"sync"

	"github.com/tprasadtp/pkg/log"
)

// Discard handler. Used for testing
type Handler struct {
	mu     sync.Mutex
	closed bool
	id     string
	level  log.Level
}

func New(id string, level log.Level) *Handler {
	return &Handler{
		id:    id,
		level: level,
	}
}

func (h *Handler) Id() string {
	return "discard"
}

func (h *Handler) Level() log.Level {
	return h.level
}

func (h *Handler) Close() error {
	h.mu.Lock()
	defer h.mu.Unlock()
	return nil
}

func (h *Handler) Flush() error {
	return nil
}

func (h *Handler) Write(e *log.Entry) error {
	h.mu.Lock()
	defer h.mu.Unlock()
	if h.closed {
		return nil
	}
	return nil
}
