package log

import (
	"strings"
	"sync"
	"time"
)

var eventPool = sync.Pool{
	New: func() any {
		return &Event{
			Fields: make([]Field, 0, fieldsBucketSize),
		}
	},
}

// Field is logger fields. Logger contains a slice of [Fields]
// Optionally with a namespace.
type Field struct {
	Namespace string
	Key       string
	Value     any
}

// Includes caller info if available.
type CallerInfo struct {
	// Defined represents whether caller entry is defined.
	Defined bool
	// Line number of the caller
	// If not available, this is 0.
	Line uint
	// File containing the code
	// If not available this is empty string.
	File string
	// Function name of the caller.
	// this includes full path of the package.
	// except for main package.
	// This is empty if information is not available.
	Func string
}

// Event represents a single Log event. Event should be considered immutable.
// If underlying handler pools events it must store it in storage backed by
// its own pool or arrays. Logger will release the event back to the pool
type Event struct {
	// Namespace is namespace of the logger that generated this event.
	Namespace string

	// Time (Global)
	Time time.Time

	// Log Level (Global)
	Level Level

	// Log Message (Global)
	Message string

	// Error (Global)
	Error error

	// Caller
	Caller CallerInfo

	Fields []Field
}

// Returns a new Field optionally with a namespace.
func F(key string, value any, namespace ...string) Field {
	switch len(namespace) {
	case 0:
		return Field{
			Key:   key,
			Value: ToValue(value),
		}
	case 1:
		return Field{
			Namespace: namespace[0],
			Key:       key,
			Value:     ToValue(value),
		}
	default:
		return Field{
			Namespace: strings.Join(namespace, "."),
			Key:       key,
			Value:     ToValue(value),
		}
	}
}
