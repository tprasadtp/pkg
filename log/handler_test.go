package log

import (
	"fmt"
	"testing"
)

// No-Op Handler.
type testHandler struct {
	alwaysErr bool
	count     uint32
	flush     uint32
	events    []Event
	level     Level
}

// Enabled Checks if given level is enabled.
func (t *testHandler) Enabled(level Level) bool {
	return t.level >= level
}

// Handle the Event.
func (t *testHandler) Handle(e Event) error {
	t.count++
	t.events = append(t.events, e)
	if t.alwaysErr {
		return fmt.Errorf("test handler Handle() error")
	}
	return nil
}

// Flushes the handler, because this handler is no-op, flush in also a no-op.
func (t *testHandler) Flush() error {
	t.flush++
	if t.alwaysErr {
		return fmt.Errorf("test handler Flush() error")
	}
	return nil
}

func TestHandler(t *testing.T) {
	logger := New(&testHandler{})
	if logger != nil {
		t.Errorf("Test is not implemented")
	}
}
