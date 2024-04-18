// SPDX-FileCopyrightText: Copyright 2024 Prasad Tengse
// SPDX-License-Identifier: GPLv3-only

package sprout

import (
	"strings"
	"text/template"

	"golang.org/x/text/cases"
)

func FuncMap() template.FuncMap {
	return template.FuncMap{
		"semver": semver,
		"trim":   strings.TrimSpace,
		"upper":  strings.ToUpper,
		"lower":  strings.ToLower,
		"title":  cases.Title,
		// Switch order so that "foobar" | contains "foo"
		"contains":  func(substr string, s string) bool { return strings.Contains(s, substr) },
		"hasPrefix": func(substr string, s string) bool { return strings.HasPrefix(s, substr) },
		"hasSuffix": func(substr string, s string) bool { return strings.HasSuffix(s, substr) },
	}
}
