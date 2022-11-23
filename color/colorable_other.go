//go:build !linux && !darwin && !windows

package color

func isColorable(flag string, istty bool) bool {
	return false
}

// isTerminal returns false on all unsupported platforms.
func isTerminal(fd int) bool {
	return false
}
