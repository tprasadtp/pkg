package log_test

import "github.com/tprasadtp/pkg/log"

var events = []log.Event{
	// 1 DEBUG
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
