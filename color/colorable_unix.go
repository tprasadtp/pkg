//go:build linux || darwin || dragonfly || freebsd || netbsd || openbsd

package color

import (
	"os"
	"strings"

	"golang.org/x/sys/unix"
)

// isTerminal returns true if given file descriptor is a terminal
func isTerminal(fd uintptr) bool {
	_, err := unix.IoctlGetTermios(int(fd), ioctlReadTermios)
	return err == nil
}

// isTerminal returns true if colors are forced or can be enabled.
func isColorable(flag string, istty bool) bool {
	switch strings.TrimSpace(strings.ToLower(flag)) {
	case "never", "false", "no", "disable", "none":
		return false
	case "force", "always":
		return true
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

	// screen, dumb and linux ttys do not support 24 bit color.
	switch t := os.Getenv("TERM"); {
	case t == "linux":
		return false
	case t == "dumb":
		return false
	case strings.HasPrefix(t, "screen"):
		// tmux supports 24 bit color but screen does not
		if os.Getenv("TERM_PROGRAM") == "tmux" {
			return istty
		}
		return false
	}

	return istty
}
