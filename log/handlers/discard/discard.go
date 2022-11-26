package discard

import (
	"fmt"
	"sync"

	"github.com/tprasadtp/pkg/log"
)

var (
	_ log.Handler = &Handler{}
)

// Discard handler. Used for testing
type Handler struct {
	mu         sync.Mutex
	closed     bool
	id         string
	level      log.Level
	callerinfo bool
}

func New(id string, level log.Level) *Handler {
	return &Handler{
		id:    id,
		level: level,
	}
}

func NewWithCallerInfo(id string, level log.Level) *Handler {
	return &Handler{
		id:         id,
		level:      level,
		callerinfo: true,
	}
}

func (h *Handler) Init() error {
	return nil
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
	h.closed = true
	return nil
}

func (h *Handler) Enabled(l log.Level) bool {
	return h.level >= l
}

func (h *Handler) IncludeCallerInfo() bool {
	return h.callerinfo
}

func (h *Handler) Flush() error {
	h.mu.Lock()
	defer h.mu.Unlock()
	return nil
}

func (h *Handler) Write(e log.Event) error {
	if h.closed {
		return fmt.Errorf("log.handler.discard: Handler is closed")
	}
	h.mu.Lock()
	defer h.mu.Unlock()
	return nil
}
