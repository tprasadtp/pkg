package log

import (
	"runtime"
	"sync"
	"time"
)

// Pooled callers for alloc optimization.
var callerPool = sync.Pool{
	New: func() any {
		println("CREATE - CALLERS_POOL")
		return &caller{
			pcs:    nil,
			frames: nil,
		}
	},
}

// caller holds runtime.Frames and slice of program counters
// to determine caller of the log function. This is pooled,
// to reduce allocations, as [runtime.Callers] returns a pointer.
type caller struct {
	pcs     []uintptr // program counters, always a sub-slice of storage.
	storage [10]uintptr
	frames  *runtime.Frames
}

// Reset resets pointer to be eligible for pool.
// Allocation is still is in the heap.
func (c *caller) Reset() {
	c.frames = nil
	c.pcs = nil
}

func getCallerInfo(depth int) CallerInfo {
	//nolint:errcheck // This linter is useless here.
	caller := callerPool.Get().(*caller)
	caller.pcs = caller.storage[:1]
	defer func() {
		caller.Reset()
		callerPool.Put(caller)
		println("PUT - CALLERS_POOL")
	}()
	//nolint:gomnd // Skips runtime.Callers, and this function.
	runtime.Callers(depth+2, caller.pcs)
	caller.frames = runtime.CallersFrames(caller.pcs)
	frame, _ := caller.frames.Next()

	return CallerInfo{
		Defined: true,
		Line:    uint(frame.Line),
		Func:    frame.Function,
		File:    frame.File,
	}
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

	return log.handler.Write(event)
}
