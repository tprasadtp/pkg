package log

import (
	"runtime"
	"time"
)

const maxHelpers = 10

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

	// return if handler is not enabled
	if !log.handler.Enabled(level) {
		return
	}

	// build log Event
	event := Event{
		Level:   level,
		Message: message,
		Error:   log.err,
		Time:    time.Now(),
	}

	// // // Caller Tracing
	// // caller := callerPool.Get().(*caller)
	// // defer callerPool.Put(caller)
	// caller := caller{}

	// // Skip two extra frames to account for this function
	// // and runtime.Callers itself.
	// //nolint:gomnd // ignore this magic number.
	// n := runtime.Callers(int(depth+2), caller.pcs[:])
	// caller.frames = runtime.CallersFrames(caller.pcs[:n])
	// for i := 0; i < maxHelpers; i++ {
	// 	frame, more := caller.frames.Next()
	// 	_, helper := helpers.Map.Load(frame.Function)
	// 	// We ran out of frames (This implies bug in log package)
	// 	if !more {
	// 		event.Caller = CallerInfo{
	// 			Defined: true,
	// 			Line:    0,
	// 			File:    "INVALID_FRAME",
	// 			Func:    "INVALID_FRAME",
	// 		}
	// 		break
	// 	}

	// 	if !helper {
	// 		event.Caller = CallerInfo{
	// 			Defined: true,
	// 			Line:    uint(frame.Line),
	// 			File:    frame.File,
	// 			Func:    frame.Function,
	// 		}
	// 		break
	// 	}
	// }
	log.handler.Write(event)
}
