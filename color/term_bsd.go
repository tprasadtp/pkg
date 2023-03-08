// SPDX-License-Identifier: BSD-3-Clause
// SPDX-FileCopyrightText: 2013 The Go Authors

//go:build darwin || dragonfly || freebsd || netbsd || openbsd

package color

import "golang.org/x/sys/unix"

const ioctlReadTermios = unix.TIOCGETA
