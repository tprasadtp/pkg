package helpers

import "sync"

// Map stores helper function names.
//
//nolint:gochecknoglobals // This MUST be global by design.
var Map sync.Map
