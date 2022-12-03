package log

// Compile time check for handler.
// This will fail if MockHandler does not implement Handler interface.
var _ Handler = &MockHandler{}

// This MUST only be used in unit tests, as all log.Events
// are simply appended to a slice, which can lead to
// memory exhaustion.
//
// MockHandler is a mock handler which is used for tests.
// This holds some counters for tracking state to be used in tests.
// This handler lacks sync semantics aka this is not concurrent safe.
// If you are looking for handler to use with testing.TB, see
// [github.com/tprasadtp/pkg/log/handlers/testing.Handler].
type MockHandler struct {
	// Number of times handler call invoked
	// This is incremented even when methods return an error.
	HandleCount int
	// Always return an error on Flush and Write methods.
	AlwaysErr bool
	// Events slice stores all the events successfully processed
	// but not yet flushed by the handler.
	//  - Upon Flush(), this slice is cleared if AlwaysErr not set to true.
	//  - If AlwaysErr is set to true, then it assumes that handler
	//    failed to write the event, thus not added to the slice.
	Events []Event
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
	m.HandleCount++
	if m.closed {
		return ErrHandlerClosed
	}
	if m.AlwaysErr {
		return ErrHandlerWrite
	}
	m.Events = append(m.Events, event)
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
	m.Events = []Event{}
	return nil
}

// Close clears its internal Events slice.
// and closes writing to this handler.
func (m *MockHandler) Close() error {
	if m.closed {
		return ErrHandlerClosed
	}
	m.closed = true
	m.Events = []Event{}
	return nil
}
