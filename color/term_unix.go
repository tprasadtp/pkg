// SPDX-License-Identifier: BSD-3-Clause
// SPDX-FileCopyrightText: 2019 The Go Authors

//go:build darwin || dragonfly || freebsd || linux || netbsd || openbsd

package color

import "golang.org/x/sys/unix"

func isTerminal(fd uintptr) bool {
	_, err := unix.IoctlGetTermios(int(fd), ioctlReadTermios)
	return err == nil
}
