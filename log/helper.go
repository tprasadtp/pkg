package log

import (
	"runtime"

	"github.com/tprasadtp/pkg/log/internal/helpers"
)

// Helper marks the calling function as a helper
// and skips it for source location information.
// It's the log's equivalent of testing.TB.Helper(), but with
// following limitations.
//   - This will ignore if called from main()
//   - There can be maximum of 10 helpers in a call stack.
func Helper() {
	pc, _, _, ok := runtime.Caller(1)
	if ok {
		f := runtime.FuncForPC(pc)
		// Ignore if called from main().
		if f.Name() != "main.main" {
			// We just want the function to be stored, make value as nil.
			helpers.Map.LoadOrStore(f.Name(), nil)
		}
	}
}
