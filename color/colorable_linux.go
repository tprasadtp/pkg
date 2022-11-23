//go:build linux

package color

import "golang.org/x/sys/unix"

const ioctlReadTermios = unix.TCGETS
