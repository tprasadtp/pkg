// SPDX-FileCopyrightText: Copyright 2024 Prasad Tengse
// SPDX-License-Identifier: MIT

//go:build darwin

package macos

import (
	"runtime"
	"syscall"
	_ "unsafe"
)

// CFRef is an opaque reference to a Core Foundation object.
type CFRef struct {
	pointer uintptr
	pinner  runtime.Pinner
}

func NewCFRef(pointer uintptr) *CFRef {
	rv := &CFRef{
		pointer: pointer,
	}
	rv.pinner.Pin(pointer)
	return rv
}

func (r *CFRef) Pointer() uintptr {
	return r.pointer
}

func (r *CFRef) Free() {
	r.pinner.Unpin()
	syscall_syscall(cf_trampoline_release_addr, uintptr(r.pointer), 0, 0)
}

// Defined in package [runtime] as [runtime.syscall_syscall],
// which is pushed to [syscall] as [syscall.syscall_syscall].
//
// [runtime.syscall_syscall]: https://go.googlesource.com/go/+/a10e42f219abb9c5bc4e7d86d9464700a42c7d57/src/runtime/sys_darwin.go#21
// [syscall.syscall_syscall]: https://go.googlesource.com/go/+/a10e42f219abb9c5bc4e7d86d9464700a42c7d57/src/runtime/sys_darwin.go#19
//
//go:linkname syscall_syscall syscall.syscall
//nolint:revive,nonamedreturns // ignore
func syscall_syscall(fn, a1, a2, a3 uintptr) (r1, r2 uintptr, err syscall.Errno)
