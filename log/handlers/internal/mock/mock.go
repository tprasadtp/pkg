package mock

import (
	"fmt"

	"github.com/tprasadtp/pkg/log"
)

// Compile time check for handler.
// This will fail if MockHandler does not implement log.Handler interface.
var _ log.Handler = &Handler{}

// MockHandler is a mock handler which is used for tests.
// This holds some counters for tracking state to be used in tests.
// This handler lacks sync semantics aka this is not concurrent safe.
type Handler struct {
	// Number of times handler call invoked Write
	// This is incremented even when methods return an error.
	WriteCalls uint
	// Number of times handler successfully saved an event entry.
	EventsWritten uint
	// Always return an error on Flush and Write methods.
	AlwaysErr bool
	// Handler Level
	Level log.Level
	// Closed
	closed bool
}

// Enabled can be replaced with a custom function.
// If EnabledFunc is set to nil, this will return true.
func (m *Handler) Enabled(level log.Level) bool {
	return level >= m.Level
}

// Handle simply saves the event in to Events slice.
// If AlwaysErr is true, then event is not saved it internal slice,
// and method returns an error.
func (m *Handler) Write(event log.Event) error {
	m.WriteCalls++
	if m.closed {
		return log.ErrHandlerClosed
	}
	if m.AlwaysErr {
		return log.ErrHandlerWrite
	}
	m.EventsWritten++
	return nil
}

// Flush clears its internal Events slice.
// If AlwaysErr is true, then Events is not cleared
// and method returns an error.
func (m *Handler) Flush() error {
	if m.closed {
		return log.ErrHandlerClosed
	}
	if m.AlwaysErr {
		return log.ErrHandlerWrite
	}
	m.EventsWritten = 0
	return nil
}

// Close clears its internal Events slice.
// and closes writing to this handler.
func (m *Handler) Close() error {
	if m.closed {
		return fmt.Errorf("mock handler error: %w", log.ErrHandlerClosed)
	}
	m.closed = true
	m.EventsWritten = 0
	return nil
}
