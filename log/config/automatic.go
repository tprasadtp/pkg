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
	// This is useful to use console logger on Google Cloud Shell.
	SkipStackdriver bool
	// This is useful to use console logger on Google Cloud Shell.
	SkipCloudwatch bool

	// Meta logger
	// DO NOT USE THIS unless you wish to debug Automatic logger selection itself.
	MetaLogWriter io.Writer
}

// Automatically configures the logger and its handler.
// HAndlers are selected in following order of priority
//  1. Google Stackdriver (stackriver)
//  2. AWS CloudWatch (cloudwatch)
//  3. On Linux, Journald if running as a systemd unit or user unit. (journald)
//  4. On Windows, EventLog if running as a windows service (eventlog)
//  5. If running in container, stderr in json format.
//     This has lower priority than journal because some
//     container run-times (like podman) do expose host system's
//     journald socket for containers.
//  6. If Log File is specified  with support for
//     log rotation via plugin (jsonfile, plugins/logrotate)
//  7. If TTY is attached, output to stderr in pretty print format (console)
//
// This is a complicated process and as it involves setting up the logging
// itself errors and debug information are not logged.
func Automatic(o AutomaticOptions) (*log.Logger, error) {
	if o.MetaLogWriter == nil {
		o.MetaLogWriter = io.Discard
	}
	return nil, log.ErrLoggerInvalid
}
