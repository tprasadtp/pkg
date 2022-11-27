//go:build windows

package color

import (
	"os"
	"strings"

	"golang.org/x/sys/windows"
)

// nolint: gochecknoglobals
var osVersion = windows.RtlGetVersion()

// isTerminal returns true if given file handle is a terminal
// and virtual terminal processing can be enabled.
func isTerminal(fd uintptr) bool {
	var handle = windows.Handle(fd)
	var mode uint32
	var err error

	err = windows.GetConsoleMode(handle, &mode)
	if err == nil {
		err = windows.SetConsoleMode(handle, mode|windows.ENABLE_VIRTUAL_TERMINAL_PROCESSING)
		if err == nil {
			return true
		}
		windows.SetConsoleMode(handle, mode)
	}
	return false
}

func isColorable(flag string, istty bool) bool {
	switch strings.TrimSpace(strings.ToLower(flag)) {
	case "never", "false", "no", "disable", "none":
		return false
	case "force", "always":
		return true
	}
	// No true color support before Windows 10 1709 (Redstone 3)
	if osVersion.BuildNumber < 16299 || osVersion.MajorVersion < 10 {
		return false
	}
	// CLICOLOR_FORCE != 0 and CLICOLOR_FORCE is not empty
	if os.Getenv("CLICOLOR_FORCE") != "0" &&
		len(strings.TrimSpace(os.Getenv("CLICOLOR_FORCE"))) > 0 {
		return true
	}
	// CLICOLOR == 0 or NO_COLOR is set and not empty
	if len(strings.TrimSpace(os.Getenv("NO_COLOR"))) > 0 ||
		os.Getenv("CLICOLOR") == "0" {
		return false
	}
	// CI
	if strings.ToLower(os.Getenv("CI")) == "true" {
		return true
	}
	return istty
}
