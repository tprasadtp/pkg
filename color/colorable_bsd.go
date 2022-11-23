//go:build darwin || dragonfly || freebsd || netbsd || openbsd

package color

import "golang.org/x/sys/unix"

const ioctlReadTermios = unix.TIOCGETA
