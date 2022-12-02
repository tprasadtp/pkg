// Automatically selects the most suitable handler.
//   - On windows if running as a service, this is the WindowsEvent log
//   - On Linux if running as systemd unit (user or system) this is journald
//   - Fallback to specified file if above conditions are not met.
package auto

import (
	"os"

	"github.com/tprasadtp/pkg/log"
)

// Options.
type Options struct {
	// Windows Eventlog name
	WinEventLogName string

	// Fallback file
	File *os.File
}

// Get best suited handler.
// Be careful with this method, as it returns an interface,
// which is not idomatic Go.
func Configure(o Options) (log.Handler, error) {
	return &log.MockHandler{}, nil
}
