// Copyright (c) 2020 The Go Authors. All rights reserved.
// Spdx-License-Identifier: BSD-3-Clause OR Apache-2.0

package slices

// Clip removes unused capacity from the slice, returning s[:len(s):len(s)].
func Clip[S ~[]E, E any](s S) S {
	return s[:len(s):len(s)]
}
