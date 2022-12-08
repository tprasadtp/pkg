package log

import (
	"net"
	"time"
)

type RequestInfo struct {
	Valid bool

	RemoteAddr net.IP

	Proto string

	Host string
	URL  string

	Referer   string
	UserAgent string

	Method       string
	RequestSize  uint64
	ResponseCode uint64
	ResponseSize int64

	Latency time.Duration

	SpanID       string
	TraceID      string
	TraceSampled bool
}
