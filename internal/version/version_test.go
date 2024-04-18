// SPDX-FileCopyrightText: Copyright 2024 Prasad Tengse
// SPDX-License-Identifier: GPLv3-only

package version

import (
	"encoding/json"
	"testing"
)

func TestJSON(t *testing.T) {
	v := GetInfo()
	out, err := json.MarshalIndent(v, "", "\t")
	if err != nil {
		t.Error("Failed to marshal JSON")
	}
	if out == nil {
		t.Error("JSON Marshal is empty")
	}
}

func TestGetWithOverride(t *testing.T) {
	tests := []struct {
		name    string
		version string
		expect  string
	}{
		{
			name:    "with-prefix",
			version: "v1.22.333+dev",
			expect:  "v1.22.333+dev",
		},
		{
			name:    "without-prefix",
			version: "1.22.333+dev",
			expect:  "1.22.333+dev",
		},
		{
			name:    "non-semver",
			version: "2022-01-31.2",
			expect:  "2022-01-31.2",
		},
		{
			name:    "non-semver-with-prefix",
			version: "v2022-01-31.2",
			expect:  "v2022-01-31.2",
		},
	}
	for _, tc := range tests {
		t.Run(tc.name, func(t *testing.T) {
			version = tc.version
			got := GetInfo()
			if got.Version != tc.expect {
				t.Errorf("got=%v, expected=%v", got, tc.expect)
			}
		})
	}
}

// disabled because of
// - https://github.com/golang/go/issues/33976,
// - https://github.com/golang/go/issues/52600
// func TestGetWithoutOverride(t *testing.T) {
// 	info := GetInfo()
// 	if info.Version == "" {
// 		t.Errorf("GetInfo().Version s empty when it should be populated automatically")
// 	}
// }
