// SPDX-FileCopyrightText: Copyright 2024 Prasad Tengse
// SPDX-License-Identifier: MIT

package shared

import "testing"

func TestRenderTemplateToFile(t *testing.T) {
	tt := []struct {
		name     string
		template string
		data     any
		expect   []byte
		ok       bool
	}{}

	for _, tc := range tt {
		t.Run(tc.name, func(t *testing.T) {})
	}
}
