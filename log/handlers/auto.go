package handlers

import (
	"github.com/tprasadtp/pkg/log"
)

// Options.
type AutoConfigOptions struct {
	// Windows Eventlog name
	WinEventLogName string

	// Fallback file
	FileName string
}

// Get best suited handler.
// Be careful with this method, as it returns an interface,
// which is not idomatic Go.
func AutoConfigure(o AutoConfigOptions) (log.Handler, error) {
	return &log.MockHandler{}, nil
}
