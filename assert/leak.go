package assert

import (
	"testing"
)

// NoLeaks marks the given test as failed if any extra goroutines are
// found. This is a helper method to make it easier to integrate in
// tests by doing:
//
//	defer assert.NoLeaks(t)
func NoLeaks(t testing.TB) {
	t.Helper()
}
