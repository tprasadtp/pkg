// SPDX-FileCopyrightText: Copyright 2023 Prasad Tengse
// SPDX-License-Identifier: MIT

//go:build !darwin && !dragonfly && !freebsd && !linux && !netbsd && !openbsd && !windows

package color

func isTerminal(fd uintptr) bool {
	return false
}
