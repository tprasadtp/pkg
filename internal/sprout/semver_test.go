// SPDX-FileCopyrightText: Copyright 2024 Prasad Tengse
// SPDX-License-Identifier: GPLv3-only

package sprout

import (
	"bytes"
	"testing"
	"text/template"
)

func TestSemver(t *testing.T) {
	tt := []struct {
		name     string
		template string
		expect   string
	}{}
	for _, tc := range tt {
		buf := &bytes.Buffer{}
		tpl := template.Must(template.New("test").Funcs(FuncMap()).Parse(tc.template))
		tpl.Execute(buf, nil)
		if tc.expect != buf.String() {
			t.Errorf("expected=%q, got=%q", tc.expect, buf)
		}
	}
}
