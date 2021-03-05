package clicolors

// EnableVTProcessing enables VT processing on Windows terminals.
// This uses SetConsoleMode API, and might not be supported on
// older version of Windows.
// Returns nil if BOTH stderr AND stdout handles accept SetConsoleMode params.
// On non Windows systems this simply returns nil.
func EnableVTProcessing() error {
	return nil
}
