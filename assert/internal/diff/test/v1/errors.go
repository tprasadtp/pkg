// Copyright (c) 2022 Prasad Tengse. All rights reserved.
// SPDX-License-Identifier: MIT

package v1

type Error struct{}

func (e Error) Error() string {
	return ""
}
