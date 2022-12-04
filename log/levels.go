package log

// Level represents log level.
// Zero value of Level represents InfoLevel, which
// is what most people would expect.
type Level int

// Named Log Levels. Level Constants do not really matter,
// as handlers can remap them anyway. There should be plenty of
// headroom for custom levels.
const (
	// Trace level logs include very low level logs.
	// Stuff like raw packet dumps.
	TraceLevel Level = -30

	// DebugLevel is lowest named level and typically used for debugging.
	DebugLevel Level = -20

	// VerboseLevel is used when you need more insights on inner workings
	// of a program/application. This typically includes just enough information
	// for debugging stuff like network issues. For low level application
	// internals use DebugLevel or TraceLevel.
	VerboseLevel Level = -10

	// InfoLevel is typical level where application logs its events. This is includes
	// stuff like request logs, and user presentable information.
	// This is set to 0 so that zero value of Level is valid and most commonly used
	// level.
	InfoLevel Level = 0

	// SuccessLevel is mostly tailored for CLI applications,
	// and usually you do not need it in a web/server application.
	SuccessLevel Level = 10

	// NoticeLevel is something important like application
	// live reload or any other significant events.
	NoticeLevel Level = 20

	// WarningLevel is for errors or events which application
	// handled gracefully by using a fallback experience or retries.
	WarningLevel Level = 30

	// ErrorLevel is for application level errors.
	// These should trigger an alert in your APM solution.
	ErrorLevel Level = 40

	// CriticalLevel is for errors which have the potential to disrupt the application
	// but for the moment application can handle it. panic/recover flow can log at
	// this level. Stuff like running on low disk space, nearing API quota etc.
	CriticalLevel Level = 50

	// FatalLevel is for errors which lead to application crashes and application
	// cannot recover from this type of error.
	// If event is logged at this level, Logger will invoke Flush method on its
	// handler to avoid losing log data. However, this is not guaranteed.
	// Logger.Fatal() will invoke defined exit function. (defaults to [os.Exit]).
	FatalLevel Level = 90
)
