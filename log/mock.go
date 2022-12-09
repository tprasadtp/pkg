package log

import "fmt"

// Compile time check for handler.
// This will fail if MockHandler does not implement Handler interface.
var _ Handler = &MockHandler{}

// MockHandler is a mock handler which is used for tests.
// This holds some counters for tracking state to be used in tests.
// This handler lacks sync semantics aka this is not concurrent safe.
type MockHandler struct {
	// Number of times handler call invoked Write
	// This is incremented even when methods return an error.
	WriteCalls uint
	// Number of times handler successfully saved an event entry.
	EventsWritten uint
	// Always return an error on Flush and Write methods.
	AlwaysErr bool
	// Handler Level
	Level Level
	// Closed
	closed bool
}

// Enabled can be replaced with a custom function.
// If EnabledFunc is set to nil, this will return true.
func (m *MockHandler) Enabled(level Level) bool {
	return level >= m.Level
}

// Handle simply saves the event in to Events slice.
// If AlwaysErr is true, then event is not saved it internal slice,
// and method returns an error.
func (m *MockHandler) Write(event Event) error {
	m.WriteCalls++
	if m.closed {
		return ErrHandlerClosed
	}
	if m.AlwaysErr {
		return ErrHandlerWrite
	}
	m.EventsWritten++
	return nil
}

// Flush clears its internal Events slice.
// If AlwaysErr is true, then Events is not cleared
// and method returns an error.
func (m *MockHandler) Flush() error {
	if m.closed {
		return ErrHandlerClosed
	}
	if m.AlwaysErr {
		return ErrHandlerWrite
	}
	m.EventsWritten = 0
	return nil
}

// Close clears its internal Events slice.
// and closes writing to this handler.
func (m *MockHandler) Close() error {
	if m.closed {
		return fmt.Errorf("mock handler error: %w", ErrHandlerClosed)
	}
	m.closed = true
	m.EventsWritten = 0
	return nil
}
