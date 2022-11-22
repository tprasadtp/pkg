package log

import "time"

// var (
// 	_ Interface = Entry{}
// )

// Log Entry
type Entry struct {
	Logger     *Logger
	Level      Level
	Timestamp  time.Time
	SpanID     SpanID
	TraceID    TraceID
	Message    string
	Error      error
	Namespaces []string
	Fields     []Field
	StackTrace StackTrace
	CallerInfo CallerInfo
}
