package log

import (
	"runtime"
	"sync"
	"time"
)

var framesPool = sync.Pool{
	New: func() any {
		return &caller{
			pcs: make([]uintptr, 20),
		}
	},
}

type caller struct {
	frames *runtime.Frames
	pcs    []uintptr
}

// write is an internal wrapper which writes event to log.Handler.
// All other named levels and methods use this with some form or other.
// this must be called directly by the method logging an event and not some
// wrapper as caller info might be wrong if done so.
func (log Logger) write(level Level, message string) error {
	// logger handler must not be nil.
	if log.handler == nil {
		panic(ErrLoggerInvalid)
	}

	// return if handler is not enabled
	if !log.handler.Enabled(level) {
		return nil
	}

	// build log Event
	event := Event{
		Level:   level,
		Message: message,
		Error:   log.err,
		Time:    time.Now(),
	}

	// Caller Tracing
	var pcs [1]uintptr
	const depth = 3
	runtime.Callers(depth, pcs[:])
	frames, _ := runtime.CallersFrames(pcs[:]).Next()
	event.Caller = Caller{
		Line: uint(frames.Line),
		File: frames.File,
		Func: frames.Function,
	}

	return log.handler.Write(event)
}
