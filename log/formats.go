package log

// TimeFormat is time format to use in handler.
// This is exported in log package so that all handlers can agree on
// a single constant.
type TimeFormat uint8

const (
	// Unix timestamp.
	TimeFormatUnix TimeFormat = iota
	// Unix Timestamp with nano seconds.
	TimeFormatUnixNano
	// RFC 2822 format.
	TimeRFC2822
)
