// SPDX-FileCopyrightText: Copyright 2024 Prasad Tengse
// SPDX-License-Identifier: MIT

package sprout

import (
	"strings"
	"text/template"
)

func FuncMap() template.FuncMap {
	return template.FuncMap{
		"semver": semver,
		"trim":   strings.TrimSpace,
		"upper":  strings.ToUpper,
		"lower":  strings.ToLower,
		"title":  strings.Title,
		// Switch order so that "foobar" | contains "foo" works.
		"contains":  func(substr string, s string) bool { return strings.Contains(s, substr) },
		"hasPrefix": func(substr string, s string) bool { return strings.HasPrefix(s, substr) },
		"hasSuffix": func(substr string, s string) bool { return strings.HasSuffix(s, substr) },
	}
}
