// SPDX-FileCopyrightText: Copyright 2023 Prasad Tengse
// SPDX-License-Identifier: MIT

//go:build windows

package guid

import (
	"golang.org/x/sys/windows"
)

// GUID represents a GUID/UUID.
//
// It has the same structure as
// [golang.org/x/sys/windows.GUID] so that it can be used with functions
// expecting that type. It is defined as its own type so that [fmt.Stringer],
// [json.Marshaler], [json.Unmarshaler], [encoding.TextMarshaler] and
// [encoding.TextUnmarshaler] can be supported. The representation matches
// that used by native Windows code.
type GUID windows.GUID
