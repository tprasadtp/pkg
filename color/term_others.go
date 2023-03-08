// SPDX-License-Identifier: BSD-3-Clause
// SPDX-FileCopyrightText: 2021 The Go Authors

//go:build !darwin && !dragonfly && !freebsd && !linux && !netbsd && !openbsd && !windows

package color

func isTerminal(fd uintptr) bool {
	return false
}
