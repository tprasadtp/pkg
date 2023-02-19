package log

import (
	"runtime"
	"sync"
	"time"
)

// Pooled events for alloc optimization.
var eventPool = sync.Pool{
	New: func() any {
		return &Event{
			Fields: make([]Field, 0, fieldsBucketSize),
		}
	},
}

// Get caller info.
func getCallerInfo(depth int) CallerInfo {
	pcs := make([]uintptr, 1)
	//nolint:gomnd // Skips runtime.Callers, and this function.
	runtime.Callers(depth+2, pcs)
	frames := runtime.CallersFrames(pcs)
	frame, _ := frames.Next()

	return CallerInfo{
		Defined: true,
		Line:    uint(frame.Line),
		Func:    frame.Function,
		File:    frame.File,
	}
}

// write is an internal method which writes event to log.Handler.
// All other named levels and methods use this in some form or other.
// This must be called directly by the method logging an event and not some
// wrapper as caller info might be wrong if done so.
func (l Logger) write(level Level, message string) {
	if l.handler == nil {
		return
	}

	// return if handler is not enabled
	if !l.handler.Enabled(level) {
		return
	}

	//nolint:errcheck // This linter is useless here.
	event := eventPool.Get().(*Event)
	defer func() {
		// Avoid large objects from poisoning the pool.
		// See https://golang.org/issue/23199
		const maxCap = 1 << 16
		if cap(event.Fields) < maxCap {
			event.Fields = event.Fields[:0]
			eventPool.Put(event)
		}
	}()

	// Build an event
	event.Namespace = l.namespace
	event.Ctx = l.ctx
	event.Error = l.err

	// Check if pool backed slice has enough capacity if not reallocate
	// in fieldsBucketSize increments.
	n := len(l.fields)
	if n > cap(event.Fields) {
		buckets := (n / fieldsBucketSize) + 1
		event.Fields = make([]Field, 0, fieldsBucketSize*buckets)
	}
	copy(event.Fields, l.fields)

	event.Time = time.Now()
	event.Level = level
	event.Message = message

	if l.caller {
		event.Caller = getCallerInfo(1)
	}

	if err := l.handler.Write(event); err != nil {
		panic(err)
	}
}
