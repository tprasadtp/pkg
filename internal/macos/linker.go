// SPDX-FileCopyrightText: Copyright 2024 Prasad Tengse
// SPDX-License-Identifier: MIT

//go:build darwin

package macos

import (
	_ "crypto/x509" // for https://github.com/golang/go/issues/42459.
)
