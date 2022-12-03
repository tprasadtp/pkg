//go:build windows

package color

import "golang.org/x/sys/windows"

// isColorableTerminal returns true if given file handle is a terminal
// and virtual terminal processing can be enabled.
func isColorableTerminal(fd uintptr) bool {
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
