package log

import "time"

type StackTrace struct{}

type CallerInfo struct {
	Package string
	File    string
	Line    string
}

// Log Events
type Event struct {
	Level      Level
	Timestamp  time.Time
	SpanID     string
	TraceID    string
	Message    string
	Error      error
	Namespaces []string
	Fields     []Field
	StackTrace StackTrace
	CallerInfo CallerInfo
}

type Field struct {
	Namespace string
	Key       string
	Value     any
}
