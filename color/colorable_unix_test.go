//go:build linux || darwin

package color

import (
	"fmt"
	"testing"
)

func TestUnixTERM(t *testing.T) {
	// We never want to call t.Parallel() as they use t.Setenv
	type testCase struct {
		TERM         string
		TERM_PROGRAM string
		tty          bool
		expect       bool
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
			tty:    true,
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
			tty:    true,
			expect: false,
		},
		// screen
		{
			TERM:   "screen",
			tty:    false,
			expect: false,
		},
		{
			TERM:   "screen",
			tty:    true,
			expect: false,
		},
		// tmux
		{
			TERM:         "screen",
			TERM_PROGRAM: "tmux",
			tty:          false,
			expect:       false,
		},
		{
			TERM:         "screen",
			TERM_PROGRAM: "tmux",
			tty:          true,
			expect:       true,
		},
	}

	for _, tc := range tt {
		tn := fmt.Sprintf(
			"TERM=%s,TERM_PROGRAM=%s,tty=%t",
			tc.TERM,
			tc.TERM_PROGRAM,
			tc.tty,
		)
		// t.Run blocks till func returns, or calls t.Parallel()
		// Thus, we never want to call t.Parallel() in any of these sub-tests
		// as they use t.Setenv
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

			if tc.TERM_PROGRAM != "" {
				t.Setenv("TERM_PROGRAM", tc.TERM_PROGRAM)
			} else {
				t.Setenv("TERM_PROGRAM", "")
			}

			val := isColorable("auto", tc.tty)
			if tc.expect != val {
				t.Errorf("%s => got=%v, want=%v", tn, val, tc.expect)
			}
		})
	}
}
