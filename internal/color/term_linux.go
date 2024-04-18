// SPDX-License-Identifier: BSD-3-Clause
// SPDX-FileCopyrightText: 2021 The Go Authors

//go:build linux

package color

import "golang.org/x/sys/unix"

const ioctlReadTermios = unix.TCGETS
