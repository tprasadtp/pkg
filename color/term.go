// SPDX-License-Identifier: BSD-3-Clause
// SPDX-FileCopyrightText: 2019 The Go Authors

package color

// IsTerminal returns whether the given file descriptor is a terminal.
func IsTerminal(fd uintptr) bool {
	return isTerminal(fd)
}
