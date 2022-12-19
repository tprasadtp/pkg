package log

// Level represents log level.
// Zero value of Level represents LevelInfo,
// which is what most people would expect.
type Level int

// Named Log Levels. Level Constants do not really matter,
// as handlers can remap them anyway. There should be plenty of
// headroom for custom levels.
const (
	// Trace level logs include very low level logs.
	// Stuff like raw packet dumps.
	LevelTrace Level = -30

	// Debug Level is lowest named level and typically used for debugging.
	LevelDebug Level = -20

	// LevelVerbose is used when you need more insights on inner workings
	// of a program/application. This typically includes just enough information
	// for debugging stuff like network issues. For low level application
	// internals use DebugLevel or TraceLevel.
	LevelVerbose Level = -10

	// LevelInfo is typical level where application logs its events.
	// This is includes stuff like access logs, and user presentable
	// information. This is set to 0 so that zero value of Level is
	// valid and most commonly used level.
	LevelInfo Level = 0

	// LevelSuccess is mostly tailored for CLI applications,
	// and usually you do not need it in a web/server application.
	LevelSuccess Level = 10

	// LevelNotice is something important like application
	// live reload, configuration change or any other significant events.
	LevelNotice Level = 20

	// LevelWarning is for errors or events which are not optimal/ideal,
	// but application handled gracefully by using a fallback experience
	// or in case of remote resources retries.
	LevelWarning Level = 30

	// LevelError is for application level errors.
	// These should trigger an alert in your APM solution.
	LevelError Level = 40

	// LevelCritical is for errors which have the potential to disrupt the application,
	// but for the moment application can handle it. panic/recover flow can log at
	// this level. Stuff like running on low disk space, nearing API quota etc.
	LevelCritical Level = 50

	// LevelFatal is for errors which lead to application crashes and application
	// cannot recover from this type of error. If event is logged at this level,
	// Logger will flush and close the Handler to avoid losing logs.
	// However, this is not guaranteed. Logger.Fatal() will invoke defined exit function.
	// (defaults to [os.Exit](1)).
	LevelFatal Level = 90
)
