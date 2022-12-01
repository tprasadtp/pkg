package log

import (
	"fmt"
	"strconv"
)

// Level Level represents log level.
type Level uint16

const (
	DebugLevel    Level = 10
	VerboseLevel  Level = 20
	InfoLevel     Level = 30
	SuccessLevel  Level = 35
	NoticeLevel   Level = 40
	WarningLevel  Level = 50
	ErrorLevel    Level = 60
	CriticalLevel Level = 70
	FatalLevel    Level = 80
)

// String returns a name for the level.
// If the level has a name, then that name
// in uppercase is returned.
func (l Level) String() string {
	switch {
	case l < DebugLevel:
		return fmt.Sprintf("DEBUG-%d", DebugLevel-l)
	case l == DebugLevel:
		return "DEBUG"
	case l < VerboseLevel:
		return fmt.Sprintf("VERBOSE-%d", VerboseLevel-l)
	case l == VerboseLevel:
		return "VERBOSE"
	case l < InfoLevel:
		return fmt.Sprintf("INFO-%d", InfoLevel-l)
	case l == InfoLevel:
		return "INFO"
	case l < SuccessLevel:
		return fmt.Sprintf("SUCCESS-%d", SuccessLevel-l)
	case l == SuccessLevel:
		return "SUCCESS"
	case l < NoticeLevel:
		return fmt.Sprintf("NOTICE-%d", NoticeLevel-l)
	case l == NoticeLevel:
		return "NOTICE"
	case l < WarningLevel:
		return fmt.Sprintf("WARNING-%d", WarningLevel-l)
	case l == WarningLevel:
		return "WARNING"
	case l < ErrorLevel:
		return fmt.Sprintf("ERROR-%d", ErrorLevel-l)
	case l == ErrorLevel:
		return "ERROR"
	case l < CriticalLevel:
		return fmt.Sprintf("CRITICAL-%d", CriticalLevel-l)
	case l == CriticalLevel:
		return "CRITICAL"
	case l < FatalLevel:
		return fmt.Sprintf("FATAL-%d", FatalLevel-l)
	case l == FatalLevel:
		return "FATAL"
	default:
		return fmt.Sprintf("FATAL+%d", l-FatalLevel)
	}
}

// MarshalJSON implements json.Marshaler interface.
func (l Level) MarshalJSON() ([]byte, error) {
	// AppendQuote is sufficient for JSON-encoding all Level strings.
	// They don't contain any runes that would produce invalid JSON
	// when escaped.
	return strconv.AppendQuote(nil, l.String()), nil
}
