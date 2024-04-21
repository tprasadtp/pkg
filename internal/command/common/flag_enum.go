// SPDX-FileCopyrightText: Copyright 2024 Prasad Tengse
// SPDX-License-Identifier: MIT

package common

import (
	"fmt"
	"slices"

	"github.com/spf13/pflag"
)

var _ pflag.Value = (*stringSetFlag)(nil)

func NewStringSetFlagValue(allowed []string) pflag.Value {
	return &stringSetFlag{
		allowed: allowed,
	}
}

// stringSetFlag is a custom flag which restricts flag values to ones specified.
type stringSetFlag struct {
	allowed []string
	value   string
}

func (e *stringSetFlag) String() string {
	return e.value
}

func (e *stringSetFlag) Type() string {
	return "string"
}

func (e *stringSetFlag) Set(p string) error {
	if !slices.Contains(e.allowed, p) {
		return fmt.Errorf("value %q must belong to set %v", p, e.allowed)
	}
	e.value = p
	return nil
}
