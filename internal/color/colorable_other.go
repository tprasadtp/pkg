// SPDX-FileCopyrightText: Copyright 2023 Prasad Tengse
// SPDX-License-Identifier: MIT

//go:build !linux && !windows && !darwin && !dragonfly && !freebsd && !netbsd && !openbsd

package color

func isColorable(flag string, istty bool) bool {
	return false
}
