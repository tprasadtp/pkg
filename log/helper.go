package log

import (
	"runtime"
	"sync"
)

// helpers stores helper's reference.
// sync.Map may be not the best fit here, but it is the easiest,
// and uses a well tested standard library code.
// because its use is well known and only limited to within
// this package it works fine.
//
//nolint:gochecknoglobals // This cannot be avoided, but it is not exported.
var helpers sync.Map

// Helper marks the calling function as a helper
// and skips it for source location information.
// It's the log's equivalent of testing.TB.Helper().
func Helper() {
	pc, _, _, ok := runtime.Caller(1)
	if ok {
		f := runtime.FuncForPC(pc)
		// We just want the function to be stored, make value as nil.
		helpers.LoadOrStore(f.Name(), nil)
	}
}
