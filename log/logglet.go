package log

import (
	"runtime"
	"sync"
	"time"

	"github.com/tprasadtp/pkg/log/internal/helpers"
)

const maxHelpers = 10

var callerPool = sync.Pool{
	New: func() any {
		return &caller{
			pcs: make([]uintptr, 64),
		}
	},
}

type caller struct {
	frames *runtime.Frames
	pcs    []uintptr
}

// write is an internal wrapper which writes event to log.Handler.
// All other named levels and methods use this with some form or other.
func (log Logger) write(level Level, message string, depth uint) {
	// // logger must not be nil.
	// if log.handler == nil {
	// 	panic(ErrLoggerInvalid)
	// }

	// Skip if handler is not enabled on the level
	if !log.handler.Enabled(level) {
		return
	}

	// build log Event
	event := Event{
		Level:   level,
		Context: log.ctx,
		Message: message,
		Error:   log.err,
		Time:    time.Now(),
	}

	// Caller Tracing
	caller := callerPool.Get().(*caller)

	// Skip two extra frames to account for this function
	// and runtime.Callers itself.
	//nolint:gomnd // ignore this magic number.
	n := runtime.Callers(int(depth+2), caller.pcs[:])
	caller.frames = runtime.CallersFrames(caller.pcs[:n])
	for i := 0; i < maxHelpers; i++ {
		frame, more := caller.frames.Next()
		_, helper := helpers.Map.Load(frame.Function)
		// We ran out of frames (This implies bug in log package)
		if !more {
			event.Caller = CallerInfo{
				Line: 0,
				File: "INVALID_FRAME",
				Func: "INVALID_FRAME",
			}
			break
		}

		if !helper {
			event.Caller = CallerInfo{
				Line: uint(frame.Line),
				File: frame.File,
				Func: frame.Function,
			}
			break
		}
	}
	log.handler.Write(event)
}
