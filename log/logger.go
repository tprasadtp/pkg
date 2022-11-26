package log

import "context"

// Logger
type Logger struct {
	handler   Handler
	ctx       context.Context
	namespace string
	err       error
	fields    []Field
}
