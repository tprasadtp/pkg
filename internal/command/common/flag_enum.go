// SPDX-FileCopyrightText: Copyright 2024 Prasad Tengse
// SPDX-License-Identifier: MIT

package common

import (
	"fmt"
	"slices"

	"github.com/spf13/pflag"
)

var _ pflag.Value = (*stringEnumFlag)(nil)

// stringEnumFlag is a custom flag which restricts flag values to ones specified.
type stringEnumFlag struct {
	allowed []string
	value   string
}

func (e *stringEnumFlag) String() string {
	return e.value
}

func (e *stringEnumFlag) Type() string {
	return "string"
}

func (e *stringEnumFlag) Set(p string) error {
	if !slices.Contains(e.allowed, p) {
		return fmt.Errorf("")
	}
	e.value = p
	return nil
}
