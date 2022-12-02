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
)

// This should only be used in unit tests, as all log.Events
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
	// Always return an error on Flush and Handle methods.
	// This also skips incrementing EventCount.
	AlwaysErr bool
	// Number of Events pending to be written.
	// Flush will reset this counter.
	EventCount int
	// Events stores all the events passed to this handler.
	// Upon Flush(), this slice is cleared.
	Events []Event
	// Handler Level
	Level Level
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
func (m *MockHandler) Handle(e Event) error {
	m.HandleCount++
	if !m.AlwaysErr {
		m.EventCount++
		return nil
	}
	return ErrMockHandler
}

// Flush clears its internal Events slice.
// If AlwaysErr is true, then Events is not cleared
// and method returns an error.
func (m *MockHandler) Flush() error {
	if !m.AlwaysErr {
		m.EventCount = 0
		return nil
	}
	return ErrMockHandler
}
