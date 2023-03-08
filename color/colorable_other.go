//go:build !linux && !windows && !darwin && !dragonfly && !freebsd && !netbsd && !openbsd

package color

func isColorable(flag string, istty bool) bool {
	return false
}
