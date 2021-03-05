// Package clicolors is a helper to determine whether or not ot enable colored outputs.
package clicolors

import (
	"os"
)

// EnvColorDisabled detect whether to coloring is disabled based env variables.
// Returns true if one of the conditions is true.
//
// 1. env variale CLICOLOR == 0
//
// 2. env variable NO_COLOR is set and not empty
func EnvColorDisabled() bool {
	return os.Getenv("NO_COLOR") != "" || os.Getenv("CLICOLOR") == "0"
}

// EnvColorForced detect whether to coloring is forced based env variables
// If env variale CLICOLOR_FORCE is not empty and is not set to 0
func EnvColorForced() bool {
	force, forceEnvExists := os.LookupEnv("CLICOLOR_FORCE")
	return forceEnvExists && force != "0"
}

// IsDumbTerm returns true if env variable TERM="dumb".
func IsDumbTerm() bool {
	return os.Getenv("TERM") == "dumb"
}

// EnableColors detect whether to coloring is disabled based env variables, cli-flags
// and if output is a Terminal. This supports both https://bixense.com/clicolors/
// and https://no-color.org/ standards.
//
// Flag will ALWAYS take priority. If disableColorsFlag is true, this function
// will always return false and  will ignore all environment and terminal settings.
// You should probably map this variable to your cli's --no-color/--disable-colors flag.
//
// Environment Variables (Forced)
//
// 1. If env variale CLICOLOR_FORCE != 0 function returns true if disableColorsFlag != true.
// All other environment variables and conditions are ignored.
//
// Environment Variables (Disable)
// Returns false if ANY of the following conditions are met and color is not forced.
//
// 1. If env variale CLICOLOR == 0.
//
// 2. If env variable NO_COLOR is set and is not empty (regardless of its value)
//
// 3. If env TERM is set to dumb.
//
// 4. If isTerminal is false. There is not TTY detection included in this package.
// Use IsTerminal() from package term
//
//
// Environment Variables (Enable)
//
// Returns true following conditions are met AND none of the disable conditions are true.
//
// 1. If env variale CLICOLOR != 0 and isTerminal is true.
func EnableColors(disableColorsFlag, isTerminal bool) bool {
	if disableColorsFlag {
		return false
	}

	switch {
	case EnvColorForced():
		return true
	case EnvColorDisabled() || IsDumbTerm():
		return false
	default:
		return isTerminal
	}
}
