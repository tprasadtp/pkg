// Package clicolors is a helper to determine whether or not ot enable colored outputs.
package clicolors

import (
	"os"
)

// EnableColors detect whether to coloring based env variables, cli-flag and if output is a tty.
//
// Environment Variables
//
// If env variale CLICOLOR != 0 then, ANSI colors are supported and should be used when the program isn’t piped.
// If env variale CLICOLOR == 0 then, Don’t output ANSI color escape codes.
// If env variale CLICOLOR_FORCE != 0 then, ANSI colors should be enabled no matter what.
// If both CLICOLOR_FORCE and disableColorsFlag is are true, disableColorsFlag takes precedence.
//
// Interactive terminal
//
// Please be advised that this function makes no effort to detemine if output is a TTY or
// supports ANSI colors. It is responsiblility of the user to approriately set isTerminal.
func EnableColors(disableColorsFlag, isTerminal bool) bool {
	// Flag always takes priority
	if disableColorsFlag {
		return false
	}
	force, forceEnvExists := os.LookupEnv("CLICOLOR_FORCE")
	color, colorEnvExists := os.LookupEnv("CLICOLOR")

	switch {
	case forceEnvExists && force != "0":
		return true
	case forceEnvExists && force == "0", colorEnvExists && color == "0":
		return false
	case colorEnvExists && color != "0":
		if isTerminal {
			return true
		}
		return false
	default:
		if isTerminal {
			return true
		}
	}

	return false
}
