//go:build linux || darwin

package color

import (
	"fmt"
	"testing"
)

func TestUnixTERM(t *testing.T) {
	// This test MUST not be parallel
	type testCase struct {
		TERM   string
		tty    bool
		expect bool
	}

	var tt = []testCase{
		{
			TERM:   "",
			tty:    false,
			expect: false,
		},
		{
			TERM:   "",
			tty:    true,
			expect: true,
		},
		{
			TERM:   "linux",
			tty:    false,
			expect: false,
		},
		{
			TERM:   "linux",
			tty:    false,
			expect: false,
		},
		// dumb
		{
			TERM:   "dumb",
			tty:    false,
			expect: false,
		},
		{
			TERM:   "dumb",
			tty:    false,
			expect: false,
		},
	}

	for _, tc := range tt {
		tn := fmt.Sprintf(
			"TERM=%s,tty=%t",
			tc.TERM,
			tc.tty,
		)
		t.Run(tn, func(t *testing.T) {
			t.Setenv("CLICOLOR_FORCE", "")
			t.Setenv("CLICOLOR", "")
			t.Setenv("NO_COLOR", "")
			t.Setenv("CI", "")

			if tc.TERM != "" {
				t.Setenv("TERM", tc.TERM)
			} else {
				t.Setenv("TERM", "")
			}

			val := isColorable("auto", tc.tty)
			if tc.expect != val {
				t.Errorf("%s => got=%v, want=%v", tn, val, tc.expect)
			}
		})
	}
}
