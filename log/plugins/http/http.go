package log

import (
	"time"
)

type RequestInfo struct {
	Valid bool

	TraceSampled bool
	SpanID       string
	TraceID      string

	RemoteAddr string

	Proto string

	Host string
	URL  string

	Referer   string
	UserAgent string

	Method       string
	RequestSize  uint64
	ResponseCode uint64
	ResponseSize int64
	Latency      time.Duration
}
