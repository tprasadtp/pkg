package testdata

import "github.com/tprasadtp/pkg/log"

var events = []log.Event{
	{
		Level:   log.LevelDebug,
		Message: "Debug-01",
	},
	// 2 VERBOSE
	{
		Level:   log.LevelVerbose,
		Message: "Verbose-01",
	},
	{
		Level:   log.LevelVerbose,
		Message: "Verbose-02",
	},
	// 3 INFO
	{
		Level:   log.LevelInfo,
		Message: "Info-01",
	},
	{
		Level:   log.LevelInfo,
		Message: "Info-02",
	},
	{
		Level:   log.LevelInfo,
		Message: "Info-03",
	},
	// 4 WARNING
	{
		Level:   log.LevelWarning,
		Message: "Warning-01",
	},
	{
		Level:   log.LevelWarning,
		Message: "Warning-02",
	},
	{
		Level:   log.LevelWarning,
		Message: "Warning-03",
	},
	{
		Level:   log.LevelWarning,
		Message: "Warning-04",
	},
	// 5 ERROR
	{
		Level:   log.LevelError,
		Message: "Error-01",
	},
	{
		Level:   log.LevelError,
		Message: "Error-02",
	},
	{
		Level:   log.LevelError,
		Message: "Error-03",
	},
	{
		Level:   log.LevelError,
		Message: "Error-04",
	},
	{
		Level:   log.LevelError,
		Message: "Error-05",
	},
}

const (
	// Number of Events greater than or equal to DebugLevel returned by GetEvents.
	D = 15
	// Number of Events greater than or equal to VerboseLevel returned by GetEvents.
	V = 14
	// Number of Events greater than or equal to InfoLevel returned by GetEvents.
	I = 12
	// Number of Events greater than or equal to WarningLevel returned by GetEvents.
	W = 9
	// Number of Events greater than or equal to ErrorLevel returned by GetEvents.
	E = 5
	// Number  level Debug returned by GetEvents.
)

// GetEvents returns a sample of events.
// Ir has 1 Debug, 2 Verbose, 3 Info, 4 Warn and 5 ErrorLevel Events.
func GetEvents() []log.Event {
	return events
}
