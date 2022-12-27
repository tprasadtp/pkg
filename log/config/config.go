package config

import (
	"errors"
	"io"

	"github.com/tprasadtp/pkg/log"
)

// Options.
type Options struct {
	// Windows Eventlog name
	WinEventLogName string

	// Bridge invokes log.Bridge on the logger before returning it.
	Bridge bool

	// Log file
	FileName string

	// Level
	Level log.Level

	// Skips checks
	SkipStackdriver bool

	// This is useful to use console logger on AWS Cloud shell
	// or AWS Workspaces.
	SkipCloudwatch bool

	// Meta logger
	// DO NOT USE THIS unless you wish to debug Automatic logger selection itself.
	MetaLogWriter io.Writer
}

// Create a new logger with best possible handler.
//
// # Semver Compatibility Warning
//
// This function is NOT covered by compatibility guarantees
// and might include new platform specific handlers.
//
// # Priority and Errors
//
// Handlers are selected in following order of priority.
// This function will return nil and
// [github.com/tprasadtp/pkg/log.ErrLoggerInvalid]
// if it cannot return a valid handler.
//
// # Exceptions
//
// Cloud platform specific handlers, or any os platform specific handlers
// will not be used, if ANY of the following conditions are met.
//   - Environment variable CI is set to 'true'.
//   - Environment variable CODESPACES is set to 'true'
//   - Environment variable CLOUD_SHELL is set to 'true'
//
// Handler search order or priority is given below,
//
//  1. Google Stackdriver (stackdriver).
//     This uses google-cloud-sdk for authentication.
//  2. AWS CloudWatch (cloudwatch) This uses aws-sdk for authentication.
//  3. On Linux, Journald if running as a systemd unit or user unit. (journald)
//     This will use default journald socket and is not configurable,
//     use journald handler directly, if you need more customization options.
//  4. On Windows, EventLog if running as a windows service (eventlog)
//     This requires you to define [Options.WinEventLogName],
//     or this handler is ignored.
//  5. If LogFile is specified with support for log rotation via plugin
//     (jsonfile, plugins/logrotate)
//  6. If running in container, logs to stderr.
//     This has lower priority than journal because some
//     containers might expose host system's journal.
//     This has lower priority than file, because containers
//     often mount volumes and host/cluster will handle the logs.
//     If a TTY is attached output to stderr in colored pretty print format.
//
// BUG(tprasadtp): On Linux, if journald socket is in-accessible and systemd unit
// attaches a tty via [StandardError=] and [TTYPath=] directives, logs may be
// written to stderr in pretty print format, which include ANSI escape
// codes and cause issues.
//
// [StandardError=]: https://www.freedesktop.org/software/systemd/man/systemd.exec.html#StandardError=
// [TTYPath=]: https://www.freedesktop.org/software/systemd/man/systemd.exec.html#TTYPath=
func Automatic(o Options) (*log.Logger, error) {
	if o.MetaLogWriter == nil {
		o.MetaLogWriter = io.Discard
	}
	return nil, errors.New("log.config: failed to select any log handler")
}
