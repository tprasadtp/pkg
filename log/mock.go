package log

// Compile time check for handler.
// This will fail if MockHandler does not implement Handler interface.
var _ Handler = &MockHandler{}

// mockError error used for testing.
type mockError string

// Implements Error() interface on mockError.
func (m mockError) Error() string {
	return string(m)
}

const (
	// Error returned by MockHandler.
	ErrMockHandler = mockError("mock handler error")
	// Error returned by MockHandler when writing or closing already closed handler.
	ErrMockHandlerClosed = mockError("mock handler is closed")
)

// This should only be used in unit tests, as all log.Events
// are simply appended to a slice, which can lead to
// memory exhaustion.
//
// MockHandler is a mock handler which is used for tests.
// This holds some counters for tracking state to be used in tests.
// This handler lacks sync semantics aka this is not concurrent safe.
// If you are looking for handler to use with testing.TB, see
// [github.com/tprasadtp/pkg/log/handlers/testing.Handler].
type MockHandler struct {
	// Replaces Enabled() with custom function
	EnabledFunc func(Level) bool
	// Number of times handler cal invoked
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
	if m.EnabledFunc == nil {
		return level >= m.Level
	}
	return m.EnabledFunc(level)
}

// Handle simply saves the event in to Events slice.
// If AlwaysErr is true, then event is not saved it internal slice,
// and method returns an error.
func (m *MockHandler) Write(event Event) error {
	m.HandleCount++
	if m.closed {
		return ErrMockHandlerClosed
	}
	if m.AlwaysErr {
		return ErrMockHandler
	}
	m.Events = append(m.Events, event)
	return nil
}

// Flush clears its internal Events slice.
// If AlwaysErr is true, then Events is not cleared
// and method returns an error.
func (m *MockHandler) Flush() error {
	if m.closed {
		return ErrMockHandlerClosed
	}
	if m.AlwaysErr {
		return ErrMockHandler
	}
	m.Events = []Event{}
	return nil
}

// Close clears its internal Events slice.
// and closes writing to this handler.
func (m *MockHandler) Close() error {
	if m.closed {
		return ErrMockHandlerClosed
	}
	m.closed = true
	m.Events = []Event{}
	return nil
}
