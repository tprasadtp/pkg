//go:build !linux && !windows && !darwin && !dragonfly && !freebsd && !netbsd && !openbsd

package color

func isColorable(flag string, istty bool) bool {
	return false
}

// isColorableTerminal returns false on all unsupported platforms.
func isColorableTerminal(fd int) bool {
	return false
}

// isTerminal returns false on all unsupported platforms.
func isTerminal(fd int) bool {
	return false
}
