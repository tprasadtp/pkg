package testdata

import "github.com/tprasadtp/pkg/log"

var events = []log.Event{
	{
		Level:   log.DebugLevel,
		Message: "Debug-01",
	},
	// 2 VERBOSE
	{
		Level:   log.VerboseLevel,
		Message: "Verbose-01",
	},
	{
		Level:   log.VerboseLevel,
		Message: "Verbose-02",
	},
	// 3 INFO
	{
		Level:   log.InfoLevel,
		Message: "Info-01",
	},
	{
		Level:   log.InfoLevel,
		Message: "Info-02",
	},
	{
		Level:   log.InfoLevel,
		Message: "Info-03",
	},
	// 4 WARNING
	{
		Level:   log.WarningLevel,
		Message: "Warning-01",
	},
	{
		Level:   log.WarningLevel,
		Message: "Warning-02",
	},
	{
		Level:   log.WarningLevel,
		Message: "Warning-03",
	},
	{
		Level:   log.WarningLevel,
		Message: "Warning-04",
	},
	// 5 ERROR
	{
		Level:   log.ErrorLevel,
		Message: "Error-01",
	},
	{
		Level:   log.ErrorLevel,
		Message: "Error-02",
	},
	{
		Level:   log.ErrorLevel,
		Message: "Error-03",
	},
	{
		Level:   log.ErrorLevel,
		Message: "Error-04",
	},
	{
		Level:   log.ErrorLevel,
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
	// Number of Events greater than or equal to WArningLevel returned by GetEvents.
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
