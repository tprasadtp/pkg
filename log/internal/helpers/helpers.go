// Global holds log's internal global state.
package helpers

import "sync"

// Map stores helper function names.
// sync.Map may be not the best fit here, but it is the easiest,
// and uses a well tested standard library code.
//
//nolint:gochecknoglobals // This MUST be global by design.
var Map sync.Map
