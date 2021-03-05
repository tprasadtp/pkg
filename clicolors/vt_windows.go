package clicolors

import (
	"errors"
	"os"

	"golang.org/x/sys/windows"
)

// EnableVTProcessing enables VT processing on Windows terminals.
// This uses SetConsoleMode API, and might not be supported on
// older version of Windows.
// Returns nil if BOTH stderr AND stdout handles accept SetConsoleMode params.
// On non Windows systems this simply returns nil.
func EnableVTProcessing() error {
	var originalStdOutMode uint32
	var originalStdErrMode uint32

	var stdOutColorEnablerResult bool
	var stdErrColorEnablerResult bool

	stdoutHandle := windows.Handle(os.Stdout.Fd())
	if err := windows.GetConsoleMode(stdoutHandle, &originalStdOutMode); err == nil {
		if err := windows.SetConsoleMode(stdoutHandle, originalStdOutMode|windows.ENABLE_VIRTUAL_TERMINAL_PROCESSING); err == nil {
			stdOutColorEnablerResult = true
		} else {
			// set the console mode back to original on failure
			// yes this also might fail, but mehh,
			// at this point we are not enabling VT sequences anyway.
			windows.SetConsoleMode(stdoutHandle, originalStdOutMode)
			return err
		}
	}

	stderrHandle := windows.Handle(os.Stderr.Fd())
	if err := windows.GetConsoleMode(stderrHandle, &originalStdErrMode); err == nil {
		if err := windows.SetConsoleMode(stderrHandle, originalStdErrMode|windows.ENABLE_VIRTUAL_TERMINAL_PROCESSING); err == nil {
			stdErrColorEnablerResult = true
		} else {
			// set the console mode back to original on failure
			// yes this also might fail, but mehh,
			// at this point we are not enabling VT sequences anyway.
			windows.SetConsoleMode(stderrHandle, originalStdErrMode)
			return err
		}
	}

	if stdOutColorEnablerResult && stdErrColorEnablerResult {
		return nil
	}

	return errors.New("failed to enable vt processing")
}
