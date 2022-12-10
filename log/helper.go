package log

import (
	"runtime"

	"github.com/tprasadtp/pkg/log/internal/helpers"
)

// Helper marks the calling function as a helper
// and skips it for source location information.
// It's the log's equivalent of testing.TB.Helper(), but with
// following limitations. This will ignore if called from main()
// logger limits maximum number of helpers in a call to 10.
func Helper() {
	pc, _, _, ok := runtime.Caller(1)
	if ok {
		frames := runtime.CallersFrames([]uintptr{pc})
		frame, _ := frames.Next()
		// Ignore if called from main().
		if frame.Function != "main.main" && frame.Function != "" {
			// We just want the function to be stored, make value as nil.
			helpers.Map.LoadOrStore(frame.Function, nil)
		}
	}
}
