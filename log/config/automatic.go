package config

import (
	"io"

	"github.com/tprasadtp/pkg/log"
)

// Options.
type AutomaticOptions struct {
	// Windows Eventlog name
	WinEventLogName string

	// Fallback file
	FileName string

	// Level
	Level log.Level

	// Skip checks
	// This is useful to use console logger on Google Cloud Shell
	// or Google Cloud Workspace.
	SkipStackdriver bool

	// This is useful to use console logger on AWS Cloud shell
	// or AWS Workspaces.
	SkipCloudwatch bool

	// Meta logger
	// DO NOT USE THIS unless you wish to debug Automatic logger selection itself.
	MetaLogWriter io.Writer
}

// Get the logger with best possible handler.
// This is NOT covered by compatibility guarantees.
// Handlers are selected in following order of priority.
// This function will return nil and
// [github.com/tprasadtp/pkg/log.ErrLoggerInvalid]
// if it cannot return a handler.
//
//  1. Google Stackdriver (stackriver).
//     This uses google-cloud-sdk for authentication.
//  2. AWS CloudWatch (cloudwatch) This uses aws-sdk for authentication.
//  3. On Linux, Journald if running as a systemd unit or user unit. (journald)
//     This will use default journald socket and is not configurable,
//     use journald handler directly, if you need more customization options.
//  4. On Windows, EventLog if running as a windows service (eventlog)
//     This requires you to define WinEventLogName in your handler options
//     or this will be ignored.
//  5. If running in container, stderr in json format.
//     This has lower priority than journal because some
//     container run-times (like podman) expose host system's journald for
//     containers as well.
//  6. If Log File is specified (it is not by default)  with support for
//     log rotation via plugin (jsonfile, plugins/logrotate)
//  7. If TTY is attached, output to stderr in pretty print format (console)
func Automatic(o AutomaticOptions) (*log.Logger, error) {
	if o.MetaLogWriter == nil {
		o.MetaLogWriter = io.Discard
	}
	return nil, log.ErrLoggerInvalid
}
