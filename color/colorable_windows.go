//go:build windows

package color

import (
	"os"
	"strings"

	"golang.org/x/sys/windows"
)

//nolint:gochecknoglobals // This is required to be global to mock in tests.
var osVersion = windows.RtlGetVersion()

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
	if len(os.Getenv("CLICOLOR_FORCE")) > 0 && os.Getenv("CLICOLOR_FORCE") != "0" {
		return true
	}
	// CLICOLOR == 0 or NO_COLOR is set and not empty
	if len(os.Getenv("NO_COLOR")) > 0 || os.Getenv("CLICOLOR") == "0" {
		return false
	}
	// CI
	if strings.ToLower(os.Getenv("CI")) == "true" {
		return true
	}
	return istty
}
