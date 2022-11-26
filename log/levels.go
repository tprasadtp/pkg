package log

import (
	"fmt"
	"math"
	"strconv"
)

// Level Level represents log level.
type Level uint16

const (
	DEBUG    Level = 10
	VERBOSE  Level = 20
	INFO     Level = 30
	SUCCESS  Level = 40
	NOTICE   Level = 50
	WARNING  Level = 60
	ERROR    Level = 70
	CRITICAL Level = 80
)

const (
	ALL  Level = 0
	NONE Level = math.MaxUint16
)

// String returns a name for the level.
// If the level has a name, then that name
// in uppercase is returned.
// If the level is between named values, then
// an integer is appended to the uppercase name.
// Examples:
//
//	DEBUG.String() => "DEBUG"
//	(DEBUG-2).String() => "DEBUG-2"
func (l Level) String() string {
	switch {
	case l < DEBUG:
		return fmt.Sprintf("DEBUG-%d", DEBUG-l)
	case l == DEBUG:
		return "DEBUG"
	case l < VERBOSE:
		return fmt.Sprintf("VERBOSE-%d", VERBOSE-l)
	case l == VERBOSE:
		return "VERBOSE"
	case l < INFO:
		return fmt.Sprintf("INFO-%d", INFO-l)
	case l == INFO:
		return "INFO"
	case l < SUCCESS:
		return fmt.Sprintf("SUCCESS-%d", SUCCESS-l)
	case l == SUCCESS:
		return "SUCCESS"
	case l < NOTICE:
		return fmt.Sprintf("SUCCESS-%d", NOTICE-l)
	case l == NOTICE:
		return "NOTICE"
	case l < WARNING:
		return fmt.Sprintf("WARNING-%d", WARNING-l)
	case l == WARNING:
		return "WARNING"
	case l < ERROR:
		return fmt.Sprintf("WARNING-%d", ERROR-l)
	case l == ERROR:
		return "ERROR"
	case l < CRITICAL:
		return fmt.Sprintf("WARNING-%d", CRITICAL-l)
	case l == CRITICAL:
		return "CRITICAL"
	default:
		return fmt.Sprintf("CRITICAL+%d", l-CRITICAL)
	}
}

func (l Level) MarshalJSON() ([]byte, error) {
	// AppendQuote is sufficient for JSON-encoding all Level strings.
	// They don't contain any runes that would produce invalid JSON
	// when escaped.
	return strconv.AppendQuote(nil, l.String()), nil
}
