package log

import (
	"runtime"
	"sync"
	"time"

	"github.com/tprasadtp/pkg/log/internal/helpers"
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
	const maxHelpers = 10
	pcs := make([]uintptr, 1)
	caller := CallerInfo{}
	// Skip two extra frames to account for this function
	// and runtime.Callers itself.
	//nolint:gomnd // ignore this magic number.
	n := runtime.Callers(int(depth+2), pcs[:])
	frames := runtime.CallersFrames(pcs[:n])
	for i := 0; i < maxHelpers; i++ {
		frame, more := frames.Next()
		_, helper := helpers.Map.Load(frame.Function)
		// We ran out of frames (This implies bug in log package)
		if !more {
			caller = CallerInfo{
				Defined: true,
				Line:    0,
				File:    "INVALID_FRAME",
				Func:    "INVALID_FRAME",
			}
			break
		}

		if !helper {
			caller = CallerInfo{
				Defined: true,
				Line:    uint(frame.Line),
				File:    frame.File,
				Func:    frame.Function,
			}
			break
		}
	}

	return caller
}

// write is an internal method which writes event to log.Handler.
// All other named levels and methods use this with some form or other.
// This must be called directly by the method emitting an event and not some
// wrapper as caller info might be wrong if done so.
func (log Logger) write(level Level, message string) {
	// logger handler must not be nil.
	if log.handler == nil {
		panic(ErrLoggerInvalid)
	}

	// return if handler is not enabled
	if !log.handler.Enabled(level) {
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
	event.Namespace = log.namespace
	event.Ctx = log.ctx
	event.Error = log.err
	// Check if pool backed slice has enough capacity if not reallocate
	// in fieldsBucketSize increments.
	n := len(log.fields)
	if n > cap(event.Fields) {
		buckets := (n / fieldsBucketSize) + 1
		event.Fields = make([]Field, fieldsBucketSize*buckets)
	}
	copy(event.Fields, log.fields)

	event.Time = time.Now()
	event.Level = level
	event.Message = message

	if log.caller {
		event.Caller = getCallerInfo(1)
	}

	if err := log.handler.Write(event); err != nil {
		panic(err)
	}
}
