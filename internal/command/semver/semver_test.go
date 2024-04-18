// SPDX-FileCopyrightText: Copyright 2024 Prasad Tengse
// SPDX-License-Identifier: GPLv3-only

package semver

import (
	"bytes"
	"strings"
	"testing"
	"time"

	"github.com/tprasadtp/knit/internal/testutils"
)

func TestCommand(t *testing.T) {
	tt := []struct {
		name   string
		args   []string
		output string
		ok     bool
	}{
		{
			name:   "version-valid",
			args:   []string{"version", "1.2.3"},
			output: "1.2.3",
			ok:     true,
		},
		{
			name:   "version-valid-with-prefix-v",
			args:   []string{"version", "v1.2.3"},
			output: "1.2.3",
			ok:     true,
		},
	}
	for _, tc := range tt {
		t.Run(tc.name, func(t *testing.T) {
			ctx, cancel := testutils.TestingContext(t, time.Second*5)
			defer cancel()

			stdout := &bytes.Buffer{}
			stderr := &bytes.Buffer{}
			cmd := NewCommand()
			cmd.SetOut(stdout)
			cmd.SetErr(stderr)
			cmd.SetArgs(tc.args)
			err := cmd.ExecuteContext(ctx)
			if tc.ok {
				if err != nil {
					t.Errorf("expected no error, got %s", err)
				}
				if tc.output != strings.TrimSpace(stdout.String()) {
					t.Errorf("expected output=%q, got=%q", tc.output, stdout)
				}
			} else {
				if err == nil {
					t.Errorf("expected an error, got nil")
				}
				if stdout.Len() != 0 {
					t.Errorf("expected no output but got: %q", stderr)
				}
			}
		})
	}
}
