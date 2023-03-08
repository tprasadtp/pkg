// SPDX-License-Identifier: BSD-3-Clause
// SPDX-FileCopyrightText: 2019 The Go Authors

package color

import "golang.org/x/sys/windows"

func isTerminal(fd uintptr) bool {
	var st uint32
	err := windows.GetConsoleMode(windows.Handle(fd), &st)
	return err == nil
}
