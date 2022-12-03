package config

import (
	"github.com/tprasadtp/pkg/log"
)

// Options.
type AutomaticOptions struct {
	// Windows Eventlog name
	WinEventLogName string

	// Fallback file
	FileName string

	// Level
	level log.Level
}

// Get best suited handler.
// Be careful with this method, as it returns an interface,
// which is not idomatic Go.
func Automatic(o AutomaticOptions) log.Handler {
	h, err := config(o)
	if err != nil {
		panic(err)
	}
	return h
}
