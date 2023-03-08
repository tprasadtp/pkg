//go:build linux || darwin

package color

import (
	"fmt"
	"testing"
)

func TestUnixTERM(t *testing.T) {
	// We never want to call t.Parallel() as they use t.Setenv
	type testCase struct {
		envTERM string
		//nolint:revive,stylecheck // This is an env variable you useless linter.
		envTERM_PROGRAM string
		tty             bool
		expect          bool
	}

	var tt = []testCase{
		{
			envTERM: "",
			tty:     false,
			expect:  false,
		},
		{
			envTERM: "",
			tty:     true,
			expect:  true,
		},
		{
			envTERM: "linux",
			tty:     false,
			expect:  false,
		},
		{
			envTERM: "linux",
			tty:     true,
			expect:  false,
		},
		// dumb
		{
			envTERM: "dumb",
			tty:     false,
			expect:  false,
		},
		{
			envTERM: "dumb",
			tty:     true,
			expect:  false,
		},
		// screen
		{
			envTERM: "screen",
			tty:     false,
			expect:  false,
		},
		{
			envTERM: "screen",
			tty:     true,
			expect:  false,
		},
		// tmux
		{
			envTERM:         "screen",
			envTERM_PROGRAM: "tmux",
			tty:             false,
			expect:          false,
		},
		{
			envTERM:         "screen",
			envTERM_PROGRAM: "tmux",
			tty:             true,
			expect:          true,
		},
	}

	for _, tc := range tt {
		tn := fmt.Sprintf(
			"TERM=%s,TERM_PROGRAM=%s,tty=%t",
			tc.envTERM,
			tc.envTERM_PROGRAM,
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

			if tc.envTERM != "" {
				t.Setenv("TERM", tc.envTERM)
			} else {
				t.Setenv("TERM", "")
			}

			if tc.envTERM_PROGRAM != "" {
				t.Setenv("TERM_PROGRAM", tc.envTERM_PROGRAM)
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
